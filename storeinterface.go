// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import ()

//StoreInterface is the interface of store
type StoreInterface interface {
	Get(key string) *CaptchaInfo
	Add(captcha *CaptchaInfo) string
	Del(key string)
	Destroy()
	OnConstruct()
	OnDestruct()
}
