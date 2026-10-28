package docker

import (
	"slices"
	"testing"

	"github.com/moby/moby/api/types/container"
)

func intPtr(i int) *int { return &i }

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
