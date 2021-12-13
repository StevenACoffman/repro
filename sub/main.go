package main

import (
	"fmt"
	"net/http"
)

// 	Populated during build
var (
	AppName = "unknown" // What the app is called
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s, %s!", AppName, r.URL.Path[1:])
}
