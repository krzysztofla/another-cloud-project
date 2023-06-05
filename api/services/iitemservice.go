package services

import (
	"github.com/demo-go-api/models"
)

type IItemsService interface {
	GetAll() models.Items
	GetById(id string) (*models.Item, error)
	CreateItem(item models.Item)
	UpdateItem(item models.Item)
	DeleteItem(uid string)
}
