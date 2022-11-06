package main

import (
	"github.com/marcelo-antonio-nunes/api-go-rest.git/database"
	"github.com/marcelo-antonio-nunes/api-go-rest.git/routes"
)

func main() {
	database.ConectarComBancoDeDados()
	routes.HandleResquest()

}
