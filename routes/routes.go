package routes

import (
	"github.com/arsura/lightnet-assignment-calculator/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.Pong)

	router.POST("/calculator.sum", controllers.SumHandler)
	router.POST("/calculator.sub", controllers.SubHandler)
	router.POST("/calculator.mul", controllers.MulHandler)
	router.POST("/calculator.div", controllers.DivHandler)

	return router
}
