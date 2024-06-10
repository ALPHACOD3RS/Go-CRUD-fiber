package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil{
        log.Fatalf("failed to connect to db")
    }

    db.AutoMigrate(&User{}, &Post{})

    return db

}