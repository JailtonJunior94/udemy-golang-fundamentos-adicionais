package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jailtonjunior94/udemy-golang-fundamentos-adicionais/database/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", handlers.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", handlers.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", handlers.BuscarUsuario).Methods(http.MethodGet)

	fmt.Println("🚀 Server is running on http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
