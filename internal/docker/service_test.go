package docker

import (
	"context"
	"errors"
	"slices"
	"testing"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

func intPtr(i int) *int { return &i }

type fakeDockerAPI struct {
	containers []container.Summary
	err        error
	stoppedId  string
}

func (f *fakeDockerAPI) ContainerList(ctx context.Context, options client.ContainerListOptions) (client.ContainerListResult, error) {
	if f.err != nil {
		return client.ContainerListResult{}, f.err
	}
	return client.ContainerListResult{Items: f.containers}, nil
}

func (f *fakeDockerAPI) ContainerStart(ctx context.Context, containerID string, options client.ContainerStartOptions) (client.ContainerStartResult, error) {
	return client.ContainerStartResult{}, nil
}

func (f *fakeDockerAPI) ContainerStop(ctx context.Context, containerID string, options client.ContainerStopOptions) (client.ContainerStopResult, error) {
	f.stoppedId = containerID
	return client.ContainerStopResult{}, nil
}

func (f *fakeDockerAPI) ContainerRestart(ctx context.Context, containerID string, options client.ContainerRestartOptions) (client.ContainerRestartResult, error) {
	return client.ContainerRestartResult{}, nil
}

func (f *fakeDockerAPI) ContainerPause(ctx context.Context, containerID string, options client.ContainerPauseOptions) (client.ContainerPauseResult, error) {
	return client.ContainerPauseResult{}, nil
}

func (f *fakeDockerAPI) ContainerUnpause(ctx context.Context, containerID string, options client.ContainerUnpauseOptions) (client.ContainerUnpauseResult, error) {
	return client.ContainerUnpauseResult{}, nil
}

func (f *fakeDockerAPI) ContainerRemove(ctx context.Context, containerID string, options client.ContainerRemoveOptions) (client.ContainerRemoveResult, error) {
	return client.ContainerRemoveResult{}, nil
}

func (f *fakeDockerAPI) Ping(ctx context.Context, options client.PingOptions) (client.PingResult, error) {
	return client.PingResult{}, nil
}

func (f *fakeDockerAPI) ServerVersion(ctx context.Context, options client.ServerVersionOptions) (client.ServerVersionResult, error) {
	return client.ServerVersionResult{}, nil
}

func (f *fakeDockerAPI) DaemonHost() string {
	return "unix:///var/run/docker.sock"
}

func (f *fakeDockerAPI) Close() error {
	return nil
}

func TestFormatPorts(t *testing.T) {
	testCases := []struct {
		name     string
		ports    []container.PortSummary
		expected []string
	}{
		{
			name:     "published port",
			ports:    []container.PortSummary{{PublicPort: 8080, PrivatePort: 80}},
			expected: []string{"8080→80"},
		},
		{
			name:     "exposed but not published i'll be ignored",
			ports:    []container.PortSummary{{PrivatePort: 80}},
			expected: []string{},
		},
		{
			name:     "multiple ports",
			ports:    []container.PortSummary{{PublicPort: 8080, PrivatePort: 80}, {PublicPort: 9090, PrivatePort: 90}},
			expected: []string{"8080→80", "9090→90"},
		},
		{
			name:     "no ports",
			ports:    []container.PortSummary{},
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := formatPorts(tc.ports)
			if !slices.Equal(result, tc.expected) {
				t.Errorf("formatPorts() = %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestToContainer(t *testing.T) {
	testCases := []struct {
		name     string
		input    container.Summary
		expected Container
	}{
		{
			name: "running container",
			input: container.Summary{
				ID:    "123",
				Names: []string{"/my-container"},
				Image: "my-image",
				State: "running",
				Ports: []container.PortSummary{{PublicPort: 8080, PrivatePort: 80}},
			},
			expected: Container{
				ID:       "123",
				Name:     "my-container",
				Image:    "my-image",
				State:    StateRunning,
				ExitCode: nil,
				Ports:    []string{"8080→80"},
			},
		},
		{
			name: "exited container with exit code",
			input: container.Summary{
				ID:     "456",
				Names:  []string{"/another-container"},
				Image:  "another-image",
				State:  "exited",
				Status: "Exited (1) 2 hours ago",
			},
			expected: Container{
				ID:       "456",
				Name:     "another-container",
				Image:    "another-image",
				State:    StateExited,
				ExitCode: intPtr(1),
				Ports:    []string{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := toContainer(tc.input)
			if result.ID != tc.expected.ID ||
				result.Name != tc.expected.Name ||
				result.Image != tc.expected.Image ||
				result.State != tc.expected.State ||
				(result.ExitCode == nil) != (tc.expected.ExitCode == nil) ||
				(result.ExitCode != nil && *result.ExitCode != *tc.expected.ExitCode) ||
				!slices.Equal(result.Ports, tc.expected.Ports) {
				t.Errorf("toContainer() = %+v, want %+v", result, tc.expected)
			}
		})
	}
}

func TestStopContainer(t *testing.T) {
	fakeAPI := &fakeDockerAPI{}
	service := &Service{cli: fakeAPI}

	err := service.StopContainer(context.Background(), "123")
	if err != nil {
		t.Fatalf("StopContainer() error = %v", err)
	}

	if fakeAPI.stoppedId != "123" {
		t.Errorf("StopContainer() stoppedId = %v, want %v", fakeAPI.stoppedId, "123")
	}
}

func TestListContainers(t *testing.T) {
	fakeAPI := &fakeDockerAPI{
		containers: []container.Summary{
			{ID: "123", Names: []string{"/my-container"}, Image: "my-image", State: "running"},
			{ID: "456", Names: []string{"/another-container"}, Image: "another-image", State: "exited", Status: "Exited (1) 2 hours ago"},
		},
	}
	service := &Service{cli: fakeAPI}

	result, err := service.ListContainers(context.Background())
	if err != nil {
		t.Fatalf("ListContainers() error = %v", err)
	}

	expected := []Container{
		{ID: "123", Name: "my-container", Image: "my-image", State: StateRunning, ExitCode: nil, Ports: []string{}},
		{ID: "456", Name: "another-container", Image: "another-image", State: StateExited, ExitCode: intPtr(1), Ports: []string{}},
	}

	if len(result) != len(expected) {
		t.Fatalf("ListContainers() = %v, want %v", result, expected)
	}

	for i := range result {
		if result[i].ID != expected[i].ID ||
			result[i].Name != expected[i].Name ||
			result[i].Image != expected[i].Image ||
			result[i].State != expected[i].State ||
			(result[i].ExitCode == nil) != (expected[i].ExitCode == nil) ||
			(result[i].ExitCode != nil && *result[i].ExitCode != *expected[i].ExitCode) ||
			!slices.Equal(result[i].Ports, expected[i].Ports) {
			t.Errorf("ListContainers()[%d] = %+v, want %+v", i, result[i], expected[i])
		}
	}
}

func TestListContainers_DockerError(t *testing.T) {
	fakeErr := errors.New("daemon error")
	s := &Service{cli: &fakeDockerAPI{err: fakeErr}}

	_, err := s.ListContainers(context.Background())

	if !errors.Is(err, fakeErr) {
		t.Errorf("err = %v, want it to wrap %v", err, fakeErr)
	}
}

func TestPing(t *testing.T) {
	fakeAPI := &fakeDockerAPI{}
	service := &Service{cli: fakeAPI}

	err := service.Ping(context.Background())
	if err != nil {
		t.Fatalf("Ping() error = %v", err)
	}
}
