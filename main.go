package main

import (
	"echo_rest/db"
	"echo_rest/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start("127.0.0.1:3000"))
}
