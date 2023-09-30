package order_repository

import (
	"h8-assignment-2/entity"
)

type OrderItemMapped struct {
	Order entity.Order
	Items []entity.Item
}

type OrderItem struct {
	Order entity.Order
	Item  entity.Item
}

func (oim *OrderItemMapped) HandleMappingOrderWithItems(orderItem []OrderItem) []OrderItemMapped {

	ordersItemsMapped := []OrderItemMapped{}

	for _, eachOrderItem := range orderItem {

		isOrderExist := false

		// Check if order exist in ordersItemsMapped
		for i := range ordersItemsMapped {
			// If exist, append the item to the order
			if eachOrderItem.Order.OrderId == ordersItemsMapped[i].Order.OrderId {
				isOrderExist = true
				ordersItemsMapped[i].Items = append(ordersItemsMapped[i].Items, eachOrderItem.Item)
				break
			}
		}

		// If not exist, create new order and append the item to the order
		if !isOrderExist {
			
			// Create new order
			orderItemMapped := OrderItemMapped{
				Order: eachOrderItem.Order,
			}

			// Append the item to the order
			orderItemMapped.Items = append(orderItemMapped.Items, eachOrderItem.Item)

			// Append the order to ordersItemsMapped
			ordersItemsMapped = append(ordersItemsMapped, orderItemMapped)
		}
	}

	return ordersItemsMapped
}
