package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", "0.0.0.0:5533", "[host]:[port]")

func main() {
	flag.Parse()

	root := "."
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}
	fmt.Printf("Serving dir '%s', listening on '%s'...\n", root, *addr)
	log.Fatal(http.ListenAndServe(*addr, http.FileServer(http.Dir(root))))
}
