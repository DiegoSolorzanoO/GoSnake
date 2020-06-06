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
	// var err error
	// imgC, _, err = ebitenutil.NewImageFromFile("assets/playerhead.png", ebiten.FilterDefault) //aqui se cargan las imagenes
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// imgUW, _, err = ebitenutil.NewImageFromFile("assets/wall.png", ebiten.FilterDefault)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// imgDW, _, err = ebitenutil.NewImageFromFile("assets/wall.png", ebiten.FilterDefault)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// imgLW, _, err = ebitenutil.NewImageFromFile("assets/sidewall.png", ebiten.FilterDefault)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// imgRW, _, err = ebitenutil.NewImageFromFile("assets/sidewall.png", ebiten.FilterDefault)
	// if err != nil {
	// 	log.Fatal(err)
	// }
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
	// g.pressed = nil
	// for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ { //este metodo tambien checa cuando se presiona una key
	// 	if ebiten.IsKeyPressed(k) {
	// 		g.pressed = append(g.pressed, k)
	// 	}
	// }
	// return nil
}

// Draw renders the image windows every tick
func (g *Game) Draw(screen *ebiten.Image) {
	if err := gm.Draw(screen); err != nil {
		fmt.Println(err)
	}
	// 	op := &ebiten.DrawImageOptions{}   //puntos para poner los sprites
	// 	opUW := &ebiten.DrawImageOptions{} //upwall
	// 	opDW := &ebiten.DrawImageOptions{} //downwall
	// 	opRW := &ebiten.DrawImageOptions{} //rightwall
	// 	//op.GeoM.Translate(-float64(640)/2, -float64(640)/2)
	// 	//op.GeoM.Rotate(float64(g.count%360) * 2 * math.Pi / 360)
	// 	g.count++
	// 	op.GeoM.Translate(float64(g.count), 640/2)
	// 	opUW.GeoM.Translate(0, 0)
	// 	opDW.GeoM.Translate(0, 620)
	// 	opRW.GeoM.Translate(620, 0)
	// 	screen.DrawImage(imgC, op)   //se dibuja la imagen con respecto al punto
	// 	screen.DrawImage(imgUW, nil) // respecto a 0,0
	// 	screen.DrawImage(imgDW, opDW)
	// 	screen.DrawImage(imgLW, nil)
	// 	screen.DrawImage(imgRW, opRW)
	// 	keyStrs := []string{} //aqui se van a guardar las keys presionadas

	// 	for _, p := range g.pressed {
	// 		keyStrs = append(keyStrs, p.String())
	// 	}
	// 	ebitenutil.DebugPrint(screen, "Hola") //aqui se imprimen en la ventana
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
