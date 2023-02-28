// This is a simple model for your tests
package worlds

import (
	"time"

	"github.com/ahmedsat/go-craft/app"
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

	// Show axis helper
	ah := helper.NewAxes(1.0)
	t.Scene().Add(ah)

	// Creates a grid helper and saves its pointer in the test state
	t.grid = helper.NewGrid(50, 1, &math32.Color{0.4, 0.4, 0.4})
	t.Scene().Add(t.grid)

	// Changes the camera position
	t.cameraPos.Set(0, 4, 10)
	t.cameraTarget.Set(0, 0, 0)
	t.cameraUp.Set(0, 1, 0)
	t.updateCamera()

	t.playerController()

}

// This method will be called at every frame
// You can animate your objects here.
func (t *testingWorld) Update(a *app.App, deltaTime time.Duration) {

	// Rotate the grid, just for show.
	// rps := float32(deltaTime.Seconds()) * 2 * math32.Pi
	// t.grid.RotateY(rps * 0.05)
}

// Cleanup is called once at the end of the demo.
func (t *testingWorld) Cleanup(a *app.App) {}

func (t *testingWorld) updateCamera() {
	t.Camera().SetPositionVec(&t.cameraPos)
	t.Camera().LookAt(&t.cameraTarget, &t.cameraUp)
}

func (t *testingWorld) playerController() {

	moves := func(s string, i interface{}) {
		kev := i.(*window.KeyEvent)
		if kev.Key == window.KeyUp || kev.Key == window.KeyW {
			t.cameraPos.Z--
		}
		if kev.Key == window.KeyDown || kev.Key == window.KeyS {
			t.cameraPos.Z++
		}
		if kev.Key == window.KeyRight || kev.Key == window.KeyD {
			t.cameraPos.X++
		}
		if kev.Key == window.KeyLeft || kev.Key == window.KeyA {
			t.cameraPos.X--
		}

		t.updateCamera()
	}

	t.Subscribe(window.OnKeyRepeat, moves)
	t.Subscribe(window.OnKeyDown, moves)

}
