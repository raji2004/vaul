package main

import "github.com/wailsapp/wails/v3/pkg/application"

// AppService handles application-level operations
type AppService struct {
	app *application.App
}

// NewAppService creates a new AppService instance
func NewAppService() *AppService {
	return &AppService{}
}

// SetApp sets the application reference
func (as *AppService) SetApp(app *application.App) {
	as.app = app
}

// Quit quits the application
func (as *AppService) Quit() error {
	if as.app != nil {
		as.app.Quit()
	}
	return nil
}

