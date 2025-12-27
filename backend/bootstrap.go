package backend

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/router"
	"github.com/gin-gonic/gin"
)

// StartBackendEngine initializes the new-api engine as a library
func StartBackendEngine(appDataPath string, port string) {
	// 1. Force Configuration Injection
	common.SQLitePath = filepath.Join(appDataPath, "new-api.db")
	// common.LogDir is a *string flag.
	logDir := filepath.Join(appDataPath, "logs")
	common.LogDir = &logDir

	// Disable Redis, Enable Memory Cache
	common.RedisEnabled = false
	common.MemoryCacheEnabled = true

	// Set as slave/standalone to avoid some distributed checks if applicable
	os.Setenv("NODE_TYPE", "slave")
	// Also set SQL_DSN for SQLite if new-api relies on env
	os.Setenv("SQL_DSN", common.SQLitePath)

	// Ensure directory exists
	if _, err := os.Stat(appDataPath); os.IsNotExist(err) {
		os.MkdirAll(appDataPath, 0755)
	}
	if _, err := os.Stat(*common.LogDir); os.IsNotExist(err) {
		os.MkdirAll(*common.LogDir, 0755)
	}

	// 2. Initialize new-api database
	// common.InitSQLite() does not exist in this version.
	// model.InitDB() handles connection if config is set.
	model.InitDB()

	// 3. Auto-inject Root Token
	ensureRootToken()

	// 4. Start Gin Router
	gin.SetMode(gin.ReleaseMode)

	// router.SetRouter(router *gin.Engine, buildFS embed.FS, indexPage []byte)
	// If we pass nil for router, it might crash if SetRouter expects an Engine.
	// But usually SetRouter creates one or configures the one passed.
	// Wait, standard OneAPI SetRouter signature is SetRouter(router *gin.Engine, buildFS embed.FS, indexPage []byte).
	// If it modifies the passed router, we should create one first.
	// If it returns *gin.Engine, then we might pass nil.
	// Let's check SetRouter signature again from grep: func SetRouter(router *gin.Engine, buildFS embed.FS, indexPage []byte)
	// It doesn't return anything. It returns void (or implicit in Go).
	// So we must create the engine first.

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// We need to pass valid embed.FS and []byte for indexPage even if we don't use them,
	// or passing nil might be handled.
	// But `embed.FS` is a struct, not a pointer, so we can't pass nil.
	// We need to pass an empty FS.
	var emptyFS embed.FS
	var emptyIndex []byte

	router.SetRouter(r, emptyFS, emptyIndex)

	go func() {
		err := r.Run(":" + port)
		if err != nil {
			panic(err)
		}
	}()
}

func ensureRootToken() {
	// Check if a root token exists, if not create one
	// This logic depends on model.Token
	var token model.Token
	// Try to find a specific root token or just any token
	// We'll create a specific one for the desktop app
	tokenName := "Novapi Root Token"
	err := model.DB.Where("name = ?", tokenName).First(&token).Error
	if err != nil {
		// Create it
		token = model.Token{
			Name: tokenName,
			Key: "sk-novapi-root",
			Status: common.TokenStatusEnabled,
			UnlimitedQuota: true,
		}
		model.DB.Create(&token)
	}

	// Export token to file for frontend
	// In a real Wails app, we might pass this via a bound method,
	// but writing to a file is a simple way to bootstrap.
	// We use common.LogDir's parent or appData path.
	// common.LogDir is set in StartBackendEngine.
	if common.LogDir != nil && *common.LogDir != "" {
		configDir := filepath.Dir(*common.LogDir)
		configPath := filepath.Join(configDir, "config.json")
		// Simple JSON with token
		content := fmt.Sprintf(`{"root_token": "%s", "api_url": "http://localhost:3000"}`, token.Key)
		os.WriteFile(configPath, []byte(content), 0644)
	}
}
