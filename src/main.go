package main

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	score       int
	pickPositon int //rotation
	pinPosition int // 8 states - 45 deg appart
	debugText   string
	debugshow   bool
	// window size
	windowWidth        int
	windowCenterWidth  int
	windowHeight       int
	windowCenterHeight int
	windowUnit         int
}

func (g *Game) Update() error {
	// Debug
	g.debugText = ""
	// window size
	g.debugText += "Window:\nSize " + strconv.Itoa(g.windowHeight) + "x" + strconv.Itoa(g.windowWidth) + "\nUnit " + strconv.Itoa(g.windowUnit) + "\n"

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.debugshow {
			g.debugshow = false
		} else {

			g.debugshow = true
		}
	}

	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	if g.debugshow {
		ebitenutil.DebugPrint(screen, g.debugText)
	}
	// TODO draw lock
	vector.DrawFilledCircle(screen, float32(g.windowCenterWidth), float32(g.windowUnit)*2.5, float32(g.windowUnit)/2, color.White, true)
	vector.DrawFilledCircle(screen, float32(g.windowCenterWidth), float32(g.windowUnit)*1.5, float32(g.windowUnit)/2, color.White, true)
	// TODO draw score - with text.draw() - needs font
}
func (g *Game) Layout(outWidth, outHeight int) (screenWidth, screenHeight int) {
	g.windowWidth = outWidth
	g.windowHeight = outHeight

	g.windowCenterWidth = outWidth / 2
	g.windowCenterHeight = outHeight / 2

	// find smallest window unit
	if g.windowHeight/3 < g.windowWidth {
		g.windowUnit = g.windowHeight / 3
	} else {
		g.windowUnit = g.windowWidth
	}

	return outWidth, outHeight
}

func main() {

	ebiten.SetWindowTitle("Pop The Lock")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	game := &Game{}
	game.debugshow = true //TODO remove
	ebiten.RunGame(game)
}
