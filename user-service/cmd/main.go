package main

import (
	"log"

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
}
