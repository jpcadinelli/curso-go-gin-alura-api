package main

import (
	"github.com/jpcadinelli/curso-go-gin-alura-api/database"
	"github.com/jpcadinelli/curso-go-gin-alura-api/models"
	"github.com/jpcadinelli/curso-go-gin-alura-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "João Pedro", CPF: "000.000.000-00", RG: "00.000.000-0"},
		{Nome: "Gui Lima", CPF: "111.111.111-11", RG: "11.111.111-1"},
	}
	routes.HandleRequests()
}
