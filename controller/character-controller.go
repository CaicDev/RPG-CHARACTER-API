package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"rpg-api.com/m/entity"
	"rpg-api.com/m/service"
)

type CharacterController struct {
	service *service.CharacterService
}

func New(service *service.CharacterService) *CharacterController {
	return &CharacterController{
		service: service,
	}
}

func (controller *CharacterController) FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	uid := uint32(id)

	character, err := controller.service.FindById(uid)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(200, character)
}

func (controller *CharacterController) FindAll(ctx *gin.Context) {
	ctx.JSON(200, controller.service.FindAll())
}

func (controller *CharacterController) Save(ctx *gin.Context) {
	newCharacter := entity.Character{}
	ctx.BindJSON(&newCharacter)

	controller.service.Save(newCharacter)

	ctx.JSON(200, newCharacter)
}
