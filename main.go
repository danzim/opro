package main

import (
	"fmt"
	"log"
	"net/http"

	oapi "github.com/danzim/opro/oapi"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/version", oapi.GetVersion).Methods("GET")
	router.HandleFunc("/api/projects", oapi.CreateProject).Methods("POST")
	router.HandleFunc("/api/projects", oapi.GetAllProjects).Methods("GET")
	router.HandleFunc("/api/projects/{ci}", oapi.GetOneProject).Methods("GET")
	router.HandleFunc("/api/projects/{ci}", oapi.UpdateProject).Methods("PATCH")
	router.HandleFunc("/api/projects/{ci}", oapi.DeleteProject).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
