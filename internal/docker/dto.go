package docker

type DockerInfo struct {
	Version    string `json:"version"`
	APIVersion string `json:"apiVersion"`
	Host       string `json:"host"`
}
