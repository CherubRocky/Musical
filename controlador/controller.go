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
    controller, err := NewController()
    if err != nil {
        return err
    }
    defer controller.DB.Close()
    controller.View.MineButton.OnTapped = func() {controller.ActionMine()}
    controller.View.SearchButton.OnTapped = func() {fmt.Println("Se tapeó la búsqueda")}
    controller.updateViewSong()
    controller.View.RunView()
    return nil
}

func PressedPlay(song vista.ViewSong) {
}

func (controller *Controller) ActionMine() {
    controller.View.DisableElements()
    Mine("", controller.DB)
    controller.View.EnableElements()
    err := controller.updateViewSong()
    if err != nil {
        fmt.Println("Algo salió mal al hacer un update de la lista de caciones")
    }
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
