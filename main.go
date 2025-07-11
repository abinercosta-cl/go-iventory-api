package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Define um handler para a rota principal "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API de estoque rodando")
	})

	fmt.Println("Servidor rodando na porta 8080")
	//inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}
