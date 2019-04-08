package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Width, Height int
	Red, Green    uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.Red, i.Green, 255, 255}
}

func main() {
	m := Image{Width: 100, Height: 100, Red: 78, Green: 240}
	pic.ShowImage(m)
}
