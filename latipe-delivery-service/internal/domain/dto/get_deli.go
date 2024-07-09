package dto

import "time"

type GetDeliveryByIdRequest struct {
	DeliveryId string `params:"id"`
}

type GetDeliveryByIdResponse struct {
	ID           string    `json:"_id" `
	DeliveryName string    `json:"delivery_name,omitempty"`
	DeliveryCode string    `json:"delivery_code,omitempty"`
	BaseCost     int       `json:"base_cost,omitempty" `
	Description  string    `json:"description,omitempty" `
	CreatedAt    time.Time `json:"created_at,omitempty"`
	IsActive     bool      `json:"is_active"`
}
