package dao

import (
	"GoTutorial/pokemon/model"
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

func InitializeDB(session *gocql.Session) {
	err = session.Query(`CREATE KEYSPACE IF NOT EXISTS gotutorialspace WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`).Exec()
	if err != nil {
		log.Fatal(err)
	}

	//poke_master_name is the partition key meaning each unique entry in this field/column will be in its own partition
	//we have then used pokemon_name as the clustering column to order by pokemon name
	err = session.Query(`CREATE TABLE IF NOT EXISTS goTutorialSpace.pokedex (poke_master_name TEXT, poke_master_rank TEXT, pokemon_name TEXT, pokemon_weight INT, pokemon_health INT, PRIMARY KEY (poke_master_name, pokemon_name))`).Exec()

	if err != nil {
		log.Fatal(err)
	}

	id := gocql.TimeUUID()
	if err := session.Query(`INSERT INTO goTutorialSpace.pokedex (poke_master_name, poke_master_rank, pokemon_name) VALUES (?, ?, ?)`, "Ash", "Expert", "ekans", 30, 50).Exec(); err != nil {
		log.Fatal(err)
	}

	err = session.Query(`CREATE TABLE IF NOT EXISTS goTutorialSpace.pokemon (pokemon_name TEXT, poke_master_name, pokemon_weight INT, pokemon_health INT, PRIMARY KEY (pokemon_name))`).Exec()

	if err != nil {
		log.Fatal(err)
	}

	id := gocql.TimeUUID()
	if err := session.Query(`INSERT INTO goTutorialSpace.pokemon (pokemon_name, poke_master_name, pokemon_weight, pokemon_health) VALUES (?, ?, ?)`, "ekans", "Ash", 30, 50).Exec(); err != nil {
		log.Fatal(err)
	}
}

func FetchPokemonFromDB(session *gocql.Session, pokemonName string) model.Pokemon {
	cluster := gocql.NewCluster("127.0.0.1")
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	var pokemon model.Pokemon
	if err := session.Query(`SELECT pokemon_name, pokemon_health, pokemon_weight FROM goTutorialSpace.pokemon WHERE pokemon_name = ?`, pokemonName).Consistency(gocql.One).Scan(&pokemon.Name, &pokemon.Health, &pokemon.Weight); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Retrieved pokemon: Name=%s", pokemon.Name)

	return pokemon
}
