package main

import (
	"context"
	"os"
	"path/filepath"

	service "github.com/sergereinov/go-windows-service"
	"github.com/sergereinov/go-windows-service/example/logger"
)

var (
	// You can set the Version at compile stage of your dev pipeline with:
	// go build -ldflags="-X main.Version=1.0.0" ./example
	Version     = "1.0.0"
	ServiceName = service.ExecutableFilename()
	Description = "My service"
)

func main() {
	// Example of initializing a file-logger with output to the console in debug mode
	exec, _ := os.Executable()
	logfile := filepath.Join(filepath.Dir(exec), ServiceName+".log")
	logToConsole := service.IsDebugMode()
	logger := logger.New(logfile, logToConsole)
	defer logger.Close()

	// Run service wrapper
	service.Service{
		Version:     Version,
		Name:        ServiceName,
		Description: Description,
		Logger:      logger,
	}.Proceed(func(ctx context.Context) {

		logger.Printf("Service %s v%s started", ServiceName, Version)

		//Do what the service should do

		<-ctx.Done()

		logger.Printf("Service %s v%s stopped", ServiceName, Version)
	})
}
