package services

import (
	"github.com/ardatosun/go-rest-api/models"
	"github.com/ardatosun/go-rest-api/repository"
	"github.com/gofiber/fiber"
	"log"
	"strings"
)

func FindAll(ctx *fiber.Ctx) {
	users := repository.FindAll()
	log.Printf("%d records were found.", len(users))
	ctx.JSON(users)
}

func FindById(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	user := repository.FindById(id)
	if isBlank(user.Firstname) {
		log.Printf("No record with id: %s was found.", id)
	}
	ctx.JSON(user)
}

func Save(ctx *fiber.Ctx) {
	user := new(models.User)
	var err error
	err = ctx.BodyParser(user)
	if err != nil {
		log.Print("Error on save.")
	}
	repository.Save(user)
	log.Print("User record saved.")
	ctx.JSON(user)
}

func isBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
