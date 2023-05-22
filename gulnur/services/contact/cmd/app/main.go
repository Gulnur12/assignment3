package main

import (
    "fmt"
    "gulnur.net/pkg/store/postgres"
)

func main() {
    // Параметры подключения к БД
    host := "localhost"
    port := "5432"
    user := "postgres"
    password := "toha2004"
    dbname := "example"

    // Подключение к БД PostgreSQL
    db, err := postgres.ConnectDB(host, port, user, password, dbname)
    if err != nil {
        fmt.Println("Failed to connect to the database:", err)
        return
    } else{
        fmt.Println("ok")
    }

    defer db.Close()

    // Ваши дальнейшие действия с БД...
}
