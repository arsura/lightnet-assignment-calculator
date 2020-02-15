package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Operand struct {
	FirstVal float32 `json:"a"`
	SecndVal float32 `json:"b"`
}

func Sum(a float32, b float32) float32 {
	return a + b
}

func Sub(a float32, b float32) float32 {
	return a - b
}

func Mul(a float32, b float32) float32 {
	return a * b
}

func Div(a float32, b float32) float32 {
	return a / b
}

func Calculator(c *gin.Context) {
	var operand Operand
	c.BindJSON(&operand)

	// Get operator
	urlPath := strings.Split(c.Request.URL.Path, ".")
	operator := urlPath[len(urlPath)-1]

	// Calculate
	var result float32
	switch operator {
	case "sum":
		result = Sum(operand.FirstVal, operand.SecndVal)
	case "sub":
		result = Sub(operand.FirstVal, operand.SecndVal)
	case "mul":
		result = Mul(operand.FirstVal, operand.SecndVal)
	case "div":
		result = Div(operand.FirstVal, operand.SecndVal)
	default:
		return
	}

	resultAsString := fmt.Sprintf("%.4f", result)
	c.JSON(http.StatusOK, gin.H{
		"result":           result,
		"result_as_string": resultAsString,
	})
}
