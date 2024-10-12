package modelo

import (
    "fmt"
    "database/sql"
    "path/filepath"
    _ "github.com/mattn/go-sqlite3"
)

type MusicalDB struct {
    DB *sql.DB
}

func NewMusicalDB() (*MusicalDB, error) {
    db, err := sql.Open("sqlite3", GetDBPath())
    if err != nil {
        return nil, err
    }
    return &MusicalDB {
        DB: db,
    }, nil
}

func (mDB *MusicalDB) InsertMinedSong(tags SongTags, path string) (int, error) {
    idPerformer, err := mDB.insertPerformer(tags.Performer)
    if err != nil {
        return -1, err
    }
    idAlbum, err := mDB.insertAlbum(tags, path, idPerformer)
    if err != nil {
        return -1, err
    }
    insrt := `INSERT INTO rolas (id_rola, id_performer, id_album, path, title,
                track, year, genre) VALUES (?,?,?,?,?,?,?,?)`
    stmt, err := mDB.DB.Prepare(insrt)
    defer stmt.Close()
    if err != nil {
        return -1, err
    }
    id, err := mDB.getNewID("rolas", "id_rola")
    if err != nil {
        return -1, err
    }
    _, err = stmt.Exec(id, idPerformer, idAlbum, path, tags.Title, tags.Track,
        tags.Year, tags.Genre)
    if err != nil {
        return -1, err
    }
    return id, nil
}

func (mDB *MusicalDB) insertPerformer(name string) (int, error) {
    id, err := mDB.getPerformerID(name)
    if err != nil {
        if err != sql.ErrNoRows {
            return -1, err
        }
        stmt, err := mDB.DB.Prepare(`INSERT INTO performers (id_performer, name) VALUES (?, ?)`)
        defer stmt.Close()
        if err != nil {
            return -1, err
        }
        id, err = mDB.getNewID("performers", "id_performer")
        if err != nil {
            return -1, err
        }
        _, err = stmt.Exec(id, name)
        if err !=  nil {
            return -1, err
        }
        return id, err
    }
    return id, nil
}

func (mDB *MusicalDB) getNewID(table, field string) (int, error) {
    var id int
    query := fmt.Sprintf(`SELECT COALESCE(MAX(%s), 0) FROM %s`, field, table)
    err := mDB.DB.QueryRow(query).Scan(&id)
    if err != nil {
        return -1, err
    }
    return id + 1, nil
}

func (mDB *MusicalDB) insertAlbum(tags SongTags, path string, idPerformer int) (int, error) {
    id, err := mDB.getAlbumID(tags.Album, idPerformer)
    if err != nil {
        if err != sql.ErrNoRows {
            return -1, err
        }
        // Si no existe el album (hace inserción)
        stmt, err := mDB.DB.Prepare("INSERT INTO albums (id_album, path, name, year) VALUES (?,?,?,?)")
        defer stmt.Close()
        if err != nil {
            return -1, err
        }
        id, err = mDB.getNewID("albums", "id_album")
        if err != nil {
            return -1, err
        }
        albumPath := filepath.Dir(path)
        _, err = stmt.Exec(id, albumPath, tags.Album, tags.Year)
        if err != nil {
            return -1, err
        }
        return id, nil
    }
    // Si el album ya existía (no hacer inserción)
    return id, nil
}


func (mDB *MusicalDB) insertSong() {

}

func (mDB *MusicalDB) getPerformerID(name string) (int, error) {
    var id int
    id = -1
    query := `SELECT performers.id_performer FROM performers WHERE performers.name = ?`
    err := mDB.DB.QueryRow(query, name).Scan(&id)
    if err == sql.ErrNoRows {
        return id, err
    }
    if err != nil {
        return id, err
    }
    return id, nil
}

func (mDB *MusicalDB) getAlbumID(albumName string, idPerformer int) (int, error) {
    var id int
    id = -1
    query := `SELECT albums.id_album FROM albums
                INNER JOIN rolas ON rolas.id_album = albums.id_album
                INNER JOIN performers ON performers.id_performer = rolas.id_performer
                WHERE albums.name = ? AND performers.id_performer = ?`
    err := mDB.DB.QueryRow(query, albumName, idPerformer).Scan(&id)
    if err == sql.ErrNoRows {
        return id, err
    }
    if err != nil {
        return id, err
    }
    return id, nil
}

func (mDB *MusicalDB) Close() error {
    return mDB.DB.Close()
}

func (mDB *MusicalDB) SongExists(tags SongTags) (bool, error) {
    var exists bool
    query := `SELECT EXISTS (
                SELECT 1
                FROM rolas
                INNER JOIN albums
                ON rolas.id_album = albums.id_album
                INNER JOIN performers
                ON rolas.id_performer = performers.id_performer
                WHERE rolas.title = ? AND albums.name = ? AND performers.name = ?
                )`
    err := mDB.DB.QueryRow(query, tags.Title, tags.Album, tags.Performer).Scan(&exists)
    return exists, err
}

func (mDB *MusicalDB) GeneralQuery() ([]Song, error) {
    query := `SELECT rolas.id_rola, rolas.title, rolas.path, performers.name, albums.name FROM rolas
                INNER JOIN performers ON performers.id_performer = rolas.id_performer
                INNER JOIN albums ON albums.id_album = rolas.id_album`
    rows, err := mDB.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var songs []Song
    for rows.Next() {
        var song Song
        if err := rows.Scan(&song.ID, &song.Title, &song.Path, &song.Performer, &song.Album); err != nil {
            return nil, err
        }
        songs = append(songs, song)
    }
    return songs, nil
}
