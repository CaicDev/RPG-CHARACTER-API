package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"rpg-api.com/m/controller"
	"rpg-api.com/m/service"
)

var (
	characterService    service.CharacterService       = service.New()
	characterController controller.CharacterController = controller.New(characterService)
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Cannot loading the .env file" + err.Error())
	}

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

	server.Run("127.0.0.1:8080")
}
