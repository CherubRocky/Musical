package modelo

import (
    "strings"
    "regexp"
)

func ProcessEntry(mDB *MusicalDB, entry string) ([]Song, error) {
    if entry == "*" {
        return mDB.GeneralQuery()
    }
    commaCount := strings.Count(entry, ",")
    dotCount := strings.Count(entry, ":")
    if commaCount == 0 {
        if dotCount == 0 {
            return mDB.BasicTitleQuery(entry)
        }
    }
    query, args := buildQuery(filter(entry))
    return mDB.VariableQuery(query, args)
}

func filter(entry string) Song {
    lilFilter := Song{}
    re := regexp.MustCompile(`(\w+):([^,]+)`) // Buscar patrones como "p:Luis Miguel"

    matches := re.FindAllStringSubmatch(entry, -1)
    for _, match := range matches {
        key := match[1]
        value := strings.TrimSpace(match[2])
        switch key {
        case "p":
            lilFilter.Performer = value
        case "t":
            lilFilter.Title = value
        case "a":
            lilFilter.Album = value
        }
    }

    return lilFilter
}

func buildQuery(lilFilter Song) (string, []interface{}) {
    query := `SELECT rolas.id_rola, rolas.title, rolas.path, performers.name, albums.name FROM rolas
                INNER JOIN performers ON performers.id_performer = rolas.id_performer
                INNER JOIN albums ON albums.id_album = rolas.id_album
                WHERE 1 = 1`

    var args []interface{}

    if lilFilter.Title != "" {
        query += " AND rolas.title LIKE ?"
        args = append(args, "%" + lilFilter.Title + "%")
    }
    if lilFilter.Performer != "" {
        query += " AND performers.name LIKE ?"
        args = append(args, "%" + lilFilter.Performer + "%")
    }
    if lilFilter.Album != "" {
        query += " AND albums.name LIKE ?"
        args = append(args, "%" + lilFilter.Album + "%")
    }
    return query, args
}
