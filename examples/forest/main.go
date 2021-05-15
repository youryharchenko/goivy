package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"

	"math"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/youryharchenko/goivy"
	"github.com/youryharchenko/goivy/pixelgl"
	"golang.org/x/image/colornames"
)

type Environment struct {
	spritesheet pixel.Picture
	treesFrames []pixel.Rect
	//
	batch *pixel.Batch
	//
	cam          pixel.Matrix
	camPos       pixel.Vec
	camSpeed     float64
	camZoom      float64
	camZoomSpeed float64
	//
	frames int
	second <-chan time.Time
	last   time.Time
	dt     float64
}

func (env *Environment) Input(win *pixelgl.Window, config goivy.Config) {

	if win.Pressed(pixelgl.MouseButtonLeft) {
		tree := pixel.NewSprite(env.spritesheet, env.treesFrames[rand.Intn(len(env.treesFrames))])
		mouse := env.cam.Unproject(win.MousePosition())
		tree.Draw(env.batch, pixel.IM.Scaled(pixel.ZV, 4).Moved(mouse))
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

	env.camZoom *= math.Pow(env.camZoomSpeed, win.MouseScroll().Y)
}

func (env *Environment) Run(win *pixelgl.Window, config goivy.Config) {

	env.batch = pixel.NewBatch(&pixel.TrianglesData{}, env.spritesheet)

	for x := env.spritesheet.Bounds().Min.X; x < env.spritesheet.Bounds().Max.X; x += 32 {
		for y := env.spritesheet.Bounds().Min.Y; y < env.spritesheet.Bounds().Max.Y; y += 32 {
			env.treesFrames = append(env.treesFrames, pixel.R(x, y, x+32, y+32))
		}
	}

	for !win.Closed() {
		env.dt = time.Since(env.last).Seconds()
		env.last = time.Now()

		env.cam = pixel.IM.Scaled(env.camPos, env.camZoom).Moved(win.Bounds().Center().Sub(env.camPos))
		win.SetMatrix(env.cam)

		env.Input(win, config)

		win.Clear(colornames.Forestgreen)
		env.batch.Draw(win)
		win.Update()

		env.frames++

		select {
		case <-env.second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", config.WindowConfig.Title, env.frames))
			env.frames = 0
		default:
		}
	}
}

//go:embed trees.png
var b []byte

func main() {

	//spritesheet, err := goivy.LoadPicture("trees.png")
	//if err != nil {
	//	panic(err)
	//}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	env := &Environment{
		spritesheet: pixel.PictureDataFromImage(img),
		//
		camPos:       pixel.ZV,
		camSpeed:     500.0,
		camZoom:      1.0,
		camZoomSpeed: 1.2,
		//
		frames: 0,
		second: time.Tick(time.Second),
		last:   time.Now(),
	}

	goivy.NewScreen(
		goivy.Config{
			WindowConfig: pixelgl.WindowConfig{
				Title:  "Goivy is cool!",
				Bounds: pixel.R(0, 0, 1024, 768),
				VSync:  true,
			},
			BgColor: colornames.Blanchedalmond,
		},
	).Show(env.Run)
}
