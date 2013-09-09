// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import ()

//ImageFilter is the interface of image filter
type ImageFilterNoisePoint struct {
	NoisePointNum int
}

//Proc the image
func (filter *ImageFilterNoisePoint) Proc(cimage *CImage) {
	for i := 0; i < filter.NoisePointNum; i++ {
		cimage.drawCircle(rnd(0, cimage.Bounds().Max.X), rnd(0, cimage.Bounds().Max.Y), rnd(0, 2), uint8(rnd(1, colorCount)))
	}
}
