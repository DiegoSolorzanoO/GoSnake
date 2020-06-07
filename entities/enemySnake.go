package entities

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// EnemySnake : Snake object for enemies
type EnemySnake struct {
	game     *Game
	numParts int
	lastDir  string
	headImg  ebiten.Image
	tailImg  ebiten.Image
	parts    [][]float64
	seed     rand.Source
}

// CreateEnemySnake : Generates an enemy snake
func CreateEnemySnake(g *Game) *EnemySnake {
	s := EnemySnake{
		game:     g,
		numParts: 0,
		lastDir:  "right",
	}

	s.seed = rand.NewSource(time.Now().UnixNano())
	random := rand.New(s.seed)
	iniX := float64(random.Intn(30) * 20)
	iniY := float64(random.Intn(30) * 20)

	s.parts = append(s.parts, []float64{iniX, iniY})

	headimg, _, _ := ebitenutil.NewImageFromFile("assets/enemyhead.png", ebiten.FilterDefault)
	tailimg, _, _ := ebitenutil.NewImageFromFile("assets/enemytail.png", ebiten.FilterDefault)

	s.headImg = *headimg
	s.tailImg = *tailimg

	return &s
}

// Update : Logical update of the snake
func (s *EnemySnake) Update(dotTime int) error {
	if dotTime == 1 {
		random := rand.New(s.seed)
		action := random.Intn(4)
		changingDirection := random.Intn(2)
		posX, posY := s.getHeadPos()
		if changingDirection == 1 {
			switch action {
			case 0:
				if posX < 580 {
					s.lastDir = "right"
				}
				return nil
			case 1:
				if posY < 580 {
					s.lastDir = "down"
				}
				return nil
			case 2:
				if posY > 0 {
					s.lastDir = "up"
				}
				return nil
			case 3:
				if posX > 0 {
					s.lastDir = "left"
				}
				return nil
			}
		}
		if posX >= 580 {
			s.lastDir = "left"
			return nil
		}
		if posX == 0 {
			s.lastDir = "right"
			return nil
		}
		if posY == 580 {
			s.lastDir = "up"
			return nil
		}
		if posY == 0 {
			s.lastDir = "down"
			return nil
		}
	}

	// if dotTime == 1 {
	// 	xPos, yPos := s.getHeadPos()
	// 	if xPos < 0 || xPos > 580 || yPos < 0 || yPos > 580 || s.collisionWithHimself() {
	// 		s.game.End()
	// 	}
	// }
	return nil
}

// Draw the snake
func (s *EnemySnake) Draw(screen *ebiten.Image, dotTime int) error {
	if s.game.playing {
		s.UpdatePos(dotTime)
	}
	enemyDO := &ebiten.DrawImageOptions{}

	xPos, yPos := s.getHeadPos()
	enemyDO.GeoM.Translate(xPos, yPos)

	screen.DrawImage(&s.headImg, enemyDO)

	for i := 0; i < s.numParts; i++ {
		partDO := &ebiten.DrawImageOptions{}
		xPos, yPos := s.getPartPos(i)
		partDO.GeoM.Translate(xPos, yPos)
		screen.DrawImage(&s.tailImg, partDO)
	}

	return nil
}

// UpdatePos changes position values for the snake head
func (s *EnemySnake) UpdatePos(dotTime int) {
	if dotTime == 1 {
		if s.numParts < 7 {
			s.numParts++
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

func (s *EnemySnake) getHeadPos() (float64, float64) {
	return s.parts[0][0], s.parts[0][1]
}

func (s *EnemySnake) getPartPos(pos int) (float64, float64) {
	return s.parts[pos+1][0], s.parts[pos+1][1]
}

func (s *EnemySnake) translateHeadPos(newXPos, newYPos float64) {
	newX := s.parts[0][0] + newXPos
	newY := s.parts[0][1] + newYPos
	s.updateParts(newX, newY)
}

func (s *EnemySnake) updateParts(newX, newY float64) {
	s.parts = append([][]float64{[]float64{newX, newY}}, s.parts...)
	s.parts = s.parts[:s.numParts+1]
}

func (s *EnemySnake) collisionWithHimself() bool {
	posX, posY := s.getHeadPos()
	for i := 1; i < len(s.parts); i++ {
		if posX == s.parts[i][0] && posY == s.parts[i][1] {
			return true
		}
	}
	return false
}
