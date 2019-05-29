package osero

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenWidth  = 256
	ScreenHeight = 240
)

type Game struct {
	sceneManager *SceneManager
	input        Input
}

func (g *Game) Update(r *ebiten.Image) error {
	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		g.sceneManager.GoTo(&TitleScene{})
	}

	g.input.Update()
	if err := g.sceneManager.Update(&g.input); err != nil {
		return err
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	g.sceneManager.Draw(r)
	return nil
}