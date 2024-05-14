package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/edupsousa/concursos-api/platform/config"
	"github.com/go-sql-driver/mysql"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func MustConnect() *DB {
	sqlDB := createSqlDatabase()
	gormDB := createGormDatabase(sqlDB)
	return &DB{
		gormDB,
	}
}

func createSqlDatabase() *sql.DB {
	sqlConfig := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 fmt.Sprintf("%s:%s", config.Envs.DBHost, config.Envs.DBPort),
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlDB, err := sql.Open("mysql", sqlConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return sqlDB
}

func createGormDatabase(sqlDB *sql.DB) *gorm.DB {
	db, err := gorm.Open(gmysql.New(gmysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
