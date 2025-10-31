package display

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

const (
	width  = 64
	height = 32
	scale  = 10
)

type Display struct {
	pixels      [height][width]bool
	Window      fyne.Window
	raster      *canvas.Raster
	pixelCanvas [height][width]*canvas.Rectangle
}

func NewDisplay() *Display {
	a := app.New()
	w := a.NewWindow("Chip-8 Emulator")
	w.Resize(fyne.NewSize(width*scale, height*scale))

	d := &Display{
		Window: w,
	}

	d.raster = canvas.NewRaster(func(w, h int) image.Image {
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if d.pixels[y][x] {
					img.Set(x, y, color.White)
				} else {
					img.Set(x, y, color.Black)
				}
			}
		}
		return img
	})
	d.raster.ScaleMode = canvas.ImageScalePixels

	w.SetContent(d.raster)
	return d
}

func (d *Display) Refresh() {
	d.raster.Refresh()
}

func (d *Display) Clear() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			d.pixels[y][x] = false
		}
	}
	d.Refresh()
}

func (d *Display) DrawSprite(x, y int, sprite []byte) bool {
	collision := false

	for row, data := range sprite {
		for col := 0; col < 8; col++ {
			if (data & (0x80 >> col)) != 0 {
				px := (x + col) % width
				py := (y + row) % height
				if d.pixels[py][px] {
					collision = true
				}
				d.pixels[py][px] = !d.pixels[py][px]
			}
		}
	}

	d.Refresh()

	return collision
}

func (d *Display) Run() {
	d.Window.ShowAndRun()
}
