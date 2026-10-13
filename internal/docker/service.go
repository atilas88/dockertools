package docker

import (
	"context"
	"fmt"
	"time"

	"strings"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

type Service struct {
	cli *client.Client
}

const defaultTimeout = 3 * time.Second

// NewService creates a new instance of the Service
func NewService(cli *client.Client) (*Service, error) {
	if cli == nil {
		var err error
		cli, err = client.New(client.FromEnv)
		if err != nil {
			return nil, fmt.Errorf("error al crear el cliente de Docker: %w", err)
		}
	}
	return &Service{
		cli: cli,
	}, nil
}

func (s *Service) Close() error {
	if s.cli != nil {
		return s.cli.Close()
	}
	return nil
}

func (s *Service) GetDockerInfo(ctx context.Context) (DockerInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	serverVersion, err := s.cli.ServerVersion(ctx, client.ServerVersionOptions{})
	if err != nil {
		return DockerInfo{}, fmt.Errorf("error al obtener la versión del servidor: %w", err)
	}
	return DockerInfo{
		Version:    serverVersion.Version,
		APIVersion: serverVersion.APIVersion,
		Host:       s.cli.DaemonHost(),
	}, nil
}

func (s *Service) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := s.cli.Ping(ctx, client.PingOptions{})
	return err
}

func (s *Service) ListContainers(ctx context.Context) ([]Container, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	containers, err := s.cli.ContainerList(ctx, client.ContainerListOptions{All: true})
	if err != nil {
		return nil, fmt.Errorf("error al listar los contenedores: %w", err)
	}

	result := make([]Container, 0, len(containers.Items))
	for _, c := range containers.Items {
		var exitCode *int
		if c.State == container.StateExited {
			var code int
			if _, err := fmt.Sscanf(c.Status, "Exited (%d)", &code); err == nil {
				exitCode = &code
			}
		}
		result = append(result, Container{
			ID:       c.ID[:12],
			Name:     strings.TrimPrefix(c.Names[0], "/"),
			Image:    c.Image,
			State:    ContainerState(c.State),
			ExitCode: exitCode,
			Ports:    formatPorts(c.Ports),
			Created:  c.Created,
		})
	}
	return result, nil
}

func formatPorts(ps []container.PortSummary) []string {
	out := make([]string, 0, len(ps))
	for _, p := range ps {
		if p.PublicPort == 0 {
			continue // puerto expuesto pero no publicado
		}
		out = append(out, fmt.Sprintf("%d→%d", p.PublicPort, p.PrivatePort))
	}
	return out
}

func (s *Service) StopContainer(ctx context.Context, containerID string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := s.cli.ContainerStop(ctx, containerID, client.ContainerStopOptions{})
	return err
}

func (s *Service) StartContainer(ctx context.Context, containerID string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := s.cli.ContainerStart(ctx, containerID, client.ContainerStartOptions{})
	return err
}

func (s *Service) RemoveContainer(ctx context.Context, containerID string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := s.cli.ContainerRemove(ctx, containerID, client.ContainerRemoveOptions{Force: true})
	return err
}

func (s *Service) RestartContainer(ctx context.Context, containerID string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := s.cli.ContainerRestart(ctx, containerID, client.ContainerRestartOptions{})
	return err
}

func (s *Service) PauseContainer(ctx context.Context, containerID string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := s.cli.ContainerPause(ctx, containerID, client.ContainerPauseOptions{})
	return err
}

func (s *Service) UnpauseContainer(ctx context.Context, containerID string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := s.cli.ContainerUnpause(ctx, containerID, client.ContainerUnpauseOptions{})
	return err
}
