package main

import (
	"fmt"
	"github.com/Jungle20m/electricity/config"
	"log"
)

func init() {
	if err := config.LoadConfig("config.yaml"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Printf("config: %+v", config.ServiceConfig)
}
