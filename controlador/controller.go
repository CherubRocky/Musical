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
    controller.View.RunView()
    return nil
}

func PressedPlay(song vista.ViewSong) {
}

func (controller *Controller) ActionMine() {
    controller.View.DisableElements()
    Mine("", controller.DB)
    controller.View.EnableElements()
}
