package main

import (
	"context"
	"fmt"

	"novapi/backend"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Initialize backend engine
	// Get AppData directory
	// Note: In a real app, you should use runtime.Environment(ctx) to get user data dir correctly
	// Here we use a local logic or standard Go way
	// runtime.Environment does not give AppData path directly, usually we use os.UserConfigDir or similar
	// But Wails has a way.

	// For now, let's use current directory or a hardcoded one for testing, or better, use relative to exe
	appDataPath := "./data"

	// Start the backend engine (new-api) on port 3000
	backend.StartBackendEngine(appDataPath, "3000")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
