package docker

import (
	"slices"
	"testing"

	"github.com/moby/moby/api/types/container"
)

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
