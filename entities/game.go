package entities

import (
	"github.com/hajimehoshi/ebiten"
)

// Game : Main object of the scene. Parent of everything
type Game struct {
	snake   *Snake
	cherry  []*Cherry
	playing bool
	points  int
	dotTime int
}

// NewGame : Starts a new game assigning variables
func NewGame() Game {
	g := Game{
		playing: true,
		points:  0,
		dotTime: 0,
	}
	g.snake = CreateSnake(&g)
	arrayC := make([]*Cherry, 5)
	for i := 0; i < 5; i++ {
		arrayC[i] = CreateCherry(&g, 600, 600)
	}
	g.cherry = arrayC

	return g
}

// Update the main process of the game
func (g *Game) Update() error {
	g.dotTime = (g.dotTime + 1) % 25
	if err := g.snake.Update(g.dotTime); err != nil {
		return err
	}
	for i := 0; i < 5; i++ {
		if err := g.cherry[i].Update(g.dotTime); err != nil {
			return err
		}
	}

	return nil
}

// Draw the whole interface
func (g *Game) Draw(screen *ebiten.Image) error {
	if err := g.snake.Draw(screen, g.dotTime); err != nil {
		return err
	}
	for i := 0; i < 5; i++ {
		if err := g.cherry[i].Draw(screen, g.dotTime); err != nil {
			return err
		}
	}

	return nil
}
