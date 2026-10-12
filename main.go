package main

import (
	"embed"
	"log"

	"dockertools/internal/docker"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	svc, err := docker.NewService(nil)
	if err != nil {
		log.Fatalf("Error al crear el servicio de Docker: %v", err)
	}
	app := NewApp(svc)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Dockertools",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 14, G: 17, B: 22, A: 255}, // --bg-000 (dark)
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
		EnumBind: []interface{}{
			docker.AllContainerStates,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
