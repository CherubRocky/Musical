package main

import(
    "fmt"
    "github.com/CherubRocky/Musical/controlador"
    "github.com/CherubRocky/Musical/modelo"
)

func main() {
    db, err := modelo.NewMusicalDB()
    defer db.Close()
    if err = db.DB.Ping(); err != nil {
        fmt.Println(fmt.Sprintf("Error ping lol: %", err))        
    }
    if err != nil {
        fmt.Println(fmt.Sprintf("Error: %", err))
        return
    }
    controlador.Mine("~/Documentos/Prueba/", db)
}
