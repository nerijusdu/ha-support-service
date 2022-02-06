package main

import (
	"fmt"
	"net/http"

	"hatvservice/config"
	"hatvservice/routes"

	"github.com/gorilla/mux"
)

func main() {
	config := config.GetConfig()
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	url := fmt.Sprintf("127.0.0.1:%d", config.Port)
	fmt.Println("Server started on " + url)
	http.ListenAndServe(url, r)
}
