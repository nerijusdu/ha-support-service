package main

import (
	"fmt"
	"net/http"
	"os"

	"haservice/config"
	"haservice/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config := config.GetConfig()
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a catch-all route"))
	})
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	routes.RegisterRoutes(r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	serverUrl := os.Getenv("SERVER_URL")
	url := fmt.Sprintf("%s:%d", serverUrl, config.Port)
	fmt.Println("Server started on " + url)
	http.ListenAndServe(url, loggedRouter)
}
