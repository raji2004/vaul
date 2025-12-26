package main

import (
	"log"
	"sync"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var (
	mainWindowRef *application.WebviewWindow
	mainWindowMu  sync.RWMutex
)

// WindowService handles window operations
type WindowService struct{}

// NewWindowService creates a new WindowService instance
func NewWindowService() *WindowService {
	return &WindowService{}
}

// SetMainWindowRef sets the main window reference (called from main.go)
func SetMainWindowRef(window *application.WebviewWindow) {
	mainWindowMu.Lock()
	defer mainWindowMu.Unlock()
	mainWindowRef = window
}

// OpenMainWindow shows and focuses the main window
func (ws *WindowService) OpenMainWindow() error {
	mainWindowMu.RLock()
	window := mainWindowRef
	mainWindowMu.RUnlock()

	if window == nil {
		log.Println("WindowService: main window reference is nil")
		return nil
	}

	// Always show the window, even if it was previously closed/hidden
	window.Show()
	window.Focus()

	// Bring window to front (unminimize if minimized)
	window.UnMinimise()
	return nil
}
