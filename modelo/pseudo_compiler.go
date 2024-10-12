package modelo

import (
    "strings"
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
    return nil, nil
}
