package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//cria um roteador padrão com gin
	router := gin.Default()

	//Primeiro endpoint Gin
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API de estoque com Gin Rodando!",
		})
	})

	//Roda o servidor na porta 8080
	router.Run(":8080")
}
