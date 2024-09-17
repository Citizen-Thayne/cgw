package main

import (
	"cgw/pkg/app"
)

func main() {
	router := app.CreateRouter()

	router.Run(":8080")
}
