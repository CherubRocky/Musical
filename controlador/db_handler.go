package music_db

import (
    "fmt"
    "database/sql"
    "controlador"
)

type MusicalDB struct {
    DB *sql.DB
}

func NewMusicalDB() (*MusicalDB, error) {
    db, err := sql.Open("sqlite3", "../bd/musical.db")
    if err != nil {
        return nil, err
    }
    return &MusicalDB {
        DB: db,
    }
}

func (mDB *MusicalDB) InsertMinedSong(tags controlador.SongTags, path string) {
    //podemos usar prepared
    var insrtAlbum, insrtPerformer, insrtSong string
    insrtPerformer = `INSERT INTO performers (id_performer, name) VALUES (?, ?)`
    insrtAlbum = `INSERT INTO albums (id_album, name) VALUES (?, ?)`
    insrtSong = `INSERT INTO rolas (path, title, track, year, genre) VALUES (?, ?, ?, ?, ?)`
    statementP, err1 := mDB.DB.Prepare(insrtPerformer)
    statementA, err2 := mDB.DB.Prepare(insrtAlbum)
    statementS, err3 := mDB.DB.Prepare(insrtSong)

}

func (mDB *MusicalDB) insertPerformer(name string) (int, error) {
    id, err := mDB.existsPerformer(name)
    if err != nil {
        if err != sql.ErrNoRows {
            return -1, err
        }
        return id, nil
    }
    stmt, err := mDB.DB.Prepare(`INSERT INTO performers (id_performer, name) VALUES (?, ?)`)
    if err != nil {
        return -1, err
    }
    id, err = getNewID("performers", "id_performer")
    if err != nil {
        return -1, err
    }
    _, err = stmt.Exec(id, name)
    if err !=  nil {
        return -1, err
    }
    return id, err
}

func (mDB *MusicalDB) getNewID(table, field string) (int, error) {
    var id int
    query := fmt.Sprintf(`SELECT COALESCE(MAX(%s), 0) FROM %s`, field, table)
    err := mDB.DBQueryRow(query).Scan(&id)
    if err != nil {
        return -1, err
    }
    return id + 1, nil
}

func (mDB *MusicalDB) insertAlbum() int {

}

func (mDB *MusicalDB) insertSong() {

}

func (mDB *Musical) existsPerformer(name string) (int, err) {
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

func (mDB *MusicalDB) Close() error {
    return mDB.DB.Close()
}

func (mDB *MusicalDB) SongExists(tags controlador.SongTags) (bool, error) {
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


// "SELECT rolas.title, albums.name, performers.name FROM rolas, albums, performers
// INNER JOIN albums
// ON rolas.id_albums == albums.id WHERE rolas.title == ?
// INNER JOIUN performers
// ON rolas.id_performer == performers.id WHERE rolas.title == ?"
