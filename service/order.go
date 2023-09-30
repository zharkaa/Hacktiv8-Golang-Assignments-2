package service

import (
	"fmt"
	"h8-assignment-2/dto"
	"h8-assignment-2/entity"
	"h8-assignment-2/pkg/errs"
	"h8-assignment-2/repository/item_repository"
	"h8-assignment-2/repository/order_repository"
	"net/http"
)

type orderService struct {
	OrderRepo order_repository.Repository
	ItemRepo  item_repository.Repository
}

type OrderService interface {
	CreateOrder(newOrderRequest dto.NewOrderRequest) (*dto.NewOrderResponse, errs.Error)
	GetOrders() (*dto.GetOrdersResponse, errs.Error)
	UpdateOrder(orderId int, newOrderRequest dto.NewOrderRequest) (*dto.NewOrderResponse, errs.Error)
	DeleteOrder(orderId int) (*dto.NewOrderResponse, errs.Error)
}

func NewOrderService(orderRepo order_repository.Repository, itemRepo item_repository.Repository) OrderService {
	return &orderService{
		OrderRepo: orderRepo,
		ItemRepo:  itemRepo,
	}
}

func (os *orderService) UpdateOrder(orderId int, newOrderRequest dto.NewOrderRequest) (*dto.NewOrderResponse, errs.Error) {

	_, err := os.OrderRepo.ReadOrderById(orderId)

	if err != nil {
		return nil, err
	}

	itemCodes := []any{}

	for _, eachItem := range newOrderRequest.Items {
		itemCodes = append(itemCodes, eachItem.ItemCode)
	}

	items, err := os.ItemRepo.GetItemsByCodes(itemCodes)

	if err != nil {
		return nil, err
	}

	for _, eachItemFromRequest := range newOrderRequest.Items {
		isFound := false

		for _, eachItem := range items {

			if eachItem.OrderId != orderId {

				return nil, errs.NewBadRequest(fmt.Sprintf("item with item code %s doesn't belong to the order with id %d", eachItem.ItemCode, orderId))

			}

			if eachItemFromRequest.ItemCode == eachItem.ItemCode {
				isFound = true
				break
			}
		}

		if !isFound {
			return nil, errs.NewNotFoundError(fmt.Sprintf("item with item code %sis not found", eachItemFromRequest.ItemCode))
		}
	}

	itemPayload := []entity.Item{}

	for _, eachItemFromRequest := range newOrderRequest.Items {
		item := entity.Item{
			ItemCode:    eachItemFromRequest.ItemCode,
			Description: eachItemFromRequest.Description,
			Quantity:    eachItemFromRequest.Quantity,
		}

		itemPayload = append(itemPayload, item)

	}

	orderPayload := entity.Order{
		OrderId:      orderId,
		OrderedAt:    newOrderRequest.OrderedAt,
		CustomerName: newOrderRequest.CustomerName,
	}

	err = os.OrderRepo.UpdateOrder(orderPayload, itemPayload)

	if err != nil {
		return nil, err
	}

	response := dto.NewOrderResponse{
		StatusCode: http.StatusOK,
		Message:    "order successfully updated",
		Data:       nil,
	}

	return &response, nil
}

func (os *orderService) GetOrders() (*dto.GetOrdersResponse, errs.Error) {
	orders, err := os.OrderRepo.ReadOrders()

	if err != nil {
		return nil, err
	}

	orderResult := []dto.OrderWithItems{}

	for _, eachOrder := range orders {
		order := dto.OrderWithItems{
			OrderId:      eachOrder.Order.OrderId,
			CustomerName: eachOrder.Order.CustomerName,
			OrderedAt:    eachOrder.Order.OrderedAt,
			CreatedAt:    eachOrder.Order.CreatedAt,
			UpdatedAt:    eachOrder.Order.UpdatedAt,
			Items:        []dto.GetItemResponse{},
		}

		for _, eachItem := range eachOrder.Items {
			item := dto.GetItemResponse{
				ItemId:      eachItem.ItemId,
				ItemCode:    eachItem.ItemCode,
				Quantity:    eachItem.Quantity,
				Description: eachItem.Description,
				OrderId:     eachItem.OrderId,
				CreatedAt:   eachItem.CreatedAt,
				UpdatedAt:   eachItem.UpdatedAt,
			}

			order.Items = append(order.Items, item)
		}

		fmt.Printf("%+v\n", orderResult)
		orderResult = append(orderResult, order)

	}

	response := dto.GetOrdersResponse{
		StatusCode: http.StatusOK,
		Message:    "orders successfully fetched",
		Data:       orderResult,
	}

	return &response, nil
}

func (os *orderService) CreateOrder(newOrderRequest dto.NewOrderRequest) (*dto.NewOrderResponse, errs.Error) {

	orderPayload := entity.Order{
		OrderedAt:    newOrderRequest.OrderedAt,
		CustomerName: newOrderRequest.CustomerName,
	}

	// fmt.Println(orderPayload)

	itemPayload := []entity.Item{}

	for _, eachItem := range newOrderRequest.Items {
		item := entity.Item{
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
		}

		itemPayload = append(itemPayload, item)

	}
	err := os.OrderRepo.CreateOrder(orderPayload, itemPayload)

	// fmt.Println(orderPayload)
	// fmt.Println(itemPayload)

	if err != nil {
		return nil, err
	}

	response := dto.NewOrderResponse{
		StatusCode: http.StatusCreated,
		Message:    "new order successfully created",
		Data:       nil,
	}

	return &response, nil

}

func (os *orderService) DeleteOrder(orderId int) (*dto.NewOrderResponse, errs.Error) {

	_, err := os.OrderRepo.ReadOrderById(orderId)

	if err != nil {
		return nil, err
	}

	err = os.OrderRepo.DeleteOrder(orderId)

	if err != nil {
		return nil, err
	}

	response := dto.NewOrderResponse{
		StatusCode: http.StatusOK,
		Message:    "order successfully deleted",
		Data:       nil,
	}

	return &response, nil
}
