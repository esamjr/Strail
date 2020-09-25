package main

import (
	"log"
	"os"
	"strail/db"
	"strail/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":" + port))
}
