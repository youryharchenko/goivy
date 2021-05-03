package main

import (
	_ "image/png"

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
			for !win.Closed() {
				win.Clear(config.BgColor)

				win.Update()
			}
		},
	)
}
