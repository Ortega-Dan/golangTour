package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type MyImage struct{}

// ColorModel returns the Image's color model.
func (img MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (img MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 754, 333)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (img MyImage) At(x, y int) color.Color {

	// answer := uint8((x + y) / 2)
	// answer := uint8(x * y)
	answer := uint8(x ^ y)

	return color.RGBA{25, answer, 255, 255}
}

func main() {
	m := MyImage{}
	// pic.ShowImage(m)

	out, _ := os.Create("outie.png")
	png.Encode(out, m)
}
