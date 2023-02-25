package game

func (g *Game) SetAspectRatio(aspectRatio float32) {
	g.aspectRatio = aspectRatio
}

func (g *Game) GetAspectRatio() (aspectRatio float32) {
	return g.aspectRatio
}
