package main

import (
	"strail/db"
	"strail/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
