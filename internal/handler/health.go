package handler

import (
	"encoding/json"
	"log"
	"net/http"
	Api "solid-lamp/api"
	"time"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	res := Api.HealthResponse{
		Status:    "success",
		Timestamp: time.Now(),
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Println(err)
	}
}
