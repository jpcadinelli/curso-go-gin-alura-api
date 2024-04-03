package main

import (
	"github.com/jpcadinelli/curso-go-gin-alura-api/database"
	"github.com/jpcadinelli/curso-go-gin-alura-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
