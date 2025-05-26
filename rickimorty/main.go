package main

import (
	"log"
	"net/http"
	"os"
	"prubarickmorti/db"
	"prubarickmorti/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db.LlenarBD()

	router := gin.Default()

	router.StaticFS("/app", http.Dir("./static"))

	baseRouter := router.Group("/api")

	baseRouter.GET("/characterslista", handlers.CharactersLista)
	baseRouter.GET("/characters/:id", handlers.Characters)
	baseRouter.POST("/sync/characters", handlers.SyncCharacters)
	baseRouter.POST("/sync/episodes", handlers.SyncEpisodes)
	baseRouter.DELETE("/characters", handlers.DelCharacters)

	router.Run(":8080")
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
