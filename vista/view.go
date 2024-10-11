package vista

import (
    //"fmt"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

type ViewSong struct {
    ID        int
    Title     string
    Performer string
    Album     string
    Path      string
}

type View struct {
    App          fyne.App
    Window       fyne.Window
    MineButton   *widget.Button
    SearchButton *widget.Button
    ProgressWin  fyne.Window
    ProgressBar  *widget.ProgressBar
    SearchBar    *widget.Entry
    SongList     *widget.List
    Songs        *[]ViewSong
}

func NewView(listButtonHandler func(song ViewSong)) *View {
    a := app.New()
    w := a.NewWindow("Reproductor de rolas")
    mine := widget.NewButton("Importar canciones", nil)
    queryB := widget.NewButton("Buscar", nil)
    bar := widget.NewEntry()
    bar.SetPlaceHolder("Buscar...")
    topContainer := container.NewGridWithColumns(2, bar, queryB)
    songs := make([]ViewSong, 0)
    songList := CreateSongList(&songs, listButtonHandler)
    content := container.NewBorder(topContainer, mine, nil, nil, songList)
    w.SetContent(content)
    w.Resize(fyne.NewSize(600, 400))
    v := &View{
        App:          a,
        Window:       w,
        MineButton:   mine,
        SearchButton: queryB,
        SearchBar:    bar,
        SongList:     songList,
        Songs:        &songs,
    }
    return v
}

func CreateSongList(songs *[]ViewSong, buttonHandler func(song ViewSong)) *widget.List {
    songList := widget.NewList(
        func() int {
            return len(*songs)
        },
        func() fyne.CanvasObject {
            return container.NewHBox(
                widget.NewButton("▶️", func() {
                }),
                widget.NewLabel("Nombre de la canción"),
                widget.NewLabel("Artista"),
                widget.NewLabel("Álbum"),
            )
        },
        func(id widget.ListItemID, obj fyne.CanvasObject) {
            // Actualizar los elementos de la lista con los datos de cada canción
            hbox := obj.(*fyne.Container)
            playButton := hbox.Objects[0].(*widget.Button)
            hbox.Objects[1].(*widget.Label).SetText((*songs)[id].Title)
            hbox.Objects[2].(*widget.Label).SetText((*songs)[id].Performer)
            hbox.Objects[3].(*widget.Label).SetText((*songs)[id].Album)

            // Actualizar la acción del botón de reproducir
            playButton.OnTapped = func() {
                buttonHandler((*songs)[id])
            }
        },
    )
    return songList
}

func (v *View) RunView() {
    v.Window.ShowAndRun()
}

func (v *View) ShowBar() {
    v.ProgressWin.Show()
}

func (v *View) DisableElements() {
    v.SearchBar.Disable()
    v.MineButton.Disable()
    v.SearchButton.Disable()
}

func (v *View) EnableElements() {
    v.SearchBar.Enable()
    v.MineButton.Enable()
    v.SearchButton.Enable()
}

func (v *View) MakeProgressBar() {
    w := v.App.NewWindow("Minando...")
    progress := widget.NewProgressBar()
    v.ProgressBar = progress
    progress.Max = 100
    w.Resize(fyne.NewSize(400, 50))
    progressContainer := container.NewVBox(progress)
    w.SetContent(progressContainer)
    v.ProgressWin = w
}

func (v *View) UpdateProgressBar(progress float64) {
    v.ProgressBar.SetValue(progress)
}
