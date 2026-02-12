package Server

import (
	"log"
	"net/http"
)

func Run() {
	log.Println("Server Started")
	err := http.ListenAndServe(":8080", nil)

	if err == nil {
		log.Println("Error:", err)
	}
}
