// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"time"
)

type CaptchaInfo struct {
	text       string
	createTime time.Time
}
