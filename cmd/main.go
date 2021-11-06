package main

import (
	"deuvox/internal/app"
	"fmt"
	"net/http"
)

func main() {
	app := app.New()
	fmt.Println("Run in localhost:3000")
	http.ListenAndServe(":3000", app.R)
}
