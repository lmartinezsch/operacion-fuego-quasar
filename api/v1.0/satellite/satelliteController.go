package satellite

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Create Satellite godoc
// @Summary Create Satellite
// @Description post a satellite
// @Accept  json
// @Produce  json
// @Success 200 {object} SatelliteCreateResponse
// @Header 200 {string} Token "qwerty"
// @Failure 400,404
// @Failure 500
// @Failure default
// @Router /satellites [post]
func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var r RequestBody

	if err := c.BindJSON(&r); err != nil {
		c.AbortWithStatus(400)
		return
	}

	satellite := Satellite{
		Name: strings.ToLower(r.Name),
		Position: Position{
			X: r.Position.X,
			Y: r.Position.Y,
		},
	}
	db.NewRecord(satellite)
	db.Create(&satellite)
	c.JSON(201, "")
}
