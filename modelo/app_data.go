package modelo
import (
    "fmt"
    "os"
    "path/filepath"
)

func fileExists(path string) bool {
    _,err := os.Stat(path)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

func getAppConfigFile() (string, err) {
    home, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    dir := filepath.Join(home, ".config", "Musical", "info.txt")
    return dir, nil
}

func getDBFile() (string, err) {
    home, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    dir := filepath.Join("/usr","local")
    return dir, nil
}

func thing() {
    dbFile, err := getDBFile()
    appFile, err:= getAppConfigFile()
    if err != nil {
    }
    if !fileExists(dbFile) {
        CreateFile(dbFile)
    }
    if !fileExists(appFile) {
        CreateFile(appFile)
    }
}

func CreateFile(path string) (*os.File, error) {
    parent := filepath.Dir(path)
    err := os.MkdirAll(parent)
    if err != nil {
        return nil, err
    }
    file, err := os.Create(path)
    if err != nil {
        return nil, err
    }
    return file, nil
}

func getMusicDir() (string, error) {
    appFile, err := getAppConfigFile()
    if !fileExists(appFile) {
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
    return string(buffer[:n]), nil
}

func getDefaultMusicPath() string {
    lang := os.Gentenv("LANG")
    if strings.Contains(strings.ToLower(lang), "es") {
        return filepath.Join(os.UserHomeDir(), "Música")
    }
    return filepath.Join(os.UserHomeDir(), "Music")
}
