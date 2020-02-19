package router

import (
	"github.com/arsura/lightnet-assignment-calculator/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controller.Pong)

	router.POST("/calculator.sum", controller.SumHandler)
	router.POST("/calculator.sub", controller.SubHandler)
	router.POST("/calculator.mul", controller.MulHandler)
	router.POST("/calculator.div", controller.DivHandler)

	return router
}
