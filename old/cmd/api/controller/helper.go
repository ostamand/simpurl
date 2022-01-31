package controller

import (
	"net/http"
)

func AllowOrigins(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Headers", "Content-Type, Accept")
}
