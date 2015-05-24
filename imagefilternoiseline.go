// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"strconv"
)

//ImageFilter is the interface of image filter
const IMAGE_FILTER_NOISE_LINE = "ImageFilterNoiseLine"

type ImageFilterNoiseLine struct {
	ImageFilterBase
}

func imageFilterNoiseLineCreator(config FilterConfigGroup) ImageFilter {
	filter := ImageFilterNoiseLine{}
	filter.SetConfig(config)
	return &filter
}

//Proc the image
func (filter *ImageFilterNoiseLine) Proc(cimage *CImage) {

	var num int
	var err error
	v, ok := filter.config.GetItem("Num")
	if ok {
		num, err = strconv.Atoi(v)
		if nil != err {
			num = 3
		}
	} else {
		num = 3
	}

	for i := 0; i < num; i++ {
		x := rnd(0, cimage.Bounds().Max.X)
		cimage.drawHorizLine(int(float32(x)/1.5), x, rnd(0, cimage.Bounds().Max.Y), uint8(rnd(1, colorCount)))
	}
}

func (filter *ImageFilterNoiseLine) GetId() string {
	return IMAGE_FILTER_NOISE_LINE
}
