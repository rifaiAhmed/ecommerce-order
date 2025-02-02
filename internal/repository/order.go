package repository

import (
	"context"
	"ecommerce-order/internal/models"

	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func (r *OrderRepo) InsertNewOrder(ctx context.Context, order *models.Order) error {
	return r.DB.Transaction(
		func(tx *gorm.DB) error {
			err := tx.Create(order).Error
			if err != nil {
				return err
			}

			return nil
		},
	)
}

func (r *OrderRepo) UpdateStatusOrder(ctx context.Context, orderID int, status string) error {
	return r.DB.Exec("UPDATE orders SET status = ? WHERE id = ?", status, orderID).Error
}

func (r *OrderRepo) GetOrderDetail(ctx context.Context, orderID int) (models.Order, error) {
	var (
		resp models.Order
		err  error
	)
	err = r.DB.Model(&models.Order{}).Preload("OrderItem").Where("id = ?", orderID).First(&resp).Error
	return resp, err
}

func (r *OrderRepo) GetOrder(ctx context.Context) ([]models.Order, error) {
	var (
		resp []models.Order
		err  error
	)
	err = r.DB.Model(&models.Order{}).Preload("OrderItem").Order("id DESC").Find(&resp).Error
	return resp, err
}
