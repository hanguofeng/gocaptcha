// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"math"
)

//ImageFilter is the interface of image filter
type ImageFilterStrike struct {
	*ImageFilterBase
	StrikeLineNum int
}

//Proc the image
func (filter *ImageFilterStrike) Proc(cimage *CImage) {

	maxx := cimage.Bounds().Max.X
	maxy := cimage.Bounds().Max.Y
	y := rnd(maxy/2, maxy-maxy/2)
	amplitude := rndf(10, 15)
	period := rndf(80, 100)
	dx := 2.0 * math.Pi / period
	for x := 0; x < maxx; x++ {
		xo := amplitude * math.Cos(float64(y)*dx)
		yo := amplitude * math.Sin(float64(x)*dx)
		for yn := 0; yn < filter.StrikeLineNum; yn++ {
			r := rnd(0, 2)
			cimage.drawCircle(x+int(xo), y+int(yo)+(yn*(filter.StrikeLineNum+1)), r/2, 1)
		}
	}
}
