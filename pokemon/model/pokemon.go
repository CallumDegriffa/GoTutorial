package model

import "fmt"

type Pokemon struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Health int    `json:"health"`
}

type PokemonMaster struct {
	Name string
	Rank string
}

type Pokedex struct {
	PokemonList   []Pokemon
	PokemonMaster PokemonMaster
}

type PokemonError struct {
	Message string
}

func (e PokemonError) Error() string {
	return fmt.Sprintf("Pokemon error occured: %s", e.Message)
}

const MaxHealth = 100
