// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import ()

var imageFilterCreators = map[string]func(FilterConfigGroup) ImageFilter{}

func RegisterImageFilter(id string, f func(FilterConfigGroup) ImageFilter) bool {
	if _, has := imageFilterCreators[id]; has {
		return false
	}
	imageFilterCreators[id] = f
	return true
}

//ImageFilter is the interface of image filter
type ImageFilter interface {
	Proc(cimage *CImage)
	GetId() string
	SetConfig(FilterConfigGroup)
	GetConfig() FilterConfigGroup
}

//ImageFilter is the base class of image filter
type ImageFilterBase struct {
	config FilterConfigGroup
}

func (filter *ImageFilterBase) Proc(cimage *CImage) {
	panic("not impl")
}

func (filter *ImageFilterBase) GetId() string {
	panic("not impl")
}

func (filter *ImageFilterBase) SetConfig(config FilterConfigGroup) {
	filter.config = config
}

func (filter *ImageFilterBase) GetConfig() FilterConfigGroup {
	return filter.config
}
