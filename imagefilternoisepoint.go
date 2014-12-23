// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"strconv"
)

const IMAGE_FILTER_NOISE_POINT = "ImageFilterNoisePoint"

//ImageFilter is the interface of image filter
type ImageFilterNoisePoint struct {
	ImageFilterBase
}

func imageFilterNoisePointCreator(config FilterConfigGroup) ImageFilter {
	filter := ImageFilterNoisePoint{}
	filter.SetConfig(config)
	return &filter
}

//Proc the image
func (filter *ImageFilterNoisePoint) Proc(cimage *CImage) {
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
		cimage.drawCircle(rnd(0, cimage.Bounds().Max.X), rnd(0, cimage.Bounds().Max.Y), rnd(0, 2), uint8(rnd(1, colorCount)))
	}
}

func (filter *ImageFilterNoisePoint) GetId() string {
	return IMAGE_FILTER_NOISE_POINT
}
