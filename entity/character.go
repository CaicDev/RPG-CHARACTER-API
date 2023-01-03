package entity

type Character struct {
	ID           uint32 `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Magic        uint32 `json:"magic" binding:"required"`
	Strength     uint32 `json:"strength" binding:"required"`
	Speed        uint32 `json:"speed" binding:"required"`
	Intelligence uint32 `json:"intelligence" binding:"required"`
}

type NewCharacter struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Magic        uint32 `json:"magic" binding:"required"`
	Strength     uint32 `json:"strength" binding:"required"`
	Speed        uint32 `json:"speed" binding:"required"`
	Intelligence uint32 `json:"intelligence" binding:"required"`
}
