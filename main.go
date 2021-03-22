package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/lmartinezsch/operacion-fuego-quasar/api"
	"github.com/lmartinezsch/operacion-fuego-quasar/database"
	"github.com/lmartinezsch/operacion-fuego-quasar/lib/middlewares"
	"github.com/lmartinezsch/operacion-fuego-quasar/services"
	"github.com/lmartinezsch/operacion-fuego-quasar/services/location"
	"github.com/lmartinezsch/operacion-fuego-quasar/services/message"
)

// @title Operación Fuego Quasar API
// @version 1.0
// @description Obtener la posición de una nave y su mensaje a partir de las distancias con los satelites

// @contact.name Leandro Martinez
// @contact.email leandro.martinez01@gmail.com

// @host localhost:4000
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initializes database
	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	app := gin.Default() // create gin app

	registerServices(db)

	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Origin", "Accept", "X-Requested-With", "Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api.ApplyRoutes(app) // apply api router
	app.Run(":" + port)  // listen to given port

	defer services.DeregisterAll()
}

func registerServices(db *gorm.DB) {

	locationService := location.NewService(db)
	services.RegisterService(location.ServiceName, locationService)

	messageService := message.NewService()
	services.RegisterService(message.ServiceName, messageService)

}
