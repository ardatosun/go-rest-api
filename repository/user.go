package repository

import (
	"github.com/ardatosun/go-rest-api/database"
	"github.com/ardatosun/go-rest-api/models"
)

func FindAll() []models.User {
	session := database.GetSession()
	var users []models.User
	session.Find(&users)
	return users
}

func FindById(id string) models.User {
	session := database.GetSession()
	var user models.User
	session.Find(&user, id)
	return user
}

func Save(user *models.User) *models.User {
	session := database.GetSession()
	session.Create(&user)
	return user
}