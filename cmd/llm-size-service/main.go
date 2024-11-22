package main

import (
	"llm-size-service/internal/server"
	"log/slog"
	"os"
	"strings"
)

func initLogging() {
	logLevelValue := os.Getenv("LOG_LEVEL")
	level := slog.LevelInfo

	switch strings.ToUpper(logLevelValue) {
	case "DEBUG":
		level = slog.LevelDebug
	case "ERROR":
		level = slog.LevelError
	}

	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})
	slog.SetDefault(slog.New(h))
}

func main() {
	initLogging()

	hfToken := os.Getenv("HF_API_KEY")
	srv := server.New(hfToken)

	err := srv.Listen()
	if err != nil {
		slog.Error("Server exited with error", "error", err)
	}
}
