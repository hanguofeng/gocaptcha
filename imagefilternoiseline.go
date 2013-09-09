// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import ()

//ImageFilter is the interface of image filter
type ImageFilterNoiseLine struct {
	NoiseLineNum int
}

//Proc the image
func (filter *ImageFilterNoiseLine) Proc(cimage *CImage) {
	for i := 0; i < filter.NoiseLineNum; i++ {
		x := rnd(0, cimage.Bounds().Max.X)
		cimage.drawHorizLine(int(float32(x)/1.5), x, rnd(0, cimage.Bounds().Max.Y), uint8(rnd(1, colorCount)))
	}
}
