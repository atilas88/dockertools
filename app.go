package main

import (
	"context"
	"dockertools/internal/docker"
)

// App struct
type App struct {
	ctx     context.Context
	service *docker.Service
}

// NewApp creates a new App application struct
func NewApp(svc *docker.Service) *App {
	return &App{
		service: svc,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	if a.service != nil {
		a.service.Close()
	}
}

func (a *App) GetDockerInfo() (docker.DockerInfo, error) {
	return a.service.GetDockerInfo(a.ctx)
}

func (a *App) GetDockerStatus() error {
	return a.service.Ping(a.ctx)
}

func (a *App) ListContainers() ([]docker.Container, error) {
	return a.service.ListContainers(a.ctx)
}
