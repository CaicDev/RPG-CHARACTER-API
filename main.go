package main

import (
	"github.com/gin-gonic/gin"
	"rpg-api.com/m/controller"
	"rpg-api.com/m/service"
)

var (
	characterService    service.CharacterService       = service.New()
	characterController controller.CharacterController = controller.New(characterService)
)

func main() {
	server := gin.Default()

	apiGroup := server.Group("/api")
	{
		apiGroup.GET("/character", characterController.FindAll)
		apiGroup.GET("/character/:id", characterController.FindById)
		apiGroup.POST("/character", characterController.Save)
	}

	server.Run("127.0.0.1:8080")
}
