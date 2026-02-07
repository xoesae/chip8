package display

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/xoesae/chip8/emulator/event"
	"github.com/xoesae/chip8/logger"
)

const (
	DisplayWidth  = 64
	DisplayHeight = 32
	PixelSize     = 10
)

type Display struct {
	app    *fyne.App
	window *fyne.Window
	pixels [][]*canvas.Rectangle
	grid   *fyne.Container
	events chan event.Event
}

func NewDisplay(a *fyne.App, eventsChannel chan event.Event) *Display {
	window := (*a).NewWindow("CHIP-8")
	window.Resize(fyne.NewSize(DisplayWidth*PixelSize, DisplayHeight*PixelSize))

	pixels := make([][]*canvas.Rectangle, DisplayHeight)
	grid := container.NewGridWithColumns(DisplayWidth)

	for y := 0; y < DisplayHeight; y++ {
		pixels[y] = make([]*canvas.Rectangle, DisplayWidth)

		for x := 0; x < DisplayWidth; x++ {
			rect := canvas.NewRectangle(color.Black)
			rect.SetMinSize(fyne.NewSize(PixelSize, PixelSize))
			pixels[y][x] = rect
			grid.Add(rect)
		}
	}

	d := &Display{
		app:    a,
		window: &window,
		pixels: pixels,
		grid:   grid,
		events: eventsChannel,
	}

	window.SetContent(grid)

	return d
}

func (d *Display) Events() chan<- event.Event {
	return d.events
}

func (d *Display) StartEventLoop() {
	go d.eventLoop()
}

func (d *Display) eventLoop() {
	for ev := range d.events {
		switch ev.(type) {
		case event.DisplayClearEvent:
			d.clearDisplay()
		case event.DisplayUpdatedEvent:
			evt := ev.(event.DisplayUpdatedEvent)
			d.updatePixel(evt)
		default:
			logger.Get().Debug("Event ignored", ev)
		}
	}
}

func (d *Display) clearDisplay() {
	fyne.Do(func() {
		for i := 0; i < DisplayHeight; i++ {
			for j := 0; j < DisplayWidth; j++ {
				d.pixels[i][j].FillColor = color.Black
				//d.pixels[i][j].Refresh()
			}
		}

		d.grid.Refresh()
	})

	logger.Get().Debug("Display clear")
}

func (d *Display) updatePixel(evt event.DisplayUpdatedEvent) {
	fyne.Do(func() {
		for i := 0; i < DisplayHeight; i++ {
			for j := 0; j < DisplayWidth; j++ {
				_color := color.Black
				if evt.Pixels[i][j] {
					_color = color.White
				}

				d.pixels[i][j].FillColor = _color
				//d.pixels[i][j].Refresh()
			}
		}

		d.grid.Refresh()
	})

	logger.Get().Debug("Display updated")
}

func (d *Display) Show() {
	(*d.window).Show()
}

func (d *Display) Run() {
	(*d.window).ShowAndRun()
}

func (d *Display) Close() {
	close(d.events)
	(*d.window).Close()
}
