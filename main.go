package main

import (
	"context"
	"fmt"

	"github.com/sushavanP/go-shopping-cart/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed sto start the app: ", err)
	}
}
