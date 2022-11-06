package main

import (
	"github.com/marcelo-antonio-nunes/api-go-rest.git/database"
	"github.com/marcelo-antonio-nunes/api-go-rest.git/models"
	"github.com/marcelo-antonio-nunes/api-go-rest.git/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome 1", Historia: "historia 1"},
		{Id: 2, Nome: "nome 2", Historia: "historia 2"},
	}
	database.ConectarComBancoDeDados()
	routes.HandleResquest()

}
