package main

import (
	"GoTutorial/handlers"
	"GoTutorial/pokemon/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"log"
	"os"
)

func getPokemonByName(session *gocql.Session) gin.HandlerFunc {
	return func(context *gin.Context) {
		handlers.GetPokemonByName(session, context)
	}
}

func AddPokemon(session *gocql.Session) gin.HandlerFunc {
	return func(context *gin.Context) {
		handlers.AddPokemon(session, context)
	}
}

func main() {
	session, _ := initializeDB()

	router := gin.Default()
	router.GET("/pokemon/pokedex/:name", handlers.GetPokedex)
	router.GET("/pokemon/:name", getPokemonByName(session))
	router.PATCH("/pokemon/health/increase/:name", handlers.IncrementPokemonHealth)
	router.POST("/pokemon/pokedex", handlers.AddPokemonToPokedex)
	router.POST("/pokemon/addPokemon", AddPokemon(session))

	//TODO error handling?
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func initializeDB() (*gocql.Session, error) {
	cassandraHost := os.Getenv("CASSANDRA_HOST")

	if cassandraHost == "" {
		cassandraHost = "localhost"
	}
	cluster := gocql.NewCluster(cassandraHost)
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	err = session.Query(`CREATE KEYSPACE IF NOT EXISTS gotutorialspace WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`).Exec()

	if err != nil {
		fmt.Println(fmt.Errorf("something went wrong %s", err))
		return nil, err
	}

	cluster.Keyspace = "gotutorialspace"
	session, err = cluster.CreateSession()

	if err != nil {
		fmt.Println(fmt.Errorf("something went wrong %s", err))
		return nil, err
	}

	dao.InitializeDB(session)
	return session, nil
}
