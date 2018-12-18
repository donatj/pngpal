package pngpal

import (
	"fmt"
	"image"
	"image/color"
)

// ErrTooManyColors is returned when there is too many colors for a palleted
// image. E.g. Greater than 256
var ErrTooManyColors = fmt.Errorf("too many colors for palette")

// ImageToPaletted losslessly converts an image.Image to an image.Paletted
// if possible.
//
// If the passed image.Image already is an image.Paletted it will be returned
// as is without modification.
//
// If the passed image.Image contains more than 256 colors a ErrTooManyColors
// will be returned
func ImageToPaletted(img image.Image) (*image.Paletted, error) {
	if pal, ok := img.(*image.Paletted); ok {
		return pal, nil
	}

	p, err := makePalette(img)
	if err != nil {
		return nil, err
	}

	img2 := image.NewPaletted(img.Bounds(), p)
	draw(img2, img)

	return img2, nil
}

// draw is a temporary fix as draw.Draw has a bug when copying non-paletted
// images on top of paletted images
func draw(dst *image.Paletted, src image.Image) {
	sb := src.Bounds()

	for y := sb.Min.Y; y < sb.Max.Y; y++ {
		for x := sb.Min.X; x < sb.Max.X; x++ {
			dst.Set(x, y, src.At(x, y))
		}
	}
}

func makePalette(i image.Image) (color.Palette, error) {
	sb := i.Bounds()
	cmap := make(map[[4]uint32]color.Color)

	for y := sb.Min.Y; y < sb.Max.Y; y++ {
		for x := sb.Min.X; x < sb.Max.X; x++ {
			c := i.At(x, y)
			y := [4]uint32{}

			y[0], y[1], y[2], y[3] = c.RGBA()

			cmap[y] = c
			if len(cmap) > 256 {
				return color.Palette{}, ErrTooManyColors
			}
		}
	}

	pal := color.Palette{}

	for _, c := range cmap {
		pal = append(pal, c)
	}

	return pal, nil
}
