// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import ()

//ImageFilter is the interface of image filter
type ImageFilter interface {
	Proc(cimage *CImage)
}

//ImageFilter is the base class of image filter
type ImageFilterBase struct {
}

func (filter *ImageFilterBase) Proc(cimage *CImage) {
	panic("not impl")
}
