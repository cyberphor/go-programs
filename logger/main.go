package main

import (
	"log/slog"
	"os"
)

func main() {
	var logger *slog.Logger
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Hello world!")
	/*
	  Output:
	  {"time":"2023-08-30T12:59:51.6032416-04:00","level":"INFO","msg":"Hello world!"}
	*/
}
