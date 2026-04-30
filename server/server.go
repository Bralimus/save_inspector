package server

import (
	"fmt"
	"net/http"

	"github.com/Bralimus/save_inspector/app"
)

func Start(app *app.App) {
	fmt.Println("Starting server on http://localhost:8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
