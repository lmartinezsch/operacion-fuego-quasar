package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lmartinezsch/operacion-fuego-quasar/lib/common"
)

// User data model
type SatelliteContact struct {
	gorm.Model
	Distance    float32
	Message     string `gorm:"type:json"`
	SatelliteID uint
	Satellite   Satellite
}

// Serialize serializes user data
func (s *SatelliteContact) Serialize() common.JSON {
	return common.JSON{
		"id":        s.ID,
		"distance":  s.Distance,
		"Message":   s.Message,
		"Satellite": s.Satellite,
	}
}

func (s *SatelliteContact) Read(m common.JSON) {
	s.ID = uint(m["id"].(float64))
	s.Distance = m["distance"].(float32)
	s.Message = m["message"].(string)
	s.Satellite = m["satellite"].(Satellite)
}
