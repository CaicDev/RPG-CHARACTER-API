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
	Delete(*gin.Context)
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

	character, err := controller.service.FindById(uid)

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(200, character)
}

func (controller *characterController) FindAll(ctx *gin.Context) {

	characters, err := controller.service.FindAll()

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(200, characters)
}

func (controller *characterController) Save(ctx *gin.Context) {
	var newCharacter entity.NewCharacter

	ctx.BindJSON(&newCharacter)

	character, err := controller.service.Save(newCharacter)

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(200, character)
}

func (controller *characterController) Delete(ctx *gin.Context) {

	id, id_err := strconv.Atoi(ctx.Param("id"))

	if id_err != nil {
		ctx.JSON(500, gin.H{
			"error": id_err.Error(),
		})
		return
	}

	uid := uint(id)

	db_err := controller.service.Delete(uid)

	if db_err != nil {
		ctx.JSON(500, gin.H{
			"error": db_err.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"message": "character delete succesfull",
	})
}