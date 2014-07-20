// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/hanguofeng/freetype-go-mirror/freetype"
)

const (
	colorCount = 20
)

//CImage is a image process tool
type CImage struct {
	*image.Paletted
	config *ImageConfig
}

func (m *CImage) drawHorizLine(fromX, toX, y int, colorIdx uint8) *CImage {
	if 0 >= colorIdx || colorIdx > colorCount {
		colorIdx = uint8(rnd(0, colorCount))
	}
	for x := fromX; x <= toX; x++ {
		m.SetColorIndex(x, y, colorIdx)
	}

	return m
}

func (m *CImage) drawCircle(x, y, radius int, colorIdx uint8) {
	f := 1 - radius
	dfx := 1
	dfy := -2 * radius
	xo := 0
	yo := radius

	m.SetColorIndex(x, y+radius, colorIdx)
	m.SetColorIndex(x, y-radius, colorIdx)
	m.drawHorizLine(x-radius, x+radius, y, colorIdx)

	for xo < yo {
		if f >= 0 {
			yo--
			dfy += 2
			f += dfy
		}
		xo++
		dfx += 2
		f += dfx
		m.drawHorizLine(x-xo, x+xo, y+yo, colorIdx)
		m.drawHorizLine(x-xo, x+xo, y-yo, colorIdx)
		m.drawHorizLine(x-yo, x+yo, y+xo, colorIdx)
		m.drawHorizLine(x-yo, x+yo, y-xo, colorIdx)
	}
}

func (m *CImage) drawString(text string) *CImage {

	fg, bg := image.Black, &image.Uniform{color.RGBA{255, 255, 255, 255}}
	draw.Draw(m, m.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetFontSize(m.config.FontSize)
	c.SetClip(m.Bounds())
	c.SetDst(m)
	c.SetSrc(fg)
	i := 0
	for _, s := range text {
		c.SetFont(m.config.fontManager.GetRandomFont())
		charX := (int(c.PointToFix32(m.config.FontSize) >> 8)) * i
		charY := int(c.PointToFix32(m.config.FontSize) >> 8)
		charPt := freetype.Pt(charX, charY)
		c.DrawString(string(s), charPt)
		i = i + 1
	}

	return m
}

//CreateCImage will create a CImage struct with config
func CreateCImage(config *ImageConfig) *CImage {
	r := new(CImage)
	r.Paletted = image.NewPaletted(image.Rect(0, 0, config.Width, config.Height), randomPalette())
	r.config = config
	if nil == r.config.fontManager {
		fm := CreateFontManager()
		for _, fontFile := range config.FontFiles {
			fm.AddFont(fontFile)
		}
		r.config.fontManager = fm
	}

	return r
}

func randomPalette() color.Palette {

	p := make([]color.Color, colorCount+1)
	// Transparent color.
	p[0] = color.RGBA{0xFF, 0xFF, 0xFF, 0x00}
	// Primary color.
	prim := color.RGBA{
		uint8(rnd(0, 255)),
		uint8(rnd(0, 255)),
		uint8(rnd(0, 255)),
		0xFF,
	}
	p[1] = prim
	// Circle colors.
	for i := 2; i <= colorCount; i++ {
		p[i] = randomBrightness(prim, 255)
	}

	return p
}

func randomBrightness(c color.RGBA, max uint8) color.RGBA {
	minc := min3(c.R, c.G, c.B)
	maxc := max3(c.R, c.G, c.B)
	if maxc > max {
		return c
	}
	n := rnd(0, int(max-maxc)) - int(minc)
	return color.RGBA{
		uint8(int(c.R) + n),
		uint8(int(c.G) + n),
		uint8(int(c.B) + n),
		uint8(c.A),
	}
}

func min3(x, y, z uint8) (m uint8) {
	m = x
	if y < m {
		m = y
	}
	if z < m {
		m = z
	}
	return
}

func max3(x, y, z uint8) (m uint8) {
	m = x
	if y > m {
		m = y
	}
	if z > m {
		m = z
	}
	return
}
