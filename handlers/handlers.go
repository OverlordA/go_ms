package handlers

import (
	"fmt"
	"net/http"
)

func getMessage() string {
	return "Hello golang"
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, getMessage())
}
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logger: Got server version request !")
	fmt.Fprint(w, "The goland micro service: v0.0.1")
}
