package Server

import (
	"log"
	"net/http"
	handlers "solid-lamp/internal/handler"
)

func Run() {
	log.Println("Server Started")
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe(":8080", nil)

	if err == nil {
		log.Println("Error:", err)
	}
}
