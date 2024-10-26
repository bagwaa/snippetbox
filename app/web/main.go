package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Create an application structure type, think of this like an IOC container in Laravel, for now it just
// contains a pointer to a logger, but it could contain other dependencies.
type application struct {
	logger *slog.Logger
}

func main() {
	// Grab the addr value from the command line and store it in the addr variable as
	// a string.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Read the value from the command line and store it in the addr variable
	flag.Parse()

	// Create the slog logger, using a text handler to write to os.Stdout
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	// Create a global application and store the logger in it
	app := &application{
		logger: logger,
	}

	// Start the server and pull in the routes from the routes method
	err := http.ListenAndServe(*addr, app.routes())

	// Log any errors and exit the program
	logger.Error(err.Error())
	os.Exit(1)
}
