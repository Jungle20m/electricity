package main

import (
	"context"
	"github.com/Jungle20m/electricity/config"
	"github.com/Jungle20m/electricity/modules/api/model"
	"github.com/Jungle20m/electricity/modules/api/storage"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func init() {
	if err := config.LoadConfig("config.yaml"); err != nil {
		log.Fatal(err)
	}
}

type StorageInterface interface {
	InsertService(ctx context.Context, service model.ServiceModel) error
	InsertCategory(ctx context.Context, category model.CategoryModel) error
}

func main() {
	db, err := gorm.Open(
		mysql.Open("anhnv:anhnv!@#456@tcp(1.53.252.177:3306)/healthnet?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Fatal(err)
	}

	sto := storage.NewStorage(db)

	service := model.ServiceModel{
		Title:      "1",
		SubTitle:   "1",
		Body:       "1",
		Image:      "1",
		CategoryID: 0,
		Active:     0,
	}

	//category := model.CategoryModel{
	//	Name:        "",
	//	Description: "",
	//	Active:      0,
	//}

	sto.InsertService(context.Background(), service)

}
