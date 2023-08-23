package main

import (
	"fmt"

	"github.com/PedroAndriola/go_REST_API/database"
	"github.com/PedroAndriola/go_REST_API/models"
	"github.com/PedroAndriola/go_REST_API/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o serviço Rest com go")
	routes.HandleRequest()
}
