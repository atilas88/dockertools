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

func (a *App) StartContainer(containerID string) error {
	return a.service.StartContainer(a.ctx, containerID)
}

func (a *App) StopContainer(containerID string) error {
	return a.service.StopContainer(a.ctx, containerID)
}

func (a *App) RemoveContainer(containerID string) error {
	return a.service.RemoveContainer(a.ctx, containerID)
}

func (a *App) RestartContainer(containerID string) error {
	return a.service.RestartContainer(a.ctx, containerID)
}

func (a *App) PauseContainer(containerID string) error {
	return a.service.PauseContainer(a.ctx, containerID)
}

func (a *App) UnpauseContainer(containerID string) error {
	return a.service.UnpauseContainer(a.ctx, containerID)
}
