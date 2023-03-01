// This is a simple model for your tests
package worlds

import (
	"fmt"
	"time"

	"github.com/ahmedsat/go-craft/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
)

// This is your test object. You can store state here.
// By convention and to avoid conflict with other demo/tests name it
// using your test category and name.
type testingWorld struct {
	*app.App

	worldID string

	grid *helper.Grid // Pointer to a GridHelper created in 'Start'

	cameraPos    math32.Vector3
	cameraTarget math32.Vector3
	cameraUp     math32.Vector3
	speed        float32

	coordinateLabel *gui.Label
	fpsLabel        *gui.Label
}

func (t *testingWorld) ID() string {
	return t.worldID
}

// This method will be called once when the test is selected from the G3ND list.
// 'a' is a pointer to the G3ND application.
// It allows access to several methods such as a.Scene(), which returns the current scene,
// a.DemoPanel(), a.Camera(), a.Window() among others.
// You can build your scene adding your objects to the a.Scene()
func (t *testingWorld) Start(a *app.App) {

	t.App = a

	t.speed = 10

	t.coordinateLabel = gui.NewLabel("")
	t.coordinateLabel.SetPosition(40, 60)
	t.coordinateLabel.SetSize(40, 40)
	t.coordinateLabel.SetColor(math32.NewColor("Black"))
	t.Scene().Add(t.coordinateLabel)

	t.fpsLabel = gui.NewLabel("")
	t.fpsLabel.SetPosition(40, 40)
	t.fpsLabel.SetSize(40, 40)
	t.fpsLabel.SetColor(math32.NewColor("Black"))
	t.Scene().Add(t.fpsLabel)

	t.updateGUI()

	// Show axis helper
	ah := helper.NewAxes(1.0)
	t.Scene().Add(ah)

	// Creates a grid helper and saves its pointer in the test state
	t.grid = helper.NewGrid(50, 1, &math32.Color{R: 0.4, G: 0.4, B: 0.4})
	t.Scene().Add(t.grid)

	// Changes the camera position
	t.cameraPos.Set(0, 4, 10)
	t.cameraTarget.Set(0, 0, 0)
	t.cameraUp.Set(0, 1, 0)
	camera.NewOrbitControl(t.Camera())
	t.updateCamera()

	t.playerController()

}

// This method will be called at every frame
// You can animate your objects here.
func (t *testingWorld) Update(a *app.App, deltaTime time.Duration) {

	// Rotate the grid, just for show.
	// rps := float32(deltaTime.Seconds()) * 2 * math32.Pi
	// t.grid.RotateY(rps * 0.05)
	t.updateGUI()
}

// Cleanup is called once at the end of the demo.
func (t *testingWorld) Cleanup(a *app.App) {}

func (t *testingWorld) updateCamera() {
	t.Camera().SetPositionVec(&t.cameraPos)
	t.Camera().LookAt(&t.cameraTarget, &t.cameraUp)

}

func (t *testingWorld) updateGUI() {
	t.coordinateLabel.SetText(fmt.Sprintf("(%v,%v,%v)", t.cameraPos.X, t.cameraPos.Y, t.cameraPos.Z))
	t.fpsLabel.SetText(fmt.Sprint(t.FPS()))
}

func (t *testingWorld) playerController() {
	moves := func(s string, i interface{}) {
		if t.FPS() < 1 {
			return
		}
		kev := i.(*window.KeyEvent)
		if kev.Key == window.KeyUp || kev.Key == window.KeyW {
			t.cameraPos.Z += t.speed * 1 / float32(t.FPS())
		}
		if kev.Key == window.KeyDown || kev.Key == window.KeyS {
			t.cameraPos.Z -= t.speed * 1 / float32(t.FPS())
		}
		if kev.Key == window.KeyRight || kev.Key == window.KeyD {
			t.cameraPos.X += t.speed * 1 / float32(t.FPS())
		}
		if kev.Key == window.KeyLeft || kev.Key == window.KeyA {
			t.cameraPos.X -= t.speed * 1 / float32(t.FPS())
		}
		if kev.Key == window.KeyRightShift || kev.Key == window.KeyPageUp {
			t.cameraPos.Y += t.speed * 1 / float32(t.FPS())
		}
		if kev.Key == window.KeyLeftShift || kev.Key == window.KeyPageDown {
			t.cameraPos.Y -= t.speed * 1 / float32(t.FPS())
		}

		t.updateCamera()
	}

	t.Subscribe(window.OnKeyRepeat, moves)
	t.Subscribe(window.OnKeyDown, moves)

}
