package backend

import (
	"fmt"
	"os/exec"
	"net"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/model"
)

// LocalProcessManager handles local model processes
type LocalProcessManager struct{}

// StartLocalModel starts a local model process and registers it to new-api
func (m *LocalProcessManager) StartLocalModel(modelPath string, modelName string) (string, error) {
	// 1. Get available port
	port, err := getAvailablePort()
	if err != nil {
		return "", err
	}

	// 2. Start process (Stub: assumes llama-server is in path or relative)
	// In a real app, we would manage the child process lifecycle properly
	cmd := exec.Command("./bin/llama-server", "-m", modelPath, "--port", port)
	err = cmd.Start()
	if err != nil {
		return "", fmt.Errorf("failed to start model: %w", err)
	}

	// Wait a bit for server to start (naive)
	time.Sleep(2 * time.Second)

	// 3. Register to new-api
	baseURL := fmt.Sprintf("http://localhost:%s", port)
	channel := model.Channel{
		Type: constant.ChannelTypeOllama, // Use constant package
		Name: "Local-" + modelName,
		BaseURL: &baseURL, // Use pointer
		Key: "sk-ignore",
		Models: modelName,
		Status: common.ChannelStatusEnabled,
	}

	// Check if channel already exists (by BaseURL maybe?)
	// For simplicity, just create new one or update if we had a stable ID
	// Here we just insert
	if err := model.DB.Create(&channel).Error; err != nil {
		return "", fmt.Errorf("failed to register channel: %w", err)
	}

	// 4. Refresh Cache
	// Since we are adding a new channel, we should update the memory cache.
	// We can use model.InitChannelCache() which rebuilds the cache from DB,
	// or if there is a more granular method.
	// model.CacheUpdateChannel accepts a *Channel, but based on the code I read,
	// it only updates existing map entries or adds to ID map, but might not update group2model2channels index completely
	// unless we call InitChannelCache or if CacheUpdateChannel handles it (it seems it doesn't handle group map update in the file I read).
	// However, InitChannelCache() is safe to call as it rebuilds everything.
	if common.MemoryCacheEnabled {
		model.InitChannelCache()
	}

	return fmt.Sprintf("Model %s started on port %s", modelName, port), nil
}

func getAvailablePort() (string, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}
	defer listener.Close()
	port := listener.Addr().(*net.TCPAddr).Port
	return fmt.Sprintf("%d", port), nil
}
