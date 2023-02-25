package game

import (
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
)

type Game struct {
	App          *app.Application
	CurrentScene Scene
	Scenes       map[string]Scene
	aspectRatio  float32
}

func (g *Game) Init(callback func(*Scene)) {
	// Create application and scene
	g.App = app.App()
	g.CurrentScene.Init(callback)

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(g.CurrentScene)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	g.App.Subscribe(window.OnWindowSize, g.onResize)
	g.onResize("", nil)

}

// Set up callback to update viewport and camera aspect ratio when the window is resized
func (g *Game) onResize(evname string, ev interface{}) {
	// Get framebuffer size and update viewport accordingly
	width, height := g.App.GetSize()
	g.App.Gls().Viewport(0, 0, int32(width), int32(height))
	// Update the camera's aspect ratio
	g.CurrentScene.Camera.SetAspect(float32(width) / float32(height))
}

func (g *Game) Run() {
	// Set background color to gray
	g.App.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	// Run the application
	g.App.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		g.App.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(g.CurrentScene, g.CurrentScene.Camera)
	})
}
