package services

import (
	"fmt"
	"log"
	"os"

	"github.com/demo-go-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ItemService struct {
	database gorm.DB
}

func NewItemService() (*ItemService, error) {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: psqlInfo,
	}))

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if err = db.AutoMigrate(&models.Item{}); err != nil {
		log.Println(err)
		return nil, err
	}

	return &ItemService{database: *db}, nil
}

func (is *ItemService) GetAll() models.Items {
	var items models.Items
	is.database.Find(&items)
	return items
}

func (is *ItemService) GetById(id string) (*models.Item, error) {
	var item models.Item
	if err := is.database.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (is *ItemService) CreateItem(item models.Item) {
	is.database.Create(&item)
}

func (is *ItemService) UpdateItem(item models.Item) {
	is.database.Save(&item)
}

func (is *ItemService) DeleteItem(id string) {
	var item models.Item
	is.database.Where("id = ?", id).Delete(&item)
}
