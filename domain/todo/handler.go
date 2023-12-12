package todo

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

var orders = []Order{}

type Item struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ItemCode    int       `json:"itemcode"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderId     int       `json:"orderId"`
}

type Order struct {
	Id           int       `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CustomerName string    `json:"customer_name"`
	Items        []Item    `json:"items"`
}

type OrderRequest struct {
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Items        Item      `json:"items"`
}

type OrderUpdate struct {
	CustomerName string `json:"customerName"`
	Items        []Item `json:"items"`
}

type Response struct {
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func (h Handler) CreateOrder(c *gin.Context) {
	var req OrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "fail",
			Error:   err.Error(),
		})
		return
	}

	if req.CustomerName == "" {
		c.JSON(http.StatusBadRequest, Response{
			Message: "fail",
			Error:   "name is required",
		})
		return

	}

	orders = append(orders, Order{
		Id:           len(orders) + 1,
		CreatedAt:    req.OrderedAt,
		CustomerName: req.CustomerName,
		Items:        []Item{req.Items},
	})

	c.JSON(http.StatusCreated, Response{
		Message: "success",
	})
}

func (h Handler) GetAll(c *gin.Context) {

	c.JSON(http.StatusCreated, Response{
		Message: "success",
		Payload: orders,
	})
}

func (h Handler) DeleteById(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "fail",
			Error:   err.Error(),
		})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusBadRequest, Response{
			Message: "fail",
			Error:   "empty order list",
		})
		return
	}

	index := -1
	for i, order := range orders {
		if order.Id == idInt {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, Response{
			Message: "fail",
			Error:   "no id in this resources",
		})
		return
	}

	orders = append(orders[:index], orders[index+1:]...)

	c.JSON(http.StatusOK, Response{
		Message: "success",
		Payload: orders,
	})
}

func (h Handler) UpdatedOrder(c *gin.Context) {

	var req OrderUpdate

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "fail",
			Error:   err.Error(),
		})
		return
	}

	index := -1
	var order *Order
	for i, o := range orders {
		if o.Id == idInt {
			index = i
			order = &orders[i]
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, Response{
			Message: "fail",
			Error:   "no id in this resources",
		})
		return
	}

	existingOrder := *order

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "fail",
			Error:   err.Error(),
		})
		return
	}

	existingOrder.CustomerName = req.CustomerName
	existingOrder.Items = req.Items
	existingOrder.UpdatedAt = time.Now()

	orders[index] = existingOrder

	c.JSON(http.StatusOK, Response{
		Message: "success",
		Payload: orders,
	})
}
