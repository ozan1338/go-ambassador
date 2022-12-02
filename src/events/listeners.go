package events

import (
	"context"
	"encoding/json"
	"go-ambassador/src/database"
	"go-ambassador/src/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Listen(message *kafka.Message) error {
	key := string(message.Key)
	
	switch key {
	case "product_created":
		return ProductCreated(message.Value)
	case "product_updated":
		return ProductUpdated(message.Value)
	case "product_deleted":
		return ProductDeleted(message.Value)
	case "order_created":
		return OrderCreated(message.Value)
	}

	return nil
}

func ProductCreated(value []byte) error {
	var product models.Product

	if err := json.Unmarshal(value, &product); err != nil {
		return err
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return err
	}
	go database.ClearCache("products_frontend", "products_backend")

	return nil
}

func ProductUpdated(value []byte) error {
	var product models.Product

	// json.Unmarshal(value, &product)
	if err := json.Unmarshal(value, &product); err != nil {
		return err
	}

	
	if err := database.DB.Model(&product).Updates(&product).Error; err != nil {
		return err
	}
	go database.ClearCache("products_frontend", "products_backend")

	return nil
}

func ProductDeleted(value []byte) error {
	var id uint

	// json.Unmarshal(value, &id)
	if err := json.Unmarshal(value, &id); err != nil {
		return err
	}

	if err := database.DB.Delete(&models.Product{}, "id = ?", id).Error; err != nil {
		return err
	}
	
	go database.ClearCache("products_frontend", "products_backend")

	return nil
}

type Order struct {
	Id uint `json:"id"`
	UserId          uint        `json:"user_id"`
	Code            string      `json:"code"`
	AmbassadorRevenue float64 `json:"ambassador_revenue"`
	AmbassadorName string `json:"ambassador_name"`
}


func OrderCreated(value []byte) error {
	var order Order

	if err := json.Unmarshal(value, &order); err != nil {
		return err
	}

	if err := database.DB.Create(&models.Order{
		Id: order.Id,
		UserId: order.UserId,
		Code: order.Code,
		Total: order.AmbassadorRevenue,
	}).Error; err != nil {
		return err
	}
	
	database.Cache.ZIncrBy(context.Background(), "rankings",  order.AmbassadorRevenue, order.AmbassadorName)

	return nil
}