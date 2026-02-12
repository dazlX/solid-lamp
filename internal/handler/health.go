package handler

import (
	"encoding/json"
	"net/http"
	Api "solid-lamp/api"
	"time"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	res := Api.HealthResponse{
		Status:    "success",
		Timestamp: time.Now(),
	}

	json.NewEncoder(w).Encode(res)
}
