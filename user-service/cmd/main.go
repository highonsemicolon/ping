package main

import (
	"log"

	"github.com/highonsemicolon/ping/user-service/pkg/utils"
)

func main() {
	_, err := utils.LoadConfig("./configs/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
}
