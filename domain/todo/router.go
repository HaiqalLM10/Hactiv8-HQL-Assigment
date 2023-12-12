package todo

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.RouterGroup) {
	handler := NewHandler()
	orders := router.Group("/orders")
	{
		orders.POST("", handler.CreateOrder)
		orders.GET("", handler.GetAll)
		orders.PUT("/:id", handler.UpdatedOrder)
		orders.DELETE("/:id", handler.DeleteById)
	}
}
