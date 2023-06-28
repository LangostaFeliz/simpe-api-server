package main

import (
	"fmt"
	"log"
	"net/http"

	"example/api-server-1/api"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome home!")
}

func main() {
	fmt.Println("starting a golang api server in http://localhost:8080")
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
