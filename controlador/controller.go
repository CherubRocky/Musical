package controlador

import (
    "fmt"
    "github.com/CherubRocky/Musical/modelo"
    "github.com/CherubRocky/Musical/vista"
)

type Controller struct {
    View *vista.View
    DB   *modelo.MusicalDB
}

func NewController() (*Controller, error) {
    c := &Controller{
        View: vista.NewView(PressedPlay),
    }
    db, err := modelo.NewMusicalDB()
    if err != nil {
        return nil, err // Retorna nil para el controlador si hay un error
    }
    c.DB = db
    return c, err
}

func Run() error {
    musicDir, lilErr := modelo.GetMusicDir()
    if lilErr != nil {
        return lilErr
    }
    if !modelo.FileExists(modelo.GetDBPath()) {
        err := modelo.CreateDBFile()
        if err != nil {
            return err
        }
    }
    controller, err := NewController()
    if err != nil {
        return err
    }
    defer controller.DB.Close()
    controller.View.MineButton.OnTapped = func() {controller.ActionMine(musicDir)}
    controller.View.SearchButton.OnTapped = func() {controller.searchTap()}
    controller.updateViewSong()
    controller.View.RunView()
    return nil
}

func PressedPlay(song vista.ViewSong) {
}

func (controller *Controller) ActionMine(minerDirPath string) {
    controller.View.DisableElements()
    progressChan := make(chan float64)
    controller.View.MakeProgressBar()
    controller.View.ShowBar()
    go Mine(minerDirPath, controller.DB, progressChan)
    go func() {
        for progressValue := range progressChan {
            controller.View.UpdateProgressBar(progressValue)
        }
        controller.View.EnableElements()
        err := controller.updateViewSong()
        if err != nil {
            fmt.Println("Algo salió mal al hacer un update de la lista de caciones")
        }
        controller.View.ProgressWin.Close()
    }()

}

func (controller *Controller) updateViewSong() error {
    modelSongList, err := controller.DB.GeneralQuery()
    if err != nil {
        return err
    }
    *controller.View.Songs = (*controller.View.Songs)[:0]
    for _, song := range modelSongList {
        lilSong := vista.ViewSong {
            ID:        song.ID,
            Title:     song.Title,
            Performer: song.Performer,
            Album:     song.Album,
            Path:      song.Path,
        }
        *controller.View.Songs = append(*controller.View.Songs, lilSong)
    }
    controller.View.SongList.Refresh()
    return nil
}

func (controller *Controller) searchTap() {
    text := controller.View.SearchBar.Text
    defer controller.View.SearchBar.SetText("")
    modelSong, err := modelo.ProcessEntry(controller.DB, text)
    if err != nil {
        controller.View.ShowErrorDialog(err)
        err = controller.updateViewSong()
        if err != nil {
            controller.View.ShowErrorDialog(err)
        }
        return
    }
    controller.convertSongs(modelSong)
    controller.View.SongList.Refresh()
    fmt.Println(text)
}

func (controller *Controller) convertSongs(modelSongList []modelo.Song) {
    *controller.View.Songs = (*controller.View.Songs)[:0]
    for _, song := range modelSongList {
        lilSong := vista.ViewSong {
            ID:        song.ID,
            Title:     song.Title,
            Performer: song.Performer,
            Album:     song.Album,
            Path:      song.Path,
        }
        *controller.View.Songs = append(*controller.View.Songs, lilSong)
    }
}
