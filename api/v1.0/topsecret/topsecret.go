package topsecret

import (
	"github.com/gin-gonic/gin"
	"github.com/lmartinezsch/operacion-fuego-quasar/lib/middlewares"
	"github.com/lmartinezsch/operacion-fuego-quasar/services"
	"github.com/lmartinezsch/operacion-fuego-quasar/services/location"
	"github.com/lmartinezsch/operacion-fuego-quasar/services/message"
)

var messageService message.Service
var locationService location.Service

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {

	messageRegisteredService, _ := services.GetService(message.ServiceName)
	messageService = messageRegisteredService.(message.Service)

	locationRegisteredService, _ := services.GetService(location.ServiceName)
	locationService = locationRegisteredService.(location.Service)

	posts := r.Group("/topsecret")
	{
		posts.POST("/", middlewares.Authorized, topSecret)
	}
}

type SatellitestRequest struct {
	Name     string   `json:"name" binding:"required"`
	Distance float32  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}

type RequestBody struct {
	Satellites []SatellitestRequest `json:"satellites" binding:"required"`
}

type TopSecretResponse struct {
	Position struct {
		X float32 `json:"x"`
		Y float32 `json:"y"`
	} `json:"position"`
	Message string `json:"message"`
}
