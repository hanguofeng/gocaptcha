// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"time"
)

// CaptchaInfo is the entity of a captcha
// text:the content text,for the image display and user to recognize
// createTime:the time when the captcha is created
type CaptchaInfo struct {
	Text       string
	CreateTime time.Time
	ShownTimes int
}
