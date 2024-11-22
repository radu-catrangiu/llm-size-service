package server

import (
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	hfToken string
}

func New(hfToken string) *Server {
	server := &Server{
		hfToken: hfToken,
	}

	return server
}

func getAddr() string {
	host, found := os.LookupEnv("SERVER_HOST")
	if !found {
		host = "127.0.0.1"
	}

	port, found := os.LookupEnv("SERVER_PORT")
	if !found {
		port = "3000"
	}

	return fmt.Sprintf("%s:%s", host, port)
}

func (s *Server) Listen() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /evaluate", s.evaluateHandler)

	addr := getAddr()
	println("Server started on", addr)
	// slog.Info("Server started", "addr", addr)
	return http.ListenAndServe(addr, mux)
}
