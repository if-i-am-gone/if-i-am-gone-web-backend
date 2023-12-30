package main

import (
	"fmt"

	"github.com/ididie/ifidie_backend/internal/app"
)

func main() {

	err := app.Run()

	if err != nil {
		mainErr := fmt.Errorf("main err on app: %w", err)
		panic(mainErr)
	}
}
