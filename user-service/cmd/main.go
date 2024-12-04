package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/highonsemicolon/ping/user-service/pkg/api"
	"github.com/highonsemicolon/ping/user-service/pkg/db"
	"github.com/highonsemicolon/ping/user-service/pkg/utils"
)

func main() {
	config, err := utils.LoadConfig("./configs/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dbConn, err := db.InitMySQL(config.MySQL)
	if err != nil {
		log.Fatalf("Error connecting to MySQL %v", err)
	}
	defer dbConn.Close()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router := gin.Default()
	router.Use(cors.New(corsConfig))
	api.SetupRoutes(router, dbConn, config.Auth0)

	address := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	log.Printf("Starting user service on %s ...", address)
	log.Fatal(router.Run(address))
}
