package main

import (
	"log"
	"os"

	_ "gitconnect-backend/docs" // Import Swagger docs
	"gitconnect-backend/config"
	"gitconnect-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GitConnect API
// @version 1.0
// @description API documentation for GitConnect
// @termsOfService http://swagger.io/terms/
// @contact.name Victor Muthomi
// @contact.email victor@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host 0.0.0.0:8080
// @BasePath /api
func main() {
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug"
	}
	gin.SetMode(mode)

	if err := config.ConnectDatabase(); err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}
	log.Println("✅ Database connected successfully.")

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.SetTrustedProxies(nil)

	// Updated CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://gitconnect-frontend.vercel.app",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	routes.AuthRoutes(router)
	routes.PostRoutes(router)
	routes.ProfileRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serverAddr := "0.0.0.0:" + port
	log.Printf("✅ Server running on %s", serverAddr)

	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}

