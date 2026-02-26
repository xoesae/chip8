package display

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/xoesae/chip8/emulator/shared"
)

type KeyEvent struct {
	Key     uint8
	Pressed bool
}

type Display struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func NewDisplay() (*Display, error) {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		return nil, err
	}

	window, err := sdl.CreateWindow(
		"CHIP-8",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		shared.DisplayWidth*shared.PixelSize,
		shared.DisplayHeight*shared.PixelSize,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}

	return &Display{
		window:   window,
		renderer: renderer,
	}, nil
}

func (d *Display) Render(pixels [shared.DisplayHeight][shared.DisplayWidth]bool) {
	d.renderer.SetDrawColor(0, 0, 0, 255)
	d.renderer.Clear()

	for y := range pixels {
		for x := range pixels[y] {
			if pixels[y][x] {
				rect := sdl.Rect{
					X: int32(x * shared.PixelSize),
					Y: int32(y * shared.PixelSize),
					W: shared.PixelSize,
					H: shared.PixelSize,
				}

				d.renderer.SetDrawColor(255, 255, 255, 255)
				d.renderer.FillRect(&rect)
			}
		}
	}

	d.renderer.Present()
}

func (d *Display) PollEvents() ([]KeyEvent, bool) {
	var events []KeyEvent
	running := true

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {

		case *sdl.QuitEvent:
			running = false

		case *sdl.KeyboardEvent:
			key, ok := mapKey(e.Keysym.Sym)
			if ok {
				events = append(events, KeyEvent{
					Key:     key,
					Pressed: e.Type == sdl.KEYDOWN,
				})
			}
		}
	}

	return events, running
}

func (d *Display) Close() {
	d.renderer.Destroy()
	d.window.Destroy()
	sdl.Quit()
}

func mapKey(key sdl.Keycode) (uint8, bool) {
	switch key {
	case sdl.K_1:
		return 0x1, true
	case sdl.K_2:
		return 0x2, true
	case sdl.K_3:
		return 0x3, true
	case sdl.K_4:
		return 0xC, true
	case sdl.K_q:
		return 0x4, true
	case sdl.K_w:
		return 0x5, true
	case sdl.K_e:
		return 0x6, true
	case sdl.K_r:
		return 0xD, true
	case sdl.K_a:
		return 0x7, true
	case sdl.K_s:
		return 0x8, true
	case sdl.K_d:
		return 0x9, true
	case sdl.K_f:
		return 0xE, true
	case sdl.K_z:
		return 0xA, true
	case sdl.K_x:
		return 0x0, true
	case sdl.K_c:
		return 0xB, true
	case sdl.K_v:
		return 0xF, true
	}

	return 0, false
}
