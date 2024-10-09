package main

import(
    "fmt"
    "github.com/CherubRocky/Musical/controlador"
    "github.com/CherubRocky/Musical/modelo"
)

func main() {
    db, err := modelo.NewMusicalDB()
    if err != nil {
        fmt.Println(fmt.Sprintf("Error: %", err))
        return
    }
    controlador.Mine("~/Documentos/Prueba/", db)
}
