package docker

type DockerInfo struct {
	Version    string `json:"version"`
	APIVersion string `json:"apiVersion"`
	Host       string `json:"host"`
}

// ContainerState mirrors Docker's container states (same string values).
type ContainerState string

const (
	StateCreated    ContainerState = "created"
	StateRunning    ContainerState = "running"
	StatePaused     ContainerState = "paused"
	StateRestarting ContainerState = "restarting"
	StateRemoving   ContainerState = "removing"
	StateExited     ContainerState = "exited"
	StateDead       ContainerState = "dead"
)

// AllContainerStates is registered with Wails EnumBind so the TypeScript
// enum docker.ContainerState is generated in wailsjs/go/models.ts.
var AllContainerStates = []struct {
	Value  ContainerState
	TSName string
}{
	{StateCreated, "CREATED"},
	{StateRunning, "RUNNING"},
	{StatePaused, "PAUSED"},
	{StateRestarting, "RESTARTING"},
	{StateRemoving, "REMOVING"},
	{StateExited, "EXITED"},
	{StateDead, "DEAD"},
}

type Container struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Image    string         `json:"image"`
	State    ContainerState `json:"state"`
	ExitCode *int     `json:"exitCode,omitempty"`
	Ports    []string `json:"ports"`
	Created  int64    `json:"created"`
}
