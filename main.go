package main

import (
	"ybmecc/db"
	"ybmecc/route"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
