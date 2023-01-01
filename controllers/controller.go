package controllers

import (
	"api/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAnime(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Success!"})

	err := schemas.Create(c.Param("name"), c.Param("description"), c.Param("genre"))
	if err != nil {
		fmt.Printf("Ocorreu um erro ao salvar as informações no banco de dados.")
	}
}

func DeleteAnime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success!"})

	err := schemas.Remove(c.Param("name"))
	if err != nil {
		fmt.Printf("Ocorreu um erro ao Deleter as esse item do banco de dados.")
	}
}
