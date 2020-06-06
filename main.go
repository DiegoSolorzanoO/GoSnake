package main

import (
	"Snakez/entities"
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var gm entities.Game

func init() {
	gm = entities.NewGame()
}

// Game interface of ebiten
type Game struct {
}

// Update the main thread of the game
func (g *Game) Update(screen *ebiten.Image) error {
	if err := gm.Update(); err != nil {
		return err
	}
	return nil
}

// Draw renders the image windows every tick
func (g *Game) Draw(screen *ebiten.Image) {
	if err := gm.Draw(screen); err != nil {
		fmt.Println(err)
	}
}

// Layout : Function which executes when it needs to reajust
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 600
}

func main() {
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Gosnakes")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
