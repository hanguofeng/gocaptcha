// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import ()

var storeCreators = map[string]func(*StoreConfig) (StoreInterface, error){}

//StoreInterface is the interface of store
type StoreInterface interface {
	Get(key string) *CaptchaInfo
	Add(captcha *CaptchaInfo) string
	Update(key string, captcha *CaptchaInfo) bool
	Del(key string)
	Destroy()
	OnConstruct()
	OnDestruct()
}

func RegisterStore(name string, f func(*StoreConfig) (StoreInterface, error)) bool {
	if _, has := storeCreators[name]; has {
		return false
	}
	storeCreators[name] = f
	return true
}
