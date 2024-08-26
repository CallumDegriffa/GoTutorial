package handlers

import (
	"GoTutorial/pokemon/model"
	"GoTutorial/pokemon/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"net/http"
)

func GetPokedex(context *gin.Context) {
	pokedex, err := service.GetPokedex(getNameParam(context))

	if err != nil {
		fmt.Println(fmt.Errorf("could not fetch pokemon %s", err))
		return
	}

	context.IndentedJSON(http.StatusOK, pokedex)
}

func GetPokemonByName(session *gocql.Session, context *gin.Context) {
	pokemon, err := service.FetchPokemonByName(session, getNameParam(context))

	if err != nil {
		fmt.Println(fmt.Errorf("could not fetch pokemon %s", err))
		return
	}

	context.IndentedJSON(http.StatusOK, pokemon)
}

func IncrementPokemonHealth(context *gin.Context) {
	pokemon, err := service.IncrementPokemonHealth(getNameParam(context), context.Query("pokeMasterName"))

	if err != nil {
		fmt.Println(fmt.Errorf("could not increment pokemon %s", err))
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	context.IndentedJSON(http.StatusOK, pokemon)
}

func AddPokemon(context *gin.Context) {
	var newPokemon model.Pokemon

	if err := context.BindJSON(&newPokemon); err != nil {
		fmt.Errorf("invalid pokemon %w", model.PokemonError{Message: "pokemon provided is invalid"})
		return
	}
	name := context.Query("name")

	pokedex, err := service.AddPokemonToPokedex(name, newPokemon)

	if err != nil {
		fmt.Println(fmt.Errorf("could not add to pokedex %s", err))
	}

	context.IndentedJSON(http.StatusCreated, pokedex)
}

func getNameParam(context *gin.Context) string {
	return context.Param("name")
}
