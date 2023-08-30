package main

import (
	"io"
	"log/slog"
	"os"
	"time"
)

func main() {
	var file *os.File
	var err error
	var writer io.Writer
	var handler slog.Handler
	var logger *slog.Logger

	file, err = os.Create("demo.log")
	if err != nil {
		panic(err)
	}
	writer = io.Writer(file)
	handler = slog.NewJSONHandler(writer, nil)
	logger = slog.New(handler)
	for i := 1; i <= 10; i++ {
		logger.Info("Hello world!")
		time.Sleep(1 * time.Second)
	}
}
