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
