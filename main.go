package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/book/", BookHandler)
	// Handlers serão adicionados aqui

	log.Println("Servidor iniciado na porta 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
