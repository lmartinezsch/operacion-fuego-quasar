package topsecret

import (
	"github.com/gin-gonic/gin"
)

// TopSecret godoc
// @Summary Get Top Secret
// @Description get position and message of ship
// @Accept  json
// @Produce  json
// @Success 200 {object} TopSecretResponse
// @Header 200 {string} Token "qwerty"
// @Failure 404
// @Failure 500
// @Failure default
// @Router /topsecret [post]
func topSecret(c *gin.Context) {

	var r RequestBody
	var response TopSecretResponse

	if err := c.BindJSON(&r); err != nil {
		c.AbortWithStatus(404)
		return
	}

	x, y := locationService.GetLocation(r.Satellites[0].Distance, r.Satellites[1].Distance, r.Satellites[2].Distance)

	//TODO: se debe pasar de otra forma los parámetros por si llegan a ser más salites
	message := messageService.GetMessage(r.Satellites[0].Message, r.Satellites[1].Message, r.Satellites[2].Message)

	response.Position.X = x
	response.Position.Y = y
	response.Message = message

	c.JSON(200, response)
}
