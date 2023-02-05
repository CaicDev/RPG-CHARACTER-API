package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"rpg-api.com/m/controller"
	"rpg-api.com/m/service"
)


const PORT = "8080"

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Cannot loading the .env file" + err.Error())
	}

	db, err := gorm.Open(sqlite.Open("rpg-character.db"), &gorm.Config{})

	if err != nil {
		panic("cannot create database " + err.Error())
	}

	var (
		characterService    service.CharacterService       = service.New(db)
		characterController controller.CharacterController = controller.New(characterService)
	)

	admin, password := os.Getenv("ADMIN"), os.Getenv("PASSWORD")

	server := gin.Default()

	server.Use(gin.BasicAuth(gin.Accounts{
		admin: password,
	}))

	apiGroup := server.Group("/api")
	{
		apiGroup.GET("/character", characterController.FindAll)
		apiGroup.GET("/character/:id", characterController.FindById)
		apiGroup.POST("/character", characterController.Save)
	}

	server.Run(":"+PORT)
}
