package main

import (
	"server_pc/db"
	"server_pc/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start("127.0.0.1:3000"))
}
