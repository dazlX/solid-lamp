package Server

import (
	"log"
	"net/http"
)

func Run() {
	log.Println("Server Started")
	http.HandleFunc("/", RootHandler)
	err := http.ListenAndServe(":8080", nil)

	if err == nil {
		log.Println("Error:", err)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Root route used!!")
}
