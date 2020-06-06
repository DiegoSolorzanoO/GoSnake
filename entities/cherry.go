package entities

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Cherry : Object which snakes eats

type Cherry struct {
	game   *Game
	cherry ebiten.Image
	xLimit float64
	yLimit float64
	xPos   float64
	yPos   float64
	eaten  bool
}

// CreateCherry : Generates a Cherry
func CreateCherry(g *Game, xCLimit float64, yCLimit float64) *Cherry {
	c := Cherry{
		game:   g,
		xLimit: xCLimit,
		yLimit: yCLimit,
		eaten:  false,
	}

	cherry, _, _ := ebitenutil.NewImageFromFile("assets/cookie.png", ebiten.FilterDefault)

	c.cherry = *cherry

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	c.xPos = random.Float64() * c.xLimit
	c.yPos = random.Float64() * c.yLimit

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(c.xPos)
	fmt.Println(c.yPos)

	return &c
}

// Update : Logical update of the snake
func (c *Cherry) Update(dotTime int) error {
	if c.eaten == false {
		return nil // Return aviso de que ya se la comieron
	}
	return nil
}

// Draw the cherry
func (c *Cherry) Draw(screen *ebiten.Image, dotTime int) error {
	characterDO = &ebiten.DrawImageOptions{}
	characterDO.GeoM.Translate(c.xPos, c.yPos)
	screen.DrawImage(&c.cherry, characterDO)

	return nil
}
