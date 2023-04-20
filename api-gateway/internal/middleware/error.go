package middleware

import (
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type ErrorHandlerFunc func(http.ResponseWriter, *http.Request) error

type Error struct {
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func ErrorMiddleware(handler ErrorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {

			requestID, ok := r.Context().Value(middleware.RequestIDHeader).(string)
			if !ok {
				requestID = "unknown"
			}

			response := Error{
				Message:   err.Error(),
				RequestID: requestID,
			}

			jsonResponse, err := json.Marshal(response)
			if err != nil {
				msg := "failed to marshal error json"
				log.Println(msg)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(msg))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonResponse)
		}
	}
}
