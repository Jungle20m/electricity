package main

import (
	"fmt"
	"github.com/Jungle20m/electricity/config"
	"log"

	"github.com/Jungle20m/electricity/sdk/mysql"
)

func init() {
	if err := config.LoadConfig("config.yaml"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	dns := "anhnv:anhnv!@#456@tcp(1.53.252.177:3306)/healthnet?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := mysql.New(dns)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

}
