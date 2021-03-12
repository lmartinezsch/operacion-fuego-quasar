package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/velopert/gin-rest-api/api"
	"github.com/velopert/gin-rest-api/database"
	"github.com/velopert/gin-rest-api/lib/middlewares"
)

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
}
