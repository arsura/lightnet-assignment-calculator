package main

import (
	"github.com/arsura/lightnet-assignment-calculator/routes"
)

func main() {
	router := routes.Router()
	router.Run()
}
