package main

import (
	"context"
	"dockertools/internal/docker"
	"fmt"

	"github.com/moby/moby/client"
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
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetDockerInfo() (docker.DockerInfo, error) {
	apiClient, err := client.New(client.FromEnv)
	if err != nil {
		return docker.DockerInfo{}, fmt.Errorf("error al conectar con Docker: %w", err)
	}
	defer apiClient.Close()
	serverVersion, err := apiClient.ServerVersion(a.ctx, client.ServerVersionOptions{})
	if err != nil {
		return docker.DockerInfo{}, fmt.Errorf("error al obtener la versión del servidor: %w", err)
	}
	return docker.DockerInfo{
		Version:    serverVersion.Version,
		APIVersion: serverVersion.APIVersion,
	}, nil
}
