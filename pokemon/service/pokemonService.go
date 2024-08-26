package service

import (
	"GoTutorial/pokemon/dao"
	"GoTutorial/pokemon/model"
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"io/ioutil"
	"net/http"
)

var pokedexList = []model.Pokedex{{
	PokemonList:   []model.Pokemon{{Name: "ekans", Weight: 60, Health: 95}, {Name: "bedrill", Weight: 20, Health: 80}},
	PokemonMaster: model.PokemonMaster{Name: "Ash", Rank: "Intermediate"},
}}

func FetchPokemonByName(session *gocql.Session, name string) (*model.Pokemon, error) {
	return fetchPokemonByName(session, name)
}

func IncrementPokemonHealth(pokemonName string, pokeMasterName string) (*model.Pokemon, error) {
	pokedex, err := GetPokedex(pokeMasterName)

	if err != nil {
		fmt.Println(fmt.Errorf("could not fetch pokedex %s", err))
		return nil, err
	}

	var pokemon *model.Pokemon

	for i, currentPokemon := range pokedex.PokemonList {
		if currentPokemon.Name == pokemonName {
			pokemon = &pokedex.PokemonList[i]
		}
	}

	if pokemon == nil {
		return nil, model.PokemonError{Message: fmt.Sprintf("pokemon %s does not exist for pokemaster %s", pokemonName, pokeMasterName)}
	}

	if pokemon.Health < model.MaxHealth {
		pokemon.Health++
	}

	return pokemon, nil
}

func GetPokedex(name string) (*model.Pokedex, error) {
	for i, currentPokedex := range pokedexList {
		if currentPokedex.PokemonMaster.Name == name {
			return &pokedexList[i], nil
		}
	}

	return nil, &model.PokemonError{Message: fmt.Sprintf("could not find pokedex for master: %s", name)}
}

func AddPokemonToPokedex(name string, pokemon model.Pokemon) (*model.Pokedex, error) {
	pokedex, err := GetPokedex(name)

	if err != nil {
		return nil, err
	}

	pokedex.PokemonList = append(pokedex.PokemonList, pokemon)

	return pokedex, nil
}

func AddPokemon(session *gocql.Session, pokemon *model.Pokemon) error {
	err := dao.AddPokemon(session, pokemon)

	if err != nil {
		return err
	}

	return nil
}

func fetchPokemonByName(session *gocql.Session, name string) (*model.Pokemon, error) {
	var pokemon model.Pokemon

	pokemon = dao.FetchPokemon(session, name)
	//fetch from cassandra, only fetch API if it doesn't currently exist

	if &pokemon == nil {
		url := "https://pokeapi.co/api/v2/pokemon"

		// Perform the GET request
		resp, err := http.Get(url + "/" + name)
		if err != nil {
			fmt.Println("fetch error:", err)
			return nil, err
		}
		// Make sure to close the response body when done
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return nil, err
		}
	}

	return &pokemon, nil
}
