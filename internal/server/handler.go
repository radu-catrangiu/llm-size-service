package server

import (
	"encoding/json"
	"llm-size-service/internal/evaluator"
	"log/slog"
	"net/http"
)

func (s *Server) evaluateHandler(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Request received", "url", r.URL.String())

	query := r.URL.Query()
	model := query.Get("model")

	if model == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e := evaluator.New(model, "")
	resp, err := e.GetSize()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	resBytes, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(resBytes)
}
