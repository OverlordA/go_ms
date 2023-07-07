package main

import (
	"net/http"
)
import "./handlers"

func main() {

	http.HandleFunc("/", handlers.HelloHandler)

	http.HandleFunc("/version", handlers.VersionHandler)

	http.ListenAndServe(":80", nil)
}
