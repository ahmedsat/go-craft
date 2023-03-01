package app

import (
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/audio/al"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util"
	"github.com/g3n/engine/util/logger"
	"github.com/g3n/engine/util/stats"
	"github.com/g3n/engine/window"
)

type App struct {
	*app.Application                  // Embedded standard application object
	log              *logger.Logger   // Application logger
	scene            *core.Node       // Scene rendered
	ambLight         *light.Ambient   // Scene ambient light
	frameRater       *util.FrameRater // Render loop frame rater
	world            World            // current world

	// GUI
	stats      *stats.Stats      // statistics object
	statsTable *stats.StatsTable // statistics table panel

	camera *camera.Camera

	fps float64
}

func (a *App) FrameRater() *util.FrameRater { return a.frameRater }

func (a *App) Scene() *core.Node { return a.scene }

func (a *App) Camera() *camera.Camera { return a.camera }

func (a *App) Log() *logger.Logger { return a.log }

// OnWindowResize is default handler for window resize events.
func (a *App) OnWindowResize() {

	// Get framebuffer size and set the viewport accordingly
	width, height := a.GetFramebufferSize()
	a.Gls().Viewport(0, 0, int32(width), int32(height))

	// Set camera aspect ratio
	a.camera.SetAspect(float32(width) / float32(height))

}

// logStats generate log with current statistics
func (a *App) logStats() {

	const statsFormat = `
         Shaders: %d
            Vaos: %d
         Buffers: %d
        Textures: %d
  Uniforms/frame: %d
Draw calls/frame: %d
 CGO calls/frame: %d
`
	a.log.Info(statsFormat,
		a.stats.Glstats.Shaders,
		a.stats.Glstats.Vaos,
		a.stats.Glstats.Buffers,
		a.stats.Glstats.Textures,
		a.stats.Unisets,
		a.stats.Drawcalls,
		a.stats.Cgocalls,
	)
}

func (a *App) setupScene() {
	// If there was a previous demo running, execute its Cleanup() method
	if a.world != nil {
		a.world.Cleanup(a)
	}

	// Destroy all objects in demo scene and GUI
	a.scene.DisposeChildren(true)

	// Clear subscriptions with ID (every subscribe called by demos should use the app address as ID so we can unsubscribe here)
	a.UnsubscribeAllID(a)

	// Clear all custom cursors and reset current cursor
	a.DisposeAllCustomCursors()
	a.SetCursor(window.ArrowCursor)

	// Set default background color
	a.Gls().ClearColor(0.6, 0.6, 0.6, 1.0)

	// Reset renderer z-sorting flag
	a.Renderer().SetObjectSorting(true)

	// Reset ambient light
	a.ambLight.SetColor(&math32.Color{R: 1.0, G: 1.0, B: 1.0})
	a.ambLight.SetIntensity(0.5)

	// Reset camera
	a.camera.SetPosition(0, 0, 5)
	a.camera.UpdateSize(5)
	a.camera.LookAt(&math32.Vector3{X: 0, Y: 0, Z: 0}, &math32.Vector3{X: 0, Y: 1, Z: 0})
	a.camera.SetProjection(camera.Perspective)

	// If audio active, resets global listener parameters
	al.Listener3f(al.Position, 0, 0, 0)
	al.Listener3f(al.Velocity, 0, 0, 0)
	al.Listenerfv(al.Orientation, []float32{0, 0, -1, 0, 1, 0})

}

func (a *App) Run() {
	a.Application.Run(a.Update)
}

func (a *App) Update(rend *renderer.Renderer, deltaTime time.Duration) {
	// Start measuring this frame
	a.frameRater.Start()

	// Clear the color, depth, and stencil buffers
	a.Gls().Clear(gls.COLOR_BUFFER_BIT | gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT) // TODO maybe do inside renderer, and allow customization

	// Update the current running demo if any
	if a.world != nil {
		a.world.Update(a, deltaTime)
	}

	// Render scene
	err := rend.Render(a.scene, a.camera)
	if err != nil {
		panic(err)
	}

	// Update GUI timers
	gui.Manager().TimerManager.ProcessTimers()

	// Update statistics
	if a.stats.Update(time.Second) {
		if a.statsTable != nil {
			a.statsTable.Update(a.stats)
		}
	}

	// Control and update FPS
	a.frameRater.Wait()
	a.updateFPS()
}

// UpdateFPS updates the fps value in the window title or header label
func (a *App) updateFPS() {

	// Get the FPS and potential FPS from the frameRater
	fps, _, ok := a.frameRater.FPS(time.Second)
	if !ok {
		return
	}

	a.fps = fps
}

func (a *App) FPS() float64 { return a.fps }
