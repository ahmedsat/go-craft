package gocraft

import (
	"github.com/ahmedsat/go-craft/game"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/helper"
)

type Gocraft struct {
	game.Game
}

func NewGame() (g Gocraft) {

	g.Game.Init(g.Init)

	return
}

func (g *Gocraft) Init(s *game.Scene) {

	// Create a blue torus and add it to the scene
	geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	mat := material.NewStandard(math32.NewColor("DarkBlue"))
	mesh := graphic.NewMesh(geom, mat)
	s.Add(mesh)

	// Create and add a button to the scene
	btn := gui.NewButton("Make Green")
	btn.SetPosition(100, 40)
	btn.SetSize(40, 40)
	btn.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		mat.SetColor(math32.NewColor("DarkGreen"))
	})
	s.Add(btn)

	// Create and add lights to the scene
	s.Add(light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8))
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	s.Add(pointLight)

	// Create and add an axis helper to the scene
	s.Add(helper.NewAxes(0.5))

}

func (g *Gocraft) Run() {
	g.Game.Run()
}
