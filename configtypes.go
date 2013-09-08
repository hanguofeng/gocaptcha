// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"time"
)

//CaptchaConfig ,the captcha config
type CaptchaConfig struct {
	CaptchaLifeTime time.Duration
	GcProbability   int
	GcDivisor       int
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
