package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/health", Health)
	log.Println("Starting Server")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln("Error")
	}

}

func Health(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending health ok")
	w.WriteHeader(http.StatusOK)
}
