package main

import (
	"fmt"
	"github.com/Jungle20m/electricity/config"
	"github.com/Jungle20m/electricity/sdk/mysql"
	"log"
	"time"
)

func init() {
	if err := config.LoadConfig("config.yaml"); err != nil {
		log.Fatal(err)
	}
}

type Data struct {
	string `json:"data"`
}

func main() {
	dns := "anhnv:anhnv!@#456@tcp(1.53.252.177:3306)/healthnet?charset=utf8mb4&parseTime=True&loc=Local"

	msql, err := mysql.New(
		dns,
		mysql.WithConnectionMaxLifetime(time.Hour),
		mysql.WithMaxIdleConnection(10),
		mysql.WithMaxOpenConnection(100))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("mysql: %+v", msql)
}
