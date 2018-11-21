// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
package main

import (
	"flag"
	"fmt"
	"net/http"
)

var listen = flag.String("listen", ":8080", "listen address")
var dir    = flag.String("dir", ".", "directory to serve")

func main() {
	flag.Parse()
	fmt.Printf("Web server running. Listening on %q", *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	fmt.Printf("%v\n",err)
}
