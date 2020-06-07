package entities

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/basicfont"
)

// Hud for the game
type Hud struct {
	game   *Game
	points int
}

// CreateHud : Constructor
func CreateHud(g *Game) *Hud {
	h := Hud{
		game:   g,
		points: 0,
	}

	return &h
}

func (h *Hud) addPoint() {
	h.points++
}

func textDimension(text string) (w int, h int) {
	return 7 * len(text), 13
}

// Draw the hud
func (h *Hud) Draw(screen *ebiten.Image) error {
	text.Draw(screen, "Score: "+strconv.Itoa(h.points), basicfont.Face7x13, 20, 20, color.White)

	if !h.game.playing {
		goText := "GAME OVER"
		textW, textH := textDimension(goText)
		screenW := screen.Bounds().Dx()
		screenH := screen.Bounds().Dy()

		text.Draw(screen, goText, basicfont.Face7x13, screenW/2-textW/2, screenH/2+textH/2, color.White)
	}

	return nil
}
