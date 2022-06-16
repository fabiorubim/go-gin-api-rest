package main

import (
	"gin-api-rest/database"
	"gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	// models.Alunos = []models.Aluno{
	// 	{Nome: "Fabio", CPF: "456.798.963.65", RG: "45.568.563-9"},
	// 	{Nome: "Marina", CPF: "456.798.963.65", RG: "45.568.563-9"},
	// }
	routes.HandleRequests()
}
