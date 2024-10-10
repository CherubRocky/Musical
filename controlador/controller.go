package controlador

import (
    "fmt"
    //"github.com/CherubRocky/Musical/modelo"
    "github.com/CherubRocky/Musical/vista"
)

type Controller struct {
    View *vista.View
}

func NewController() *Controller {
    c := &Controller{vista.NewView(PressedPlay)}
    return c
}

func Run() {
    controller := NewController()
    controller.View.MineButton.OnTapped = func() {fmt.Println("Se tapeó el minero")}
    controller.View.SearchButton.OnTapped = func() {fmt.Println("Se tapeó la búsqueda")}
    controller.View.RunView()
}

func PressedPlay(song vista.ViewSong) {
}
