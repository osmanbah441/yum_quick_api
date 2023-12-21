package main

import (
	"fmt"
	"net/http"
)

func (app *application) serve() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	app.logger.Printf("[%s] server is listening on %s", app.config.env, server.Addr)

	if err := server.ListenAndServe(); err != nil {
		app.logger.Fatal(err)
	}
}
