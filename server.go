package main

import (
	"github.com/arsura/lightnet-assignment-calculator/router"
)

func main() {
	router := router.Router()
	router.Run(":8081")
}
