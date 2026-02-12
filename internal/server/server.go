package Server

import (
	"log"
	"net/http"
	"solid-lamp/internal/handler"
)

func Run() {
	log.Println("Server Started")
	http.HandleFunc("/api/health", handler.HealthHandler)
	http.HandleFunc("/", handler.RootHandler)
	err := http.ListenAndServe(":8080", nil)

	if err == nil {
		log.Println("Error:", err)
	}
}
