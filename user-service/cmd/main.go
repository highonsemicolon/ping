package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	address := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	log.Printf("Starting user service on %s ...", address)
	log.Fatal(router.Run(address))
}
