package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HelloHandler Golang")
}
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logger: Got server version request !")
	fmt.Fprint(w, "The goland micro service: v0.0.1")
}
