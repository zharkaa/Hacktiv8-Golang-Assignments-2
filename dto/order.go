package dto

import "time"

type NewOrderRequest struct {
	OrderedAt    time.Time        `json:"orderedAt"`
	CustomerName string           `json:"customerName"`
	Items        []NewItemRequest `json:"items"`
}

type GetItemResponse struct {
	ItemId      int       `json:"itemId"`
	ItemCode    string    `json:"itemCode"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	OrderId     int       `json:"orderId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type NewOrderResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type GetOrdersResponse struct {
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	Data       []OrderWithItems `json:"data"`
}
