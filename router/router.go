package router

import (
	"api/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

func StartRouter() {

	router := gin.Default()

	//Rotas
	router.GET("anime/:name", controllers.DeleteAnime)
	router.GET("anime/:name/:description/:genre", controllers.GetAnime)

	address := os.Getenv("ADDR")
	port := os.Getenv("PORT")

	router.Run(address + port)
}
