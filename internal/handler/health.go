package handler

import (
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	res := api.healthResponse
}
