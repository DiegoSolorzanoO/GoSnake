package main

import (
	_ "image/png"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var imgC *ebiten.Image
var imgUW *ebiten.Image
var imgDW *ebiten.Image
var imgLW *ebiten.Image
var imgRW *ebiten.Image
var direction string

func init() {
	var err error
	imgC, _, err = ebitenutil.NewImageFromFile("pokeball.png", ebiten.FilterDefault) //aqui se cargan las imagenes
	if err != nil {
		log.Fatal(err)
	}
	imgUW, _, err = ebitenutil.NewImageFromFile("wall.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgDW, _, err = ebitenutil.NewImageFromFile("wall.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgLW, _, err = ebitenutil.NewImageFromFile("sidewall.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgRW, _, err = ebitenutil.NewImageFromFile("sidewall.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	count   int
	pressed []ebiten.Key
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.pressed = nil
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ { //este metodo tambien checa cuando se presiona una key
		if ebiten.IsKeyPressed(k) {
			g.pressed = append(g.pressed, k)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}   //puntos para poner los sprites
	opUW := &ebiten.DrawImageOptions{} //upwall
	opDW := &ebiten.DrawImageOptions{} //downwall
	opRW := &ebiten.DrawImageOptions{} //rightwall
	//op.GeoM.Translate(-float64(640)/2, -float64(640)/2)
	//op.GeoM.Rotate(float64(g.count%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(640/2, 640/2)
	opUW.GeoM.Translate(0, 0)
	opDW.GeoM.Translate(0, 620)
	opRW.GeoM.Translate(620, 0)
	screen.DrawImage(imgC, op)   //se dibuja la imagen con respecto al punto
	screen.DrawImage(imgUW, nil) // respecto a 0,0
	screen.DrawImage(imgDW, opDW)
	screen.DrawImage(imgLW, nil)
	screen.DrawImage(imgRW, opRW)
	keyStrs := []string{} //aqui se van a guardar las keys presionadas

	for _, p := range g.pressed {
		keyStrs = append(keyStrs, p.String())
	}
	ebitenutil.DebugPrint(screen, strings.Join(keyStrs, ", ")) //aqui se imprimen en la ventana
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func main() {
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Gosnakes")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
