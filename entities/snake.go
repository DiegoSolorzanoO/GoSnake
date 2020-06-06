package entities

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var characterDO *ebiten.DrawImageOptions

// Snake : Object which the player controls
type Snake struct {
	game    *Game
	parts   int
	lastDir string
	headImg ebiten.Image
	tailImg ebiten.Image
	xPos    float64
	yPos    float64
}

// CreateSnake : Generates a snake
func CreateSnake(g *Game) *Snake {
	s := Snake{
		game:    g,
		parts:   0,
		lastDir: "right",
		xPos:    300,
		yPos:    300,
	}

	headimg, _, _ := ebitenutil.NewImageFromFile("assets/playerhead.png", ebiten.FilterDefault)
	tailimg, _, _ := ebitenutil.NewImageFromFile("assets/playertail.png", ebiten.FilterDefault)

	s.headImg = *headimg
	s.tailImg = *tailimg

	return &s
}

// Update : Logical update of the snake
func (s *Snake) Update(dotTime int) error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) && s.lastDir != "right" {
		s.lastDir = "right"
		return nil
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && s.lastDir != "down" {
		s.lastDir = "down"
		return nil
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) && s.lastDir != "up" {
		s.lastDir = "up"
		return nil
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && s.lastDir != "left" {
		s.lastDir = "left"
		return nil
	}
	return nil
}

// Draw the snake
func (s *Snake) Draw(screen *ebiten.Image, dotTime int) error {
	s.UpdatePos(dotTime)
	characterDO = &ebiten.DrawImageOptions{}
	characterDO.GeoM.Translate(s.xPos, s.yPos)
	screen.DrawImage(&s.headImg, characterDO)

	ebitenutil.DebugPrint(screen, s.lastDir)

	return nil
}

// UpdatePos changes position values for the snake head
func (s *Snake) UpdatePos(dotTime int) {
	if dotTime == 1 {
		switch s.lastDir {
		case "up":
			s.yPos -= 20
		case "down":
			s.yPos += 20
		case "right":
			s.xPos += 20
		case "left":
			s.xPos -= 20
		}
	}
}
