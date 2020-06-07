package entities

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

var juegar = true

// Game : Main object of the scene. Parent of everything
type Game struct {
	snake   *Snake
	cherry  []*Cherry
	cherryN int
	playing bool
	points  int
	dotTime int
}

// NewGame : Starts a new game assigning variables
func NewGame(cherrys int) Game {
	g := Game{
		playing: true,
		points:  0,
		dotTime: 0,
	}
	g.snake = CreateSnake(&g)
	g.cherryN = cherrys
	arrayC := make([]*Cherry, g.cherryN)
	for i := 0; i < g.cherryN; i++ {
		arrayC[i] = CreateCherry(&g)
		time.Sleep(20)
	}
	g.cherry = arrayC

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
		//fmt.Println("game stopped")
	}

	for i := 0; i < g.cherryN; i++ {
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
	for i := 0; i < g.cherryN; i++ {
		if err := g.cherry[i].Draw(screen, g.dotTime); err != nil {
			return err
		}
	}

	return nil
}
