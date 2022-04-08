package main

import (
	"context"
	"raverte/raverte"
)

// App struct
type App struct {
	ctx     context.Context
	raverte *raverte.Raverte
}

// NewApp creates a new App application struct
func NewApp(r *raverte.Raverte) *App {
	app := App{}
	app.raverte = r
	return &app
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.raverte.LoadProfile()
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (a *App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}
