package main

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func dbCreate() {
    dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    db.Exec("CREATE DATABASE shoping_list")
    fmt.Println("Database created successfully!")
}
