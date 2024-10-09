package controlador

import (
    "fmt"
    "os"
    "time"
    "path/filepath"
    "strings"
    "github.com/dhowden/tag"
    "github.com/CherubRocky/Musical/modelo"
)

// Mine recolecta todas las etiquetas de los archivos mp3 en el directorio
// especificado y las inserta en la base de datos.
func Mine(path string, mDB *modelo.MusicalDB) error {
    processMP3Files(fullPath(), mDB)
    return nil
}

// processMP3Files itera y busca archivos mp3 en el directorio
// y procesa el archivo.
func processMP3Files(path string, mDB *modelo.MusicalDB) (err error) {
    err = filepath.Walk(path, func(pathString string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() && strings.HasSuffix(info.Name(), ".mp3") {
            tags, err := getSongTags(pathString)
            if err != nil {
                fmt.Println("Hubo un error al leer las etiquetas.")
            }
            _, err = mDB.InsertMinedSong(tags, pathString)
            if err != nil {
                fmt.Sprintf("Hubo un error al insertar la canción.")
                return err
            }
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
func getSongTags(path string) (modelo.SongTags, error){
    file, err := os.Open(path)
    if err != nil {
        return modelo.SongTags{}, fmt.Errorf("Error al intentar leer etiquetas: %v", err)
    }
    defer file.Close()
    metadata, err := tag.ReadID3v2Tags(file)
    data := modelo.SongTags {
        Title:      getStringTag(metadata.Title()),
        Performer:  getStringTag(metadata.Artist()),
        Album:      getStringTag(metadata.Album()),
        Track:      getTrackTag(metadata.Track()),
        Year:       getYearTag(metadata.Year()),
        Genre:      getStringTag(metadata.Genre()),
    }
    fmt.Println("Performer(s): ", metadata.Artist())
    fmt.Println("Composer: ", metadata.Composer())
    fmt.Println("Album Artist ", metadata.AlbumArtist())
    fmt.Println("Title: ", metadata.Title())
    fmt.Println("Album: ", metadata.Album())
    fmt.Println("Year: ", metadata.Year())
    fmt.Println("Genre: ", metadata.Genre())
    number, total := metadata.Track()
    fmt.Println("Track: ", number, "/", total)
    return data, nil
}

func fullPath() string {
    home, _ := os.UserHomeDir()
    return filepath.Join(home, "Documentos", "Prueba")
}

func getStringTag(tag string) string {
    if tag == "" {
        return "Unknown"
    }
    return tag
}

func getYearTag(year int) int {
    if year == 0 {
        now := time.Now()
        return now.Year()
    }
    return year
}

func getTrackTag(track int, total int) int {
    return track
}
