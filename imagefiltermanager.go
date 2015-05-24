// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import ()

//ImageFilterManager
type ImageFilterManager struct {
	filters []ImageFilter
}

func init() {
	RegisterImageFilter(IMAGE_FILTER_NOISE_LINE, imageFilterNoiseLineCreator)
	RegisterImageFilter(IMAGE_FILTER_NOISE_POINT, imageFilterNoisePointCreator)
	RegisterImageFilter(IMAGE_FILTER_STRIKE, imageFilterStrikeCreator)
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

func CreateImageFilterManagerByConfig(config *FilterConfig) *ImageFilterManager {
	mgr := new(ImageFilterManager)
	mgr.filters = []ImageFilter{}

	for _, cfgId := range config.Filters {
		creator, has := imageFilterCreators[cfgId]
		if !has {
			continue
		}
		cfgGroup, has := config.GetGroup(cfgId)
		if !has {
			continue
		}
		mgr.AddFilter(creator(cfgGroup))
	}

	return mgr
}
