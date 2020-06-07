package entities

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var characterDO *ebiten.DrawImageOptions
var xFP float64
var yFP float64
var losse = false

// Snake : Object which the player controls
type Snake struct {
	game     *Game
	numParts int
	lastDir  string
	headImg  ebiten.Image
	tailImg  ebiten.Image
	parts    [][]float64
}

// CreateSnake : Generates a snake
func CreateSnake(g *Game) *Snake {
	s := Snake{
		game:     g,
		numParts: 0,
		lastDir:  "right",
	}

	s.parts = append(s.parts, []float64{300, 300})

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

	xPos, yPos := s.getHeadPos()
	if xPos < 0 || xPos > 580 {
		if !losse {
			s.game.End()
		}

	}
	if yPos < 0 || yPos > 580 {
		if !losse {
			s.game.End()
		}
	}
	return nil
}

// Draw the snake
func (s *Snake) Draw(screen *ebiten.Image, dotTime int) error {
	s.UpdatePos(dotTime)
	characterDO = &ebiten.DrawImageOptions{}

	xPos, yPos := s.getHeadPos()
	characterDO.GeoM.Translate(xPos, yPos)

	screen.DrawImage(&s.headImg, characterDO)

	for i := 0; i < s.numParts; i++ {
		partDO := &ebiten.DrawImageOptions{}
		xPos, yPos := s.getPartPos(i)
		partDO.GeoM.Translate(xPos, yPos)
		screen.DrawImage(&s.tailImg, partDO)
	}

	//ebitenutil.DebugPrint(screen, s.lastDir)

	return nil
}

// UpdatePos changes position values for the snake head
func (s *Snake) UpdatePos(dotTime int) {
	if dotTime == 1 {
		// DEBUG the snake size
		if s.numParts < 3 {
			s.addPoint()
		}
		switch s.lastDir {
		case "up":
			s.translateHeadPos(0, -20)
		case "down":
			s.translateHeadPos(0, +20)
		case "right":
			s.translateHeadPos(20, 0)
		case "left":
			s.translateHeadPos(-20, 0)
		}

	}
}

func (s *Snake) addPoint() {
	s.numParts++
}

func (s *Snake) getHeadPos() (float64, float64) {
	return s.parts[0][0], s.parts[0][1]
}

func (s *Snake) getPartPos(pos int) (float64, float64) {
	return s.parts[pos+1][0], s.parts[pos+1][1]
}

func (s *Snake) translateHeadPos(newXPos, newYPos float64) {
	newX := s.parts[0][0] + newXPos
	newY := s.parts[0][1] + newYPos
	s.updateParts(newX, newY)
}

func (s *Snake) updateParts(newX, newY float64) {
	s.parts = append([][]float64{[]float64{newX, newY}}, s.parts...)
	s.parts = s.parts[:s.numParts+1]
}
