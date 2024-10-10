package main

import(
    "fmt"
    "github.com/CherubRocky/Musical/controlador"
)

func main() {
    err := controlador.Run()
    if err != nil {
        fmt.Println(fmt.Sprintf("Hubo une error: %v", err))
    }
    // db, err := modelo.NewMusicalDB()
    // defer db.Close()
    // if err = db.DB.Ping(); err != nil {
    //     fmt.Println(fmt.Sprintf("Error ping lol: %", err))
    // }
    // if err != nil {
    //     fmt.Println(fmt.Sprintf("Error: %", err))
    //     return
    // }
    // controlador.Mine("~/Documentos/Prueba/", db)
}
