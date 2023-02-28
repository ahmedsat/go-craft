package app

import (
	"fmt"
	"time"
)

// World is the interface that must be satisfied by all demos.
type World interface {
	ID() string
	Start(*App)                 // Called once at the start of the demo
	Update(*App, time.Duration) // Called every frame
	Cleanup(*App)               // Called once at the end of the demo
}

// Worlds maps the demo name string to its object
// Individual demos sets the keys of this map
var worlds = map[string]World{}

func AddWorld(world World) error {

	if _, ok := worlds[world.ID()]; ok {
		return fmt.Errorf("%s world is already added", world.ID())
	}

	worlds[world.ID()] = world
	return nil

}
