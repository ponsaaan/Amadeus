package main

import (
	"github.com/labstack/gommon/log"
	"github.com/ponsaaan/Amadeus.git/app/presentation"
	"time"
)

func init() {
	time.Local = time.FixedZone("JST", 9*60*60)
}

func main() {
	println("Hello Amadeus!")
	r := presentation.Router()
	r.Logger.SetLevel(log.DEBUG)
	r.Logger.Fatal(r.Start(":80"))
}
