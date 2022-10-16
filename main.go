package main

import (
	"github.com/madfelps/api-gin-go/database"
	"github.com/madfelps/api-gin-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
