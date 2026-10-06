package docker

import (
	"context"
	"fmt"
	"time"

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
	}, nil
}

func (s *Service) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := s.cli.Ping(ctx, client.PingOptions{})
	return err
}
