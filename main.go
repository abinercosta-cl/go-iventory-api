package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Funçoes Handler (Controladores)

func createProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Endpoint para criar Produtos!"})
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Endpoint para criar Produtos!"})
}

func getProductsByID(c *gin.Context) {
	//Pegando o ID do produto da rota
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Endpoint para obter o produto com ID: " + id})
}

// Função Principal do servidor
func main() {
	//cria um roteador padrão com gin
	router := gin.Default()

	//Agrupando as Rotas de produtos
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("/", createProduct)
		productRoutes.GET("/", getProducts)
		productRoutes.GET("/:id", getProductsByID)

	}

	//Roda o servidor na porta 8080
	router.Run(":8080")
}
