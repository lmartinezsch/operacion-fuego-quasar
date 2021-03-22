package satellite

import (
	"github.com/gin-gonic/gin"
	"github.com/lmartinezsch/operacion-fuego-quasar/database/models"
	"github.com/lmartinezsch/operacion-fuego-quasar/lib/common"
	"github.com/lmartinezsch/operacion-fuego-quasar/lib/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/satellites")
	{
		posts.POST("/", middlewares.Authorized, create)
		/*posts.GET("/", list)
		posts.GET("/:id", read)
		posts.DELETE("/:id", middlewares.Authorized, remove)
		posts.PATCH("/:id", middlewares.Authorized, update)*/
	}
}

// Satellite type alias
type Satellite = models.Satellite

// Position type alias
type Position = models.Position

// User type alias
type User = models.User

// JSON type alias
type JSON = common.JSON

type RequestBody struct {
	Name     string   `json:"name" binding:"required"`
	Position Position `json:"position" binding:"required"`
}
