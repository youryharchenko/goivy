package goivy

import (
	"image/color"

	"github.com/faiface/pixel/pixelgl"
)

type Screen struct {
	config Config
}

type Config struct {
	WindowConfig pixelgl.WindowConfig
	BgColor      color.Color
}

func NewScreen(config Config) (screen *Screen) {
	screen = &Screen{
		config: config,
	}
	return
}

func (screen *Screen) Show(run func(win *pixelgl.Window, config Config)) {
	pixelgl.Run(func() {
		win, err := pixelgl.NewWindow(screen.config.WindowConfig)
		if err != nil {
			panic(err)
		}

		win.SetSmooth(true)

		run(win, screen.config)
	})
}
