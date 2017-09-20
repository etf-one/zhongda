package main

import (
	"flag"
)

func init() {
	flag.Parse()
}

func main() {
	app := App{}
	app.initialize()
	app.Run(":3001")
}
