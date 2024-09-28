package controlador

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "github.com/dhowden/tag"
)

// Mine recolecta todas las etiquetas de los archivos mp3 en el directorio
// especificado y las inserta en la base de datos.
func Mine(path string) error {
    getTags(fullPath())
    return nil
}

// Honestamente no sé qué hace esta función muy bien pero básicamente busca
// archivos mp3 en el directorio y procesa el archivo.
func getTags(path string) (err error) {
    err = filepath.Walk(path, func(pathString string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() && strings.HasSuffix(info.Name(), ".mp3") {
            processEntry(pathString)
        }

        return nil
    })
    if err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}

// Esta función está pensada para insertar esta canción a la base de datos
// o no hacer nada en caso de que ya exista.
func processEntry(path string) {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("Error al leer las etiquetas: ", err)
        return
    }
    defer file.Close()
    metadata, err := tag.ReadID3v2Tags(file)
    fmt.Println("Performer(s): ", metadata.Artist())
    fmt.Println("Title: ", metadata.Title())
    fmt.Println("Album: ", metadata.Album())
    fmt.Println("Year: ", metadata.Year())
    fmt.Println("Genre: ", metadata.Genre())
    number, total := metadata.Track()
    fmt.Println("Track: ", number, "/", total)
}

func fullPath() string {
    home, _ := os.UserHomeDir()
    return filepath.Join(home, "Documentos", "Prueba")
}
