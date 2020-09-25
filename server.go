package main

import (
	"os"
	"strail/db"
	"strail/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":" + port))
}
