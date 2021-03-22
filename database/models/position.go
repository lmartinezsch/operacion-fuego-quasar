package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lmartinezsch/operacion-fuego-quasar/lib/common"
)

// User data model
type Position struct {
	gorm.Model
	X float32
	Y float32
}

// Serialize serializes user data
func (p *Position) Serialize() common.JSON {
	return common.JSON{
		"id": p.ID,
		"x":  p.X,
		"y":  p.Y,
	}
}

func (p *Position) Read(m common.JSON) {
	p.ID = uint(m["id"].(float64))
	p.X = m["x"].(float32)
	p.Y = m["y"].(float32)
}
