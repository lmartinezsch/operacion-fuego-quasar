package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lmartinezsch/operacion-fuego-quasar/lib/common"
)

// User data model
type Satellite struct {
	gorm.Model
	Name       string
	PositionID uint
	Position   Position
}

// Serialize serializes user data
func (s *Satellite) Serialize() common.JSON {
	return common.JSON{
		"id":       s.ID,
		"name":     s.Name,
		"position": s.Position,
	}
}

func (s *Satellite) Read(m common.JSON) {
	s.ID = uint(m["id"].(float64))
	s.Name = m["name"].(string)
	s.Position = m["position"].(Position)
}
