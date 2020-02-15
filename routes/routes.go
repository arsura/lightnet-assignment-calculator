package routes

import (
	"github.com/arsura/lightnet-assignment-calculator/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.Pong)

	router.POST("/calculator.sum", controllers.Calculator)
	router.POST("/calculator.sub", controllers.Calculator)
	router.POST("/calculator.mul", controllers.Calculator)
	router.POST("/calculator.div", controllers.Calculator)

	return router
}
