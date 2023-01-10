package main

import (
	"fmt"

	_router "github.com/online.scheduling-api/src/core"
)

func main() {
	serve()
}

func serve() {
	router := _router.ConfigureRouter()

	if err := router.Run("localhost:8080"); err != nil {
		fmt.Println("Error serving API")
	}
}
