package game

import (
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
)

type Scene struct {
	*core.Node
	Camera *camera.Camera
}

func (s *Scene) Init(callback func(*Scene)) {
	s.Node = core.NewNode()

	// Create perspective camera
	s.Camera = camera.New(1)
	s.Camera.SetPosition(0, 0, 3)
	s.Add(s.Camera)

	// Set up orbit control for the camera
	camera.NewOrbitControl(s.Camera)

	callback(s)
}
