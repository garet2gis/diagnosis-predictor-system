package server

import (
	"encoding/json"
	"net/http"
)

type TestModelInput struct {
	BloodAnalysis string `json:"blood_analysis"`
	Symptom       string `json:"symptom"`
}

func (s *Server) TestModel(w http.ResponseWriter, r *http.Request) error {
	var input TestModelInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return err
	}

	// TODO: pass input to message broker

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(input.Symptom))

	return nil
}
