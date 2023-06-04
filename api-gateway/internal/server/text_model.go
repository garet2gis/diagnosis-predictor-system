package server

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type TextModelRequest struct {
	Text string `json:"text"`
}

func (s *Server) TextModel(w http.ResponseWriter, r *http.Request) error {
	var input TextModelRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return err
	}

	reqBody, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		s.cfg.ServicesURL.TextModel,
		bytes.NewBuffer(reqBody),
	)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var output map[string]string

	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return err
	}

	outputJson, err := json.Marshal(output)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(outputJson)

	return nil
}
