// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"time"
)

//CaptchaConfig ,the captcha config
type CaptchaConfig struct {
	LifeTime time.Duration
}

//FilterConfig ,the filter config
type FilterConfig struct {
	EnableStrike     bool
	EnableNoisePoint bool
	EnableNoiseLine  bool
	StrikeLineNum    int
	NoisePointNum    int
	NoiseLineNum     int
}

//ImageConfig ,the image config
type ImageConfig struct {
	Width       int
	Height      int
	FontSize    float64
	FontFiles   []string
	fontManager *FontManager
}

const STORE_ENGINE_BUILDIN = "buildin"
const STORE_ENGINE_MEMCACHE = "memcache"

//StoreConfig ,the store engine config
type StoreConfig struct {
	CaptchaConfig
	Engine        string
	Servers       []string
	GcProbability int
	GcDivisor     int
}
