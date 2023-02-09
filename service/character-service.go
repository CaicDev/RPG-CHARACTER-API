package service

import (
	"gorm.io/gorm"
	"rpg-api.com/m/entity"
)

type CharacterService interface {
	FindById(uint) (entity.Character, error)
	FindAll() ([]entity.Character, error)
	Save(entity.NewCharacter) (entity.Character, error)
	Delete(uint) error
}

type characterService struct {
	db *gorm.DB
}

func New(db *gorm.DB) CharacterService {
	return &characterService{
		db,
	}
}

func (service *characterService) FindById(id uint) (entity.Character, error) {
	var character entity.Character

	err := service.db.First(&character, id).Error

	return character, err

}

func (service *characterService) FindAll() ([]entity.Character, error) {
	var characters []entity.Character

	err := service.db.Find(&characters).Error

	return characters, err
}

func (service *characterService) Save(newCharacter entity.NewCharacter) (entity.Character, error) {

	character := entity.Character{
		Name: newCharacter.Name, Description: newCharacter.Description,
		Magic: newCharacter.Magic, Strength: newCharacter.Strength,
		Speed: newCharacter.Speed, Intelligence: newCharacter.Intelligence,
	}

	err := service.db.Create(&character).Error


	return character, err
}

func (service *characterService) Delete(id uint) error {

	err := service.db.Delete(&entity.Character{}, id).Error

	return err
}
