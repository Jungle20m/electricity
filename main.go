package main

import (
	"context"
	"fmt"
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
	WithTransaction(ctx context.Context, fn func(*storage.Storage) error) error
}

func main() {
	db, err := gorm.Open(
		mysql.Open("anhnv:anhnv!@#456@tcp(1.53.252.177:3306)/healthnet?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Fatal(err)
	}

	stor := storage.NewStorage(db)

	service := model.ServiceModel{
		Title:      "service title",
		SubTitle:   "service subtitle",
		Body:       "service body",
		Image:      "service image",
		CategoryID: 0,
		Active:     0,
	}

	category := model.CategoryModel{
		Name:        "category name",
		Description: "category description",
		Active:      0,
	}

	err = stor.WithTransaction(context.Background(), func(s *storage.Storage) error {
		if err := s.InsertService(context.Background(), service); err != nil {
			fmt.Printf("error to insert service: %+v\n", err)
			return err
		}
		if err := s.InsertCategory(context.Background(), category); err != nil {
			fmt.Printf("error to insert category: %+v\n", err)
			return err
		}
		return fmt.Errorf("some error")
	})

	fmt.Println(err)

}
