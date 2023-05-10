package main

import (
	"context"
	"log"

	service "github.com/sergereinov/go-windows-service"
)

var (
	// You can set the Version at compile stage of your dev pipeline with:
	// go build -ldflags="-X main.Version=1.0.0" ./example
	Version     = "1.0.0"
	Name        = service.ExecutableFilename()
	Description = "My service"
)

func main() {
	// Init your favorite logger
	logger := log.Default()

	// Run service wrapper
	service.Service{
		Version:     Version,
		Name:        Name,
		Description: Description,
		Logger:      logger,
	}.Proceed(func(ctx context.Context) {

		logger.Printf("Service %s v%s started", Name, Version)

		//Do what the service should do

		<-ctx.Done()

		logger.Printf("Service %s v%s stopped", Name, Version)
	})
}
