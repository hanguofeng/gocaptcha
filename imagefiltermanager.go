// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import ()

//ImageFilterManager
type ImageFilterManager struct {
	filters []ImageFilter
}

func CreateImageFilterManager() *ImageFilterManager {
	ret := new(ImageFilterManager)
	ret.filters = []ImageFilter{}
	return ret
}

func (manager *ImageFilterManager) AddFilter(filter ImageFilter) {
	manager.filters = append(manager.filters, filter)
}

func (manager *ImageFilterManager) GetFilters() []ImageFilter {
	return manager.filters
}
