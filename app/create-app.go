package app

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util"
	"github.com/g3n/engine/util/logger"
	"github.com/g3n/engine/util/stats"
	"github.com/g3n/engine/window"
)

func Create() (a *App) {
	a = new(App)

	a.world = worlds["worlds.test"]

	a.Application = app.App()

	// Creates application logger
	a.log = logger.New("GC", nil)
	a.log.AddWriter(logger.NewConsole(false))
	a.log.SetFormat(logger.FTIME | logger.FMICROS)
	a.log.SetLevel(logger.DEBUG)

	a.log.Info("Go Craft is Running ...")

	a.stats = stats.NewStats(a.Gls())

	// Log OpenGL version
	glVersion := a.Gls().GetString(gls.VERSION)
	a.log.Info("OpenGL version: %s", glVersion)

	// Set OpenGL error checking based on flag
	a.Gls().SetCheckErrors(glErrors)

	// Create scene
	a.scene = core.NewNode()
	gui.Manager().Set(a.Scene())

	// Create camera and orbit control
	width, height := a.GetSize()
	aspect := float32(width) / float32(height)
	a.camera = camera.New(aspect)
	a.scene.Add(a.camera) // Add camera to scene (important for audio demos)

	// Create and add ambient light to scene
	a.ambLight = light.NewAmbient(&math32.Color{R: 1.0, G: 1.0, B: 1.0}, 0.5)
	a.scene.Add(a.ambLight)

	// Create frame rater
	a.frameRater = util.NewFrameRater(targetFPS)

	// Sets the default window resize event handler
	a.Subscribe(window.OnWindowSize, func(evname string, ev interface{}) { a.OnWindowResize() })
	a.OnWindowResize()

	// Subscribe to key events
	a.Subscribe(window.OnKeyDown, func(evname string, ev interface{}) {
		kev := ev.(*window.KeyEvent)
		if kev.Key == window.KeyEscape { // ESC terminates the program
			a.Exit()
		}
		if kev.Key == window.KeyS && kev.Mods == window.ModAlt { // Ctr-S prints statistics in the console
			a.logStats()
		}
	})

	// Setup scene
	a.setupScene()

	a.world.Start(a)

	return
}
