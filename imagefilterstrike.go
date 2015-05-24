// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"math"
	"strconv"
)

const IMAGE_FILTER_STRIKE = "ImageFilterStrike"

//ImageFilter is the interface of image filter
type ImageFilterStrike struct {
	ImageFilterBase
}

func imageFilterStrikeCreator(config FilterConfigGroup) ImageFilter {
	filter := ImageFilterStrike{}
	filter.SetConfig(config)
	return &filter
}

//Proc the image
func (filter *ImageFilterStrike) Proc(cimage *CImage) {

	var num int
	var err error
	v, ok := filter.config.GetItem("Num")
	if ok {
		num, err = strconv.Atoi(v)
		if nil != err {
			num = 1
		}
	} else {
		num = 1
	}

	maxx := cimage.Bounds().Max.X
	maxy := cimage.Bounds().Max.Y
	y := rnd(maxy/2, maxy-maxy/2)
	amplitude := rndf(10, 15)
	period := rndf(80, 100)
	dx := 2.0 * math.Pi / period
	for x := 0; x < maxx; x++ {
		xo := amplitude * math.Cos(float64(y)*dx)
		yo := amplitude * math.Sin(float64(x)*dx)
		for yn := 0; yn < num; yn++ {
			r := rnd(0, 2)
			cimage.drawCircle(x+int(xo), y+int(yo)+(yn*(num+1)), r/2, 1)
		}
	}
}

func (filter *ImageFilterStrike) GetId() string {
	return IMAGE_FILTER_STRIKE
}
