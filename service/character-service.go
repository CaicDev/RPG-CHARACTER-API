package service

import (
	"gorm.io/gorm"
	"rpg-api.com/m/entity"
)

type CharacterService interface {
	FindById(uint) entity.Character
	FindAll() []entity.Character
	Save(entity.NewCharacter) entity.Character
	Delete(uint)
}

type characterService struct {
	db *gorm.DB
}

func New(db *gorm.DB) CharacterService {
	return &characterService{
		db,
	}
}

func (service *characterService) FindById(id uint) entity.Character {
	var character entity.Character

	service.db.First(&character, id)

	return character

}

func (service *characterService) FindAll() []entity.Character {
	var characters []entity.Character

	service.db.Find(&characters)

	return characters
}

func (service *characterService) Save(newCharacter entity.NewCharacter) entity.Character {

	character := entity.Character{
		Name: newCharacter.Name, Description: newCharacter.Description,
		Magic: newCharacter.Magic, Strength: newCharacter.Strength,
		Speed: newCharacter.Speed, Intelligence: newCharacter.Intelligence,
	}

	service.db.Create(&character)

	return character
}

func (service *characterService) Delete(id uint) {
	 
	service.db.Delete(&entity.Character{}, id)
}