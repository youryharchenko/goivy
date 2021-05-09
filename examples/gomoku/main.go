package main

import (
	_ "embed"
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"

	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"github.com/youryharchenko/goivy"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

//go:embed 	Itim-Regular.ttf
var fivo []byte

type Environment struct {
	//
	fontFace font.Face
	//
	//batch *pixel.Batch
	//
	imd        *imdraw.IMDraw
	atlas      *text.Atlas
	txt        *text.Text
	head       *text.Text
	headText   string
	status     *text.Text
	statusText string
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
	steps      []Step
	colors     [2]color.Color
	colorNames [2]string
	//
	play bool
	calc bool
	auto bool
}

func (env *Environment) Draw(win *pixelgl.Window, config goivy.Config) {
	if env.changed {
		env.imd = imdraw.New(nil)
		l := float64(env.dim) * env.d / 2

		env.head.Clear()
		env.head.Color = colornames.Black
		env.head.Dot.X = -env.txt.BoundsOf(env.headText).W() / 2
		env.head.Dot.Y = l + env.txt.BoundsOf(env.headText).H()*2
		fmt.Fprint(env.head, env.headText)

		env.status.Clear()
		env.status.Color = colornames.Black
		env.status.Dot.X = -env.txt.BoundsOf(env.statusText).W() / 2
		env.status.Dot.Y = -(l + env.txt.BoundsOf(env.statusText).H()*2)
		fmt.Fprint(env.status, env.statusText)

		env.imd.Clear()

		env.imd.Color = colornames.Black
		env.imd.Push(pixel.V(-l, -l))
		env.imd.Push(pixel.V(l, l))
		env.imd.Rectangle(3)

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

		for i, st := range env.steps {
			x := env.ch[st.x]
			y := env.ch[st.y]
			env.imd.Color = colornames.Black
			env.imd.Push(pixel.V(x, y))
			env.imd.Circle(env.d/2, 0)

			c := i % 2
			env.imd.Color = env.colors[c]
			env.imd.Push(pixel.V(x, y))
			env.imd.Circle(env.d/2-1, 0)
		}

		env.txt.Clear()
		for i, st := range env.steps {
			x := env.ch[st.x]
			y := env.ch[st.y]

			c := 1 - i%2

			s := fmt.Sprintf("%d", i+1)
			env.txt.Color = env.colors[c]
			env.txt.Dot.X = x - env.txt.BoundsOf(s).W()/2
			env.txt.Dot.Y = y - env.txt.BoundsOf(s).H()/4
			fmt.Fprint(env.txt, s)
			//log.Println("Draw Text Bounds", env.txt.Bounds())
		}
	}

	win.Clear(config.BgColor)
	env.imd.Draw(win)
	env.head.Draw(win, pixel.IM)
	env.txt.Draw(win, pixel.IM)
	env.status.Draw(win, pixel.IM)
	win.Update()

	env.frames++
	env.changed = false

}

func (env *Environment) Input(win *pixelgl.Window, config goivy.Config) {

	if win.JustReleased(pixelgl.MouseButtonLeft) {
		coord := env.cam.Unproject(win.MousePosition())
		//log.Println("MouseButtonLeft JustReleased", coord)

		step, err := env.GetStepFromCoord(coord)
		if err != nil {
			return
		}

		env.changed = true
		if env.play {
			env.steps = append(env.steps, step)
			env.calc = true
		}
	}

	if win.JustReleased(pixelgl.KeyPageDown) {
		env.changed = true
		if env.play {
			env.calc = true
		}
	}

	if win.JustReleased(pixelgl.KeyHome) {
		env.statusText = ""
		env.steps = []Step{{7, 7}}
		env.changed = true
		env.play = true
		env.calc = true
		env.auto = false
	}

	if win.JustReleased(pixelgl.KeyEnd) {
		env.changed = true
		if env.play {
			env.auto = true
		}
		env.calc = true

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
	env.play = true
	env.calc = true
	env.auto = false

	for !win.Closed() {
		env.dt = time.Since(env.last).Seconds()
		env.last = time.Now()

		env.cam = pixel.IM.Scaled(env.camPos, env.camZoom).Moved(win.Bounds().Center().Sub(env.camPos))
		win.SetMatrix(env.cam)
		win.SetSmooth(true)

		stat := ""
		if env.play && env.calc {
			env.steps, stat = calcStep(env.steps)
			env.Draw(win, config)
			if stat != "play" {
				env.play = false
				if stat == "win" {
					stat = fmt.Sprintf("%s won!", env.colorNames[1-len(env.steps)%2])
				} else if stat == "draw" {
					stat = "Draw!"
				}
				log.Println("Game over:", stat, len(env.steps))
				env.statusText = stat
			} else {
				log.Println("Game go on:", stat, len(env.steps))
			}
			if !env.auto {
				env.calc = false
			}
			env.changed = true
		}

		env.Input(win, config)

		env.Draw(win, config)

		select {
		case <-env.second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", config.WindowConfig.Title, env.frames))
			env.frames = 0
		default:
		}
	}
}

func (env *Environment) GetStepFromCoord(coord pixel.Vec) (step Step, err error) {
	x := -1
	y := -1
	for i, cx := range env.ch {
		if coord.X > cx-env.d/2 && coord.X < cx+env.d/2 {
			x = i
			break
		}
	}
	for i, cy := range env.cv {
		if coord.Y > cy-env.d/2 && coord.Y < cy+env.d/2 {
			y = i
			break
		}
	}
	if x < 0 || y < 0 {
		err = fmt.Errorf("incorrect step")
		return
	}
	step = Step{x, y}
	return
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

	rand.Seed(time.Now().Unix())

	env := &Environment{
		imd: imdraw.New(nil),
		//
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
		steps:      []Step{{7, 7}},
		colors:     [2]color.Color{colornames.Black, colornames.White},
		colorNames: [2]string{"Black", "White"},
	}

	env.fillCoord()

	env.fontFace = loadFont(18)
	env.atlas = text.NewAtlas(env.fontFace, text.ASCII)
	env.txt = text.New(pixel.ZV, env.atlas)
	env.head = text.New(pixel.ZV, env.atlas)
	env.headText = "Human Step - MouseLeftButton, Computer Step - PageDown, Auto Play - End, New Game - Home"
	env.status = text.New(pixel.ZV, env.atlas)

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

func loadFont(size float64) font.Face {

	font, err := truetype.Parse(fivo)
	if err != nil {
		panic(err)
	}

	face := truetype.NewFace(font, &truetype.Options{
		Size: size,
		//DPI:               0,
		//Hinting:           0,
		GlyphCacheEntries: 1,
		//SubPixelsX:        0,
		//SubPixelsY:        0,
	})

	return face
}
