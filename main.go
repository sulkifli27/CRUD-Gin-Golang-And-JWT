package main

import (
	"golang-test/config"
	"golang-test/routes"
	"log"

	"github.com/joho/godotenv"
)
func main() {
    // for load godotenv
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

     // database connection
    db := config.ConnectDataBase()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()

     // router
    r := routes.SetupRouter(db)
    r.Run("localhost:8080")
}