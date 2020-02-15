package routes

import (
	"github.com/arsura/lightnet-assignment-calculator/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.Pong)

	router.POST("/calculator.sum", controllers.CalculatorHandler)
	router.POST("/calculator.sub", controllers.CalculatorHandler)
	router.POST("/calculator.mul", controllers.CalculatorHandler)
	router.POST("/calculator.div", controllers.CalculatorHandler)

	return router
}
