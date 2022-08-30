package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods(http.MethodGet)
	port := 3001
	log.Printf("Server running on: "+"localhost:%v", port)

	log.Println(http.ListenAndServe(":"+strconv.Itoa(port), r))
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "pong")
}
