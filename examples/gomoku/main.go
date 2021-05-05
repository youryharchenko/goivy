package main

import (
	_ "embed"
	"fmt"
	_ "image/png"
	"log"

	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/youryharchenko/goivy"
	"golang.org/x/image/colornames"
)

type Environment struct {
	//
	//batch *pixel.Batch
	//
	imd *imdraw.IMDraw
	//
	cam          pixel.Matrix
	camPos       pixel.Vec
	camSpeed     float64
	camZoom      float64
	camZoomSpeed float64
	//
	frames int
	second <-chan time.Time
	//milisec <-chan time.Time
	last time.Time
	dt   float64
	//
	d       float64
	dim     int
	cv      []float64
	ch      []float64
	changed bool
	//
	steps []Step
}

func (env *Environment) Draw(win *pixelgl.Window, config goivy.Config) {
	if env.changed {
		env.imd = imdraw.New(nil)
		l := float64(env.dim) * env.d / 2
		for _, y := range env.cv {
			env.imd.Color = colornames.Black
			env.imd.EndShape = imdraw.RoundEndShape
			env.imd.Push(pixel.V(-l, y))
			env.imd.EndShape = imdraw.RoundEndShape
			env.imd.Push(pixel.V(l, y))
			env.imd.Line(2)
		}
		for _, x := range env.ch {
			env.imd.Color = colornames.Black
			env.imd.EndShape = imdraw.RoundEndShape
			env.imd.Push(pixel.V(x, -l))
			env.imd.EndShape = imdraw.RoundEndShape
			env.imd.Push(pixel.V(x, l))
			env.imd.Line(2)
		}
		for range env.steps {
			env.imd.Color = colornames.Black
			env.imd.Push(pixel.V(0, 0))
			env.imd.Circle(env.d/2, 0)
		}
	}

	win.Clear(config.BgColor)
	env.imd.Draw(win)
	win.Update()

	env.frames++
	env.changed = false

}

func (env *Environment) Input(win *pixelgl.Window, config goivy.Config) {

	if win.JustReleased(pixelgl.MouseButtonLeft) {
		log.Println("MouseButtonLeft JustReleased", env.cam.Unproject(win.MousePosition()))
		//tree := pixel.NewSprite(env.spritesheet, env.treesFrames[rand.Intn(len(env.treesFrames))])
		//mouse := env.cam.Unproject(win.MousePosition())
		//tree.Draw(env.batch, pixel.IM.Scaled(pixel.ZV, 4).Moved(mouse))
		env.changed = true
	}

	if win.Pressed(pixelgl.KeyLeft) {
		env.camPos.X -= env.camSpeed * env.dt
	}

	if win.Pressed(pixelgl.KeyRight) {
		env.camPos.X += env.camSpeed * env.dt
	}

	if win.Pressed(pixelgl.KeyDown) {
		env.camPos.Y -= env.camSpeed * env.dt
	}

	if win.Pressed(pixelgl.KeyUp) {
		env.camPos.Y += env.camSpeed * env.dt
	}

	scroll := win.MouseScroll().Y
	if int(scroll) != 0 {
		env.camZoom *= math.Pow(env.camZoomSpeed, win.MouseScroll().Y)
		//env.changed = true
	}
}

func (env *Environment) Run(win *pixelgl.Window, config goivy.Config) {

	//env.batch = pixel.NewBatch(&pixel.TrianglesData{}, env.spritesheet)

	//for x := env.spritesheet.Bounds().Min.X; x < env.spritesheet.Bounds().Max.X; x += 32 {
	//	for y := env.spritesheet.Bounds().Min.Y; y < env.spritesheet.Bounds().Max.Y; y += 32 {
	//		env.treesFrames = append(env.treesFrames, pixel.R(x, y, x+32, y+32))
	//	}
	//}

	env.changed = true

	for !win.Closed() {
		env.dt = time.Since(env.last).Seconds()
		env.last = time.Now()

		env.cam = pixel.IM.Scaled(env.camPos, env.camZoom).Moved(win.Bounds().Center().Sub(env.camPos))
		win.SetMatrix(env.cam)

		env.Input(win, config)

		win.SetSmooth(true)
		env.Draw(win, config)

		select {
		case <-env.second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", config.WindowConfig.Title, env.frames))
			env.frames = 0
		default:
		}
	}
}

func (env *Environment) fillCoord() {
	c := make([]float64, env.dim)
	cntr := env.dim / 2
	c[cntr] = 0
	for i := cntr + 1; i < env.dim; i++ {
		c[i] = c[i-1] + env.d
	}
	for i := cntr - 1; i >= 0; i-- {
		c[i] = c[i+1] - env.d
	}
	env.cv = c
	env.ch = c
}

func main() {

	env := &Environment{
		imd: imdraw.New(nil),
		//
		camPos:       pixel.ZV,
		camSpeed:     500.0,
		camZoom:      1.0,
		camZoomSpeed: 1.2,
		//
		frames: 0,
		second: time.Tick(time.Second),
		last:   time.Now(),
		//
		d:   40,
		dim: 15,
		//
		steps: []Step{{7, 7}},
	}

	env.fillCoord()

	goivy.NewScreen(
		goivy.Config{
			WindowConfig: pixelgl.WindowConfig{
				Title:  "Play gomoku!",
				Bounds: pixel.R(0, 0, 1024, 768),
				VSync:  true,
			},
			BgColor: colornames.Lightgrey,
		},
	).Show(env.Run)
}