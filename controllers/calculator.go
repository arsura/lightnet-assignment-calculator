package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Calculator struct {
	FirstVal	float32    `json:"a"`
	SecndVal	float32    `json:"b"`
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

func CalculatorHandler(c *gin.Context) {
	var calculator Calculator
	c.BindJSON(&calculator)

	// Get operator
	urlPath := strings.Split(c.Request.URL.Path, ".")
	op := urlPath[len(urlPath) - 1]

	// Calculate
	var result float32
    switch op {
    case "sum":
    	result = Sum(calculator.FirstVal, calculator.SecndVal)
    case "sub":
        result = Sub(calculator.FirstVal, calculator.SecndVal)
    case "mul":
		result = Mul(calculator.FirstVal, calculator.SecndVal)
	case "div":
		result = Div(calculator.FirstVal, calculator.SecndVal)
	default:
		return
    }

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}