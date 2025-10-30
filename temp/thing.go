package temp

import (
	"fmt"
	"image"
	"image/color"
)

/*
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
*/

type Image struct {
	width, height int
	colors [][][]uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{ image.Point{0,0}, image.Point{3, 4} }
}

func (i Image) At(x, y int) color.Color {
	colors := i.colors
	return color.RGBA{colors[y][x][0], colors[y][x][1], colors[y][x][2], colors[y][x][3]}
}

func main() {
	img := Image{0, 0, make([][][]uint8, 0)}
	fmt.Printf("%v", img)	
}
