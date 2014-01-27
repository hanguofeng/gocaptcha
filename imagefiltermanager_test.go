// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"testing"
)

func TestImageFilterManager(t *testing.T) {

	manager := CreateImageFilterManager()
	manager.AddFilter(&ImageFilterBase{})
	if 1 != len(manager.GetFilters()) {
		t.Error("add or get failed")
	}
}

func TestImageFilterManagerConfig(t *testing.T) {

	filterConfig := new(FilterConfig)
	filterConfig.Init()
	filterConfig.Filters = []string{"ImageFilterNoiseLine", "ImageFilterNoisePoint", "ImageFilterStrike"}
	var filterConfigGroup *FilterConfigGroup
	filterConfigGroup = new(FilterConfigGroup)
	filterConfigGroup.Init()
	filterConfigGroup.SetItem("Num", "5")
	filterConfig.SetGroup("ImageFilterNoiseLine", filterConfigGroup)
	filterConfigGroup = new(FilterConfigGroup)
	filterConfigGroup.Init()
	filterConfigGroup.SetItem("Num", "10")
	filterConfig.SetGroup("ImageFilterNoisePoint", filterConfigGroup)
	filterConfigGroup = new(FilterConfigGroup)
	filterConfigGroup.Init()
	filterConfigGroup.SetItem("Num", "3")
	filterConfig.SetGroup("ImageFilterStrike", filterConfigGroup)

	manager := CreateImageFilterManagerByConfig(filterConfig)
	filters := manager.GetFilters()

	if 3 != len(filters) {
		t.Error("create by config failed")
	}

	for _, filter := range filters {
		config := filter.GetConfig()
		var num string
		switch filter.GetId() {
		case "ImageFilterNoiseLine":

			if num, _ = config.GetItem("Num"); "5" != num {
				t.Error("ImageFilterNoiseLine ERROR")
			}
			break
		case "ImageFilterNoisePoint":
			if num, _ = config.GetItem("Num"); "10" != num {
				t.Error("ImageFilterNoisePoint ERROR")
			}
			break
		case "ImageFilterStrike":
			if num, _ = config.GetItem("Num"); "3" != num {
				t.Error("ImageFilterStrike ERROR")
			}
			break
		}

	}

}
