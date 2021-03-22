package api

import (
	"github.com/gin-gonic/gin"
	apiv1 "github.com/lmartinezsch/operacion-fuego-quasar/api/v1.0"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		apiv1.ApplyRoutes(api)
	}
}
