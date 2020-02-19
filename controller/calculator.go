package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Operand struct {
	FirstVal float64 `json:"a"`
	SecndVal float64 `json:"b"`
}

func Sum(a float64, b float64) float64 {
	return a + b
}

func Sub(a float64, b float64) float64 {
	return a - b
}

func Mul(a float64, b float64) float64 {
	return a * b
}

func Div(a float64, b float64) float64 {
	return a / b
}

func SumHandler(c *gin.Context) {
	var operand Operand
	c.BindJSON(&operand)

	result := Sum(operand.FirstVal, operand.SecndVal)

	resultAsString := fmt.Sprintf("%.4f", result)
	c.JSON(http.StatusOK, gin.H{
		"result":           result,
		"result_as_string": resultAsString,
	})
}

func SubHandler(c *gin.Context) {
	var operand Operand
	c.BindJSON(&operand)

	result := Sub(operand.FirstVal, operand.SecndVal)

	resultAsString := fmt.Sprintf("%.4f", result)
	c.JSON(http.StatusOK, gin.H{
		"result":           result,
		"result_as_string": resultAsString,
	})
}

func MulHandler(c *gin.Context) {
	var operand Operand
	c.BindJSON(&operand)

	result := Mul(operand.FirstVal, operand.SecndVal)

	resultAsString := fmt.Sprintf("%.4f", result)
	c.JSON(http.StatusOK, gin.H{
		"result":           result,
		"result_as_string": resultAsString,
	})
}

func DivHandler(c *gin.Context) {
	var operand Operand
	c.BindJSON(&operand)

	if operand.SecndVal == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "undefined",
		})
		return
	}

	result := Div(operand.FirstVal, operand.SecndVal)

	resultAsString := fmt.Sprintf("%.4f", result)
	c.JSON(http.StatusOK, gin.H{
		"result":           result,
		"result_as_string": resultAsString,
	})
}
