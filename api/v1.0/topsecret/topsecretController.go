package topsecret

import (
	"github.com/gin-gonic/gin"
)

func topSecret(c *gin.Context) {

	var r RequestBody
	var response TopSecretResponse

	if err := c.BindJSON(&r); err != nil {
		c.AbortWithStatus(400)
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
