package service

import (
	"errors"

	"rpg-api.com/m/entity"
)

type CharacterService struct {
	characters []entity.Character
}

func New() *CharacterService {
	return &CharacterService{}
}

func (service *CharacterService) FindById(id uint32) (entity.Character, error) {
	characters_list := service.characters
	for _, character := range characters_list  {
		if character.id == id {
			return character, nil
		}
	}
	return entity.Character{}, errors.New("Cannot find user with this id")

}

func (service *CharacterService) FindAll() []entity.Character {
	return service.characters
}

func (service *CharacterService) Save(character entity.Character) entity.Character {
	service.characters = append(service.characters, character)

	return character
}