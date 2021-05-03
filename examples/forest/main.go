package main

import (
	"fmt"
	_ "image/png"
	"math"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/youryharchenko/goivy"
	"golang.org/x/image/colornames"
)

func main() {

	goivy.NewScreen(
		goivy.Config{
			WindowConfig: pixelgl.WindowConfig{
				Title:  "Goivy is cool!",
				Bounds: pixel.R(0, 0, 1024, 768),
				VSync:  true,
			},
			BgColor: colornames.Blanchedalmond,
		},
	).Show(
		func(win *pixelgl.Window, config goivy.Config) {

			spritesheet, err := goivy.LoadPicture("trees.png")
			if err != nil {
				panic(err)
			}

			batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)

			var treesFrames []pixel.Rect
			for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
				for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
					treesFrames = append(treesFrames, pixel.R(x, y, x+32, y+32))
				}
			}
			var (
				camPos       = pixel.ZV
				camSpeed     = 500.0
				camZoom      = 1.0
				camZoomSpeed = 1.2
			)

			var (
				frames = 0
				second = time.Tick(time.Second)
				last   = time.Now()
			)

			for !win.Closed() {
				dt := time.Since(last).Seconds()
				last = time.Now()

				cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
				win.SetMatrix(cam)

				if win.Pressed(pixelgl.MouseButtonLeft) {
					tree := pixel.NewSprite(spritesheet, treesFrames[rand.Intn(len(treesFrames))])
					mouse := cam.Unproject(win.MousePosition())
					tree.Draw(batch, pixel.IM.Scaled(pixel.ZV, 4).Moved(mouse))
				}
				if win.Pressed(pixelgl.KeyLeft) {
					camPos.X -= camSpeed * dt
				}
				if win.Pressed(pixelgl.KeyRight) {
					camPos.X += camSpeed * dt
				}
				if win.Pressed(pixelgl.KeyDown) {
					camPos.Y -= camSpeed * dt
				}
				if win.Pressed(pixelgl.KeyUp) {
					camPos.Y += camSpeed * dt
				}

				camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

				win.Clear(colornames.Forestgreen)
				batch.Draw(win)
				win.Update()

				frames++

				select {
				case <-second:
					win.SetTitle(fmt.Sprintf("%s | FPS: %d", config.WindowConfig.Title, frames))
					frames = 0
				default:
				}
			}
		},
	)
}
