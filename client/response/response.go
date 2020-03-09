package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	type response struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	err := json.NewEncoder(w).Encode(response{Status: true, Message: message, Data: data})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, message string, err error) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	if err != nil {
		type response struct {
			Status  bool
			Message string `json:"message"`
			Error   string `json:"error"`
		}

		err := json.NewEncoder(w).Encode(response{Status: false, Message: message, Error: err.Error()})
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
	}
}
