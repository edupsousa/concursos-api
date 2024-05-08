package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/edupsousa/concursos-api/cmd/api"
	"github.com/edupsousa/concursos-api/config"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	db := openDBConnection()
	startAPIServer(db)
}

func startAPIServer(db *gorm.DB) {
	server := api.NewAPIServer(fmt.Sprintf("%s:%s", config.Envs.PublicHost, config.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func openDBConnection() *gorm.DB {
	sqlConfig := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 fmt.Sprintf("%s:%s", config.Envs.DBHost, config.Envs.DBPort),
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlConnection, err := sql.Open("mysql", sqlConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = sqlConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to the database!")

	db, err := gorm.Open(gmysql.New(gmysql.Config{Conn: sqlConnection}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
