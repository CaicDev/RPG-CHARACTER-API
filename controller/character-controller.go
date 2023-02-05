package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"rpg-api.com/m/entity"
	"rpg-api.com/m/service"
)

type CharacterController interface {
	FindById(*gin.Context)
	FindAll(*gin.Context)
	Save(*gin.Context)
}

type characterController struct {
	service service.CharacterService
}

func New(service service.CharacterService) CharacterController {
	return &characterController{
		service: service,
	}
}

func (controller *characterController) FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	uid := uint(id)

	character := controller.service.FindById(uid)

	ctx.JSON(200, character)
}

func (controller *characterController) FindAll(ctx *gin.Context) {
	ctx.JSON(200, controller.service.FindAll())
}

func (controller *characterController) Save(ctx *gin.Context) {
	var newCharacter entity.NewCharacter

	ctx.BindJSON(&newCharacter)

	controller.service.Save(newCharacter)

	ctx.JSON(200, newCharacter)
}
