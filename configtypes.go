// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"time"
)

//CaptchaConfig ,the captcha config
type CaptchaConfig struct {
	LifeTime            time.Duration
	CaseSensitive       bool
	ChangeTextOnRefresh bool
}

//FilterConfigGroup
type FilterConfigGroup struct {
	data map[string]string
}

func (this *FilterConfigGroup) Init() {
	this.data = map[string]string{}
}
func (this *FilterConfigGroup) GetItem(key string) (string, bool) {
	val, ok := this.data[key]
	return val, ok
}
func (this *FilterConfigGroup) SetItem(key string, val string) {
	this.data[key] = val
}

type FilterConfig struct {
	Filters []string
	data    map[string]*FilterConfigGroup
}

func (this *FilterConfig) Init() {
	this.Filters = []string{}
	this.data = map[string]*FilterConfigGroup{}
}
func (this *FilterConfig) GetGroup(key string) (FilterConfigGroup, bool) {
	val, ok := this.data[key]
	if !ok {
		val = new(FilterConfigGroup)
	}

	return *val, ok
}
func (this *FilterConfig) SetGroup(key string, group *FilterConfigGroup) {
	this.data[key] = group
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
	Password      string
	NetWork       string
	DB            int
}
