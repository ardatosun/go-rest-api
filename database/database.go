package database

import (
	"github.com/ardatosun/go-rest-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var session *gorm.DB

func InitDatabase() {
	var err error
	session, err = gorm.Open("postgres", "postgres://leqfncon:AIeMXzkHbhlilL7vnbc6sC8hUe2UW4sR@kandula.db.elephantsql.com:5432/leqfncon")
	if err != nil {
		panic("Failed to connect to the database")
	}
	log.Print("Successfully connected to the database")
	session.AutoMigrate(&models.User{})
}

func GetSession() *gorm.DB {
	if session.DB().Ping() == nil {
		return session
	}
	InitDatabase()
	return session
}