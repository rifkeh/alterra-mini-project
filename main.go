package main

import (
	"miniproject/config"
	"miniproject/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}
