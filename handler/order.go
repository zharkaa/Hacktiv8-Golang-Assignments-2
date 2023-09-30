package handler

import (
	"h8-assignment-2/dto"
	"h8-assignment-2/pkg/errs"
	"h8-assignment-2/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	OrderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) orderHandler {
	return orderHandler{
		OrderService: orderService,
	}
}

func (oh *orderHandler) CreateOrder(ctx *gin.Context) {
	var newOrderRequest dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := oh.OrderService.CreateOrder(newOrderRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)

		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (oh *orderHandler) GetOrders(ctx *gin.Context) {
	response, err := oh.OrderService.GetOrders()

	if err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	// fmt.Println(response.Data)

	ctx.JSON(response.StatusCode, response)
}

func (oh *orderHandler) UpdateOrder(ctx *gin.Context) {
	var newOrderRequest dto.NewOrderRequest

	var orderId, _ = strconv.Atoi(ctx.Param("orderId"))

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := oh.OrderService.UpdateOrder(orderId, newOrderRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)

		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (oh *orderHandler) DeleteOrder(ctx *gin.Context) {
	var orderId, _ = strconv.Atoi(ctx.Param("orderId"))

	response, err := oh.OrderService.DeleteOrder(orderId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)

		return
	}

	ctx.JSON(response.StatusCode, response)
}
