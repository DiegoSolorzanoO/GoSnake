package entities

import (
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Game : Main object of the scene. Parent of everything
type Game struct {
	snake       *Snake
	hud         *Hud
	cherries    []*Cherry
	numCherries int
	enemies     []*EnemySnake
	playing     bool
	points      int
	dotTime     int
}

// NewGame : Starts a new game assigning variables
func NewGame(cherrys int) Game {
	g := Game{
		playing:     true,
		points:      0,
		dotTime:     0,
		numCherries: cherrys,
	}
	arrayC := make([]*Cherry, g.numCherries)
	for i := 0; i < g.numCherries; i++ {
		arrayC[i] = CreateCherry(&g)
		time.Sleep(20)
	}
	arrayEnemies := make([]*EnemySnake, 5)
	for i := 0; i < len(arrayEnemies); i++ {
		arrayEnemies[i] = CreateEnemySnake(&g)
		time.Sleep(20)
	}
	g.cherries = arrayC
	g.enemies = arrayEnemies
	g.snake = CreateSnake(&g)
	g.hud = CreateHud(&g)

	return g
}

// End the game
func (g *Game) End() {
	g.playing = false
}

// Update the main process of the game
func (g *Game) Update() error {

	if g.playing {
		if g.numCherries == 0 {
			g.playing = false
		}

		g.dotTime = (g.dotTime + 1) % 20
		if err := g.snake.Update(g.dotTime); err != nil {
			return err
		}
		for _, enemy := range g.enemies {
			if err := enemy.Update(g.dotTime); err != nil {
				return err
			}
		}
		xPos, yPos := g.snake.getHeadPos()
		for i := 0; i < len(g.cherries); i++ {
			if xPos == g.cherries[i].xPos && yPos == g.cherries[i].yPos {
				g.cherries[i].yPos = -20
				g.cherries[i].xPos = -20
				g.hud.addPoint()
				g.numCherries--
				g.snake.addPoint()
				break
			}
		}

	} else {
		//fmt.Println("game stopped")

	}

	for i := 0; i < g.numCherries; i++ {
		if err := g.cherries[i].Update(g.dotTime); err != nil {
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
	for _, enemy := range g.enemies {
		if err := enemy.Draw(screen, g.dotTime); err != nil {
			return err
		}
	}
	if err := g.hud.Draw(screen); err != nil {
		return err
	}
	for i := 0; i < len(g.cherries); i++ {
		if err := g.cherries[i].Draw(screen, g.dotTime); err != nil {
			return err
		}
	}
	if !g.playing && g.numCherries != 0 {
		ebitenutil.DebugPrint(screen, "Game over")
	}

	if !g.playing && g.numCherries == 0 {
		ebitenutil.DebugPrint(screen, "You Win")
	}

	return nil
}
