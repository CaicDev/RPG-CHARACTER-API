package entity

type Character struct {
	ID           uint32 `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Magic        uint32 `json:"magic"`
	Strength     uint32 `json:"strength"`
	Speed        uint32 `json:"speed"`
	Intelligence uint32 `json:"intelligence"`
}

type NewCharacter struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Magic        uint32 `json:"magic"`
	Strength     uint32 `json:"strength"`
	Speed        uint32 `json:"speed"`
	Intelligence uint32 `json:"intelligence"`
}
