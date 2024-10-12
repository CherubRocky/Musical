package modelo
import (
    "fmt"
    "os"
    "log"
    "path/filepath"
    "database/sql"
    "strings"
    _ "github.com/mattn/go-sqlite3"
)

func FileExists(path string) bool {
    _,err := os.Stat(path)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

func getAppConfigFile() (string, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    file := filepath.Join(home, ".config", "Musical", "musical_info.txt")
    return file, nil
}

func CreateFile(path string) (*os.File, error) {
    parent := filepath.Dir(path)
    err := os.MkdirAll(parent, 0755)
    if err != nil {
        return nil, err
    }
    file, err := os.Create(path)
    if err != nil {
        return nil, err
    }
    return file, nil
}

func GetMusicDir() (string, error) {
    appFile, err := getAppConfigFile()
    if !FileExists(appFile) {
        file, err := CreateFile(appFile)
        if err != nil {
            return "", err
        }
        defer file.Close()
        // checar si su sistema está en inglés o en español
        mDir := getDefaultMusicPath()
        _, err = file.WriteString(mDir)
        if err != nil {
            return "", err
        }
        return mDir, nil
    }
    file, err := os.Open(appFile)
    if err != nil {
        return "", err
    }
    defer file.Close()
    buffer := make([]byte, 1024)
    n, err := file.Read(buffer)
    if err != nil {
        return "", err
    }
    fmt.Println(string(buffer[:n]))
    return string(buffer[:n]), nil
}

func getDefaultMusicPath() string {
    lang := os.Getenv("LANG")
    home, _ := os.UserHomeDir()
    if strings.Contains(strings.ToLower(lang), "es") {
        return filepath.Join(home, "Música")
    }
    return filepath.Join(home, "Music")
}
// /usr/local/share/Musical
func CreateDBFile() error {
    dbFile := GetDBPath()
    lilFile, err := CreateFile(dbFile)
    defer lilFile.Close()
    if err != nil {
        return err
    }
    db, err := sql.Open("sqlite3", dbFile)
    if err != nil {
        return err
    }
    defer db.Close()
    statements := []string{
        `CREATE TABLE types (
            id_type       INTEGER PRIMARY KEY,
            description   TEXT
        );`,
        `INSERT INTO types VALUES(0,'Person');`,
        `INSERT INTO types VALUES(1,'Group');`,
        `INSERT INTO types VALUES(2,'Unknown');`,
        `CREATE TABLE performers (
            id_performer  INTEGER PRIMARY KEY,
            id_type       INTEGER,
            name          TEXT,
            FOREIGN KEY   (id_type) REFERENCES types(id_type)
        );`,
        `CREATE TABLE persons (
            id_person     INTEGER PRIMARY KEY,
            stage_name    TEXT,
            real_name     TEXT,
            birth_date    TEXT,
            death_date    TEXT
        );`,
        `CREATE TABLE groups (
            id_group      INTEGER PRIMARY KEY,
            name          TEXT,
            start_date    TEXT,
            end_date      TEXT
        );`,
        `CREATE TABLE in_group (
            id_person     INTEGER,
            id_group      INTEGER,
            PRIMARY KEY   (id_person, id_group),
            FOREIGN KEY   (id_person) REFERENCES persons(id_person)
            FOREIGN KEY   (id_group) REFERENCES groups(id_group)
        );`,
        `CREATE TABLE albums (
            id_album      INTEGER PRIMARY KEY,
            path          TEXT,
            name          TEXT,
            year          INTEGER
        );`,
        `CREATE TABLE rolas (
            id_rola       INTEGER PRIMARY KEY,
            id_performer  INTEGER,
            id_album      INTEGER,
            path          TEXT,
            title         TEXT,
            track         INTEGER,
            year          INTEGER,
            genre         TEXT,
            FOREIGN KEY   (id_performer) REFERENCES performers(id_performer)
            FOREIGN KEY   (id_album) REFERENCES albums(id_album)
        );`,
    }
    for _, statement := range statements {
        _, err := db.Exec(statement)
        if err != nil {
            return fmt.Errorf("Error al crear la base de datos.")
        }
    }
    return nil
}

func GetDBPath() string {
    home, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    return filepath.Join(home, ".local", "share", "Musical", "musical.db")
}

/*
checar si existe el archivo
*/
