package entities

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

var juegar = true

// Game : Main object of the scene. Parent of everything
type Game struct {
	snake   *Snake
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

	return g
}

func (g *Game) End() {
	g.playing = false
	juegar = false
	//fmt.Println(g.playing)
}

// Update the main process of the game
func (g *Game) Update() error {
	//fmt.Println(g.playing)
	if juegar {

		g.dotTime = (g.dotTime + 1) % 25
		if err := g.snake.Update(g.dotTime); err != nil {
			return err
		}
	} else {
		fmt.Println("game stopped")
	}

	return nil
}

// Draw the whole interface
func (g *Game) Draw(screen *ebiten.Image) error {
	if err := g.snake.Draw(screen, g.dotTime); err != nil {
		return err
	}
	return nil
}
