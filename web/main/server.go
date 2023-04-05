package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jyouturner/intellidox/web/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/files/generate-upload-url", handlers.GenerateUploadURL).Methods("POST")
	router.HandleFunc("/api/files/{id}/confirm", handlers.ConfirmUpload).Methods("POST")
	router.HandleFunc("/api/files/{id}/status", handlers.GetStatus).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
