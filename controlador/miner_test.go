package controlador
import (
    "testing"
    "github.com/CherubRocky/Musical/modelo"
)

func TestGetSongTags(t *testing.T) {
    esperado := modelo.SongTags{"T", "A", "L", 3, 1950, "8"}
    datos, err := getSongTags("../full.mp3")
    if err != nil {
        t.Errorf("getSongTags() error: %v", err)
    }
    if esperado != datos {
        t.Errorf("getSongTags() falló. Esperado: %v\nObtenido: %v", esperado, datos)
    }
}
