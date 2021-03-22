package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/lmartinezsch/operacion-fuego-quasar/api/v1.0/auth"
	"github.com/lmartinezsch/operacion-fuego-quasar/api/v1.0/posts"
	"github.com/lmartinezsch/operacion-fuego-quasar/api/v1.0/satellite"
	"github.com/lmartinezsch/operacion-fuego-quasar/api/v1.0/topsecret"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)
		auth.ApplyRoutes(v1)
		posts.ApplyRoutes(v1)
		satellite.ApplyRoutes(v1)
		topsecret.ApplyRoutes(v1)
	}
}
