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

	fmt.Println("Escutando na porta 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))

}
