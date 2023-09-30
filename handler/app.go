package handler

import (
	"h8-assignment-2/infra/database"
	"h8-assignment-2/repository/item_repository/item_pg"
	"h8-assignment-2/repository/order_repository/order_pg"
	"h8-assignment-2/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	orderRepo := order_pg.NewOrderPG(db)

	itemRepo := item_pg.NewItemPG(db)

	orderService := service.NewOrderService(orderRepo, itemRepo)

	orderHandler := NewOrderHandler(orderService)

	r := gin.Default()

	// ROUTES
	r.GET("/orders", orderHandler.GetOrders)
	r.POST("/orders", orderHandler.CreateOrder)
	r.PUT("/orders/:orderId", orderHandler.UpdateOrder)
	r.DELETE("/orders/:orderId", orderHandler.DeleteOrder)

	r.Run(":8080")
}
