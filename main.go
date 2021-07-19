package main

import (
	"time"
)

func init() {
	time.Local = time.FixedZone("JST", 9*60*60)
}

func main() {
	println("Hello Amadeus!")
}
