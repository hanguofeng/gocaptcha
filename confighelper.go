// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"errors"
	"strings"
	"time"

	"github.com/hanguofeng/config"
)

const (
	DEFAULT_LIFE_TIME      = time.Minute * 5
	DEFAULT_FONT_SIZE      = 26
	DEFAULT_WIDTH          = 120
	DEFAULT_HEIGHT         = 40
	DEFAULT_GC_PROBABILITY = 1
	DEFAULT_GC_DIVISOR     = 100
)

func loadConfigFromFile(configFile string) (error, string, *CaptchaConfig, *ImageConfig, *FilterConfig, *StoreConfig) {

	var retErr error

	c, err := config.ReadDefault(configFile)

	//wordDict
	wordDict, err := c.String("captcha", "word_dict")
	if nil != err {
		retErr = errors.New("loadConfigFromFile Fail,Get word_dict options failed:" + err.Error())
	}
	//captchaConfig
	captchaConfig := new(CaptchaConfig)
	var lifeTime time.Duration

	cfgLifeTime, err := c.String("captcha", "life_time")
	if nil == err {
		lifeTime, err = time.ParseDuration(cfgLifeTime)
		if nil != err {
			lifeTime = DEFAULT_LIFE_TIME
		}
	} else {
		lifeTime = DEFAULT_LIFE_TIME
	}

	captchaConfig.LifeTime = lifeTime

	//imageConfig
	imageConfig := new(ImageConfig)
	var fontFiles []string
	cfgFontFiles, err := c.StringMuti("image", "font_files")
	if nil != err {
		retErr = errors.New("loadConfigFromFile Fail,font_files options failed:" + err.Error())
	} else {
		fontFiles = cfgFontFiles
	}

	imageConfig.FontFiles = fontFiles

	var fontSize float64
	cfgFontSize, err := c.Int("image", "font_size")
	if nil != err {
		fontSize = DEFAULT_FONT_SIZE
	} else {
		fontSize = float64(cfgFontSize)
	}

	imageConfig.FontSize = fontSize

	var width int
	cfgWidth, err := c.Int("image", "width")
	if nil != err {
		width = DEFAULT_WIDTH
	} else {
		width = int(cfgWidth)
	}
	imageConfig.Width = width

	var height int
	cfgHeight, err := c.Int("image", "height")
	if nil != err {
		height = DEFAULT_HEIGHT
	} else {
		height = int(cfgHeight)
	}
	imageConfig.Height = height

	//filterConfig
	filterConfig := new(FilterConfig)
	filterConfig.Init()
	cfgOpenFilter, err := c.StringMuti("filter", "open_filter")
	if nil == err {
		filterConfig.Filters = cfgOpenFilter
	} else {
		filterConfig.Filters = []string{}
	}
	for _, section := range c.Sections() {
		if strings.HasPrefix(section, "ImageFilter") {
			filterConfigGroup := new(FilterConfigGroup)
			filterConfigGroup.Init()
			options, err := c.Options(section)
			if nil == err {
				for _, option := range options {
					v, err := c.String(section, option)
					if nil == err {
						filterConfigGroup.SetItem(option, v)
					}
				}
			}
			filterConfig.SetGroup(section, filterConfigGroup)
		}
	}

	//storeConfig
	storeConfig := new(StoreConfig)
	engine, err := c.String("store", "engine")
	if nil != err {
		retErr = errors.New("loadConfigFromFile Fail,engine options failed" + err.Error())
	} else {
		storeConfig.Engine = engine
	}
	servers, err := c.StringMuti("store", "servers")
	if nil != err {
		storeConfig.Servers = []string{}
	} else {
		storeConfig.Servers = servers
	}
	gcProbability, err := c.Int("store", "gc_probability")
	if nil != err {
		storeConfig.GcProbability = gcProbability
	} else {
		storeConfig.GcProbability = DEFAULT_GC_PROBABILITY
	}
	gcDivisor, err := c.Int("store", "gc_divisor")
	if nil != err {
		storeConfig.GcDivisor = gcDivisor
	} else {
		storeConfig.GcDivisor = DEFAULT_GC_DIVISOR
	}

	if nil != err && nil == retErr {
		retErr = err
	}

	return retErr, wordDict, captchaConfig, imageConfig, filterConfig, storeConfig
}
