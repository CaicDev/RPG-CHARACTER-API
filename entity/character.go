package entity

type Character struct {
	id           uint32 `json:"id"`
	name         string `json:"name"`
	description  string `json:"description"`
	strength     uint32 `json:"strength"`
	speed        uint32 `json:"speed"`
	intelligence uint32 `json:"intelligence"`
	xp           uint32 `json:"xp"`
}
