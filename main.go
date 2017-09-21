package main

import (
	"flag"
)

var addr *string

func init() {
	addr = flag.String("addr", ":3001", "Host Address. Default - :3001")
	flag.Parse()
}

//func main() {
//	app := App{}
//	app.initialize()
//	app.Run(*addr)
//}
