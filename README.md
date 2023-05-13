# go-windows-service
A library that wraps Windows SCM operations into one nice call.

### Windows
The wrapper assumes a typical CLI for Windows services.
```
Usage: example.exe </i>|</u>|</d>
        /i - install service.
        /u - uninstall service.
        /d - debug with console.
```

Install and uninstall operations should runs as admin.

### Linux and other
The wrapper tries not to interfere with the work and immediately transfers control to the nested function.
All that the wrapper does is:
- sets a panic handler before calling the nested function
- sets signal handlers for graceful shutdown the application

See [entry_other.go](https://github.com/sergereinov/go-windows-service/blob/main/entry_other.go) file.

### Examples
```go
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
```

Examples in projects that use this library:
- https://github.com/sergereinov/nf-svc

## Similar Projects
- https://github.com/judwhite/go-svc
- more complex https://github.com/kardianos/service
- and many other at https://pkg.go.dev/
