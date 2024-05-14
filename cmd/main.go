package main

import (
	"fmt"
	"log"

	"github.com/edupsousa/concursos-api/cmd/api"
	"github.com/edupsousa/concursos-api/platform/config"
	"github.com/edupsousa/concursos-api/platform/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := database.MustConnect()
	startAPIServer(db)
}

func startAPIServer(db *database.DB) {
	server := api.NewAPIServer(fmt.Sprintf("%s:%s", config.Envs.PublicHost, config.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
