package main

import (
	"fmt"

	"github.com/if-i-am-gone/if-i-am-gone-web-backend/internal/app"
)

func main() {

	err := app.Run()

	if err != nil {
		mainErr := fmt.Errorf("main err on app: %w", err)
		panic(mainErr)
	}
}
