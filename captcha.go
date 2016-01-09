// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"errors"
	"image"
	"strings"
	"time"
)

//Captcha is the core captcha struct
type Captcha struct {
	store         StoreInterface
	wordManager   *WordManager
	filterManager *ImageFilterManager

	captchaConfig *CaptchaConfig
	imageConfig   *ImageConfig
	filterConfig  *FilterConfig
	storeConfig   *StoreConfig
}

//CreateCaptcha is a method to create new Captcha struct
func CreateCaptcha(wordManager *WordManager, captchaConfig *CaptchaConfig, imageConfig *ImageConfig, filterConfig *FilterConfig, storeConfig *StoreConfig) (*Captcha, error) {
	var retErr error

	captcha := new(Captcha)

	store, err := createStore(storeConfig)
	if nil == err {
		captcha.store = store
	} else {
		retErr = err
	}
	if nil != wordManager {
		captcha.wordManager = wordManager
	} else {
		retErr = errors.New("CreateCaptcha fail:invalid wordManager")
	}
	captcha.captchaConfig = captchaConfig
	captcha.imageConfig = imageConfig
	captcha.filterConfig = filterConfig

	captcha.filterManager = CreateImageFilterManagerByConfig(filterConfig)

	return captcha, retErr
}

//CreateCaptchaFromConfigFile is a method to create new Captcha struct
func CreateCaptchaFromConfigFile(configFile string) (*Captcha, error) {
	var captcha *Captcha
	var retErr error

	err, wordDict, captchaConfig, imageConfig, filterConfig, storeConfig := loadConfigFromFile(configFile)
	if nil == err {

		wordmgr, err := CreateWordManagerFromDataFile(wordDict)
		if nil == err {
			captcha, retErr = CreateCaptcha(wordmgr, captchaConfig, imageConfig, filterConfig, storeConfig)
		} else {
			retErr = err
		}
	} else {
		retErr = err
	}
	return captcha, retErr
}

//GetKey will generate a key with required length
func (captcha *Captcha) GetKey(length int) (string, error) {
	var retErr error
	var rst string
	text, err := captcha.wordManager.Get(length)
	if nil != err {
		retErr = err
	} else {
		info := new(CaptchaInfo)
		info.Text = text
		info.CreateTime = time.Now()
		info.ShownTimes = 0
		rst = captcha.store.Add(info)
	}
	return rst, retErr
}

//Verify will verify the user's input and the server stored captcha text
func (captcha *Captcha) Verify(key, textToVerify string) (bool, string) {
	info := captcha.store.Get(key)
	if nil == info {
		return false, "captcha info not found"
	}

	if info.CreateTime.Add(captcha.captchaConfig.LifeTime).Before(time.Now()) {
		return false, "captcha expires"
	}

	verified := false
	if captcha.captchaConfig.CaseSensitive {
		verified = info.Text == textToVerify
	} else {
		verified = strings.ToLower(info.Text) == strings.ToLower(textToVerify)
	}

	if !verified {
		return false, "captcha text not match"
	}

	captcha.store.Del(key)
	return true, ""
}

//GetImage will generate the binary image data
func (captcha *Captcha) GetImage(key string) (image.Image, error) {

	info := captcha.store.Get(key)
	if nil == info {
		return nil, errors.New("captcha info not found")
	}

	if info.CreateTime.Add(captcha.captchaConfig.LifeTime).Before(time.Now()) {
		return nil, errors.New("captcha expires")
	}

	if captcha.captchaConfig.ChangeTextOnRefresh {
		if info.ShownTimes > 0 {

			text, err := captcha.wordManager.Get(len(info.Text))
			if nil != err {
				return nil, err
			} else {
				info.Text = text
			}
		}

		info.ShownTimes++
		captcha.store.Update(key, info)
	}

	cimg := captcha.genImage(info.Text)
	return cimg, nil

}
func createStore(config *StoreConfig) (StoreInterface, error) {
	var store StoreInterface
	var err error
	switch config.Engine {
	case STORE_ENGINE_BUILDIN:
		store = CreateCStore(config.LifeTime, config.GcProbability, config.GcDivisor)
		break
	case STORE_ENGINE_MEMCACHE:
		store = CreateMCStore(config.LifeTime, config.Servers)
		break
	default:
		creator, has := storeCreators[config.Engine]
		if !has {
			err = errors.New("Not supported engine:'" + config.Engine + "'")
			break
		}
		store, err = creator(config)
	}
	return store, err
}

func (captcha *Captcha) genImage(text string) *CImage {

	cimg := CreateCImage(captcha.imageConfig)
	cimg.drawString(text)

	for _, filter := range captcha.filterManager.GetFilters() {
		filter.Proc(cimg)
	}

	return cimg
}
