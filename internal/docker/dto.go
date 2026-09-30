package docker

type DockerInfo struct {
	Version    string `json:"version"`
	APIVersion string `json:"apiVersion"`
}
