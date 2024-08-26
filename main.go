package main

import (
	"GoTutorial/handlers"
	"GoTutorial/pokemon/dao"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"log"
)

func main() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "gotutorialspace"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	dao.InitializeDB(session)

	router := gin.Default()
	router.GET("/pokemon/pokedex/:name", handlers.GetPokedex)
	router.GET("/pokemon/:name", handlers.GetPokemonByName)
	router.PATCH("/pokemon/health/increase/:name", handlers.IncrementPokemonHealth)
	router.POST("/pokemon/pokedex", handlers.AddPokemon)

	//TODO error handling?
	err := router.Run("localhost:9090")
	if err != nil {
		return
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
