package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gtihub.com/franciscoklaus/golango-simple-crud/servidor"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)
	router.HandleFunc("/contagem", servidor.ContarUsuario).Methods(http.MethodGet)

	fmt.Println("Escutando na porta 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))

}
