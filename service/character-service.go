package service

import (
	"errors"

	"rpg-api.com/m/entity"
)

type CharacterService interface {
	FindById(uint32) (entity.Character, error)
	FindAll() []entity.Character
	Save(entity.Character) entity.Character
}

type characterService struct {
	characters []entity.Character
}

func New() CharacterService {
	return &characterService{}
}

func (service *characterService) FindById(id uint32) (entity.Character, error) {
	for _, character := range service.characters {
		if character.ID == id {
			return character, nil
		}
	}
	return entity.Character{}, errors.New("Cannot find user with this id")

}

func (service *characterService) FindAll() []entity.Character {
	return service.characters
}

func (service *characterService) Save(character entity.Character) entity.Character {
	service.characters = append(service.characters, character)

	return character
}
