// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"errors"
	"image"
	"time"
)

//Captcha is the core captcha struct
type Captcha struct {
	store       *CStore
	wordManager *WordManager

	captchaConfig *CaptchaConfig
	imageConfig   *ImageConfig
	filterConfig  *FilterConfig
}

//CreateCaptcha is a method to create new Captcha struct
func CreateCaptcha(wordManager *WordManager, captchaConfig *CaptchaConfig, imageConfig *ImageConfig, filterConfig *FilterConfig) *Captcha {
	captcha := new(Captcha)
	captcha.store = CreateCStore(captchaConfig.CaptchaLifeTime, captchaConfig.GcProbability, captchaConfig.GcDivisor)
	captcha.wordManager = wordManager
	captcha.captchaConfig = captchaConfig
	captcha.imageConfig = imageConfig
	captcha.filterConfig = filterConfig

	return captcha
}

//GetKey will generate a key with required length
func (captcha *Captcha) GetKey(length int) string {
	text := captcha.wordManager.Get(length)
	info := new(CaptchaInfo)
	info.Text = text
	info.CreateTime = time.Now()

	rst := captcha.store.Add(info)
	return rst
}

//Verify will verify the user's input and the server stored captcha text
func (captcha *Captcha) Verify(key, textToVerify string) (bool, string) {
	info := captcha.store.Get(key)
	if nil == info {
		return false, "captcha info not found"
	}

	if info.CreateTime.Add(captcha.captchaConfig.CaptchaLifeTime).Before(time.Now()) {
		return false, "captcha expires"
	}

	if info.Text != textToVerify {
		return false, "captcha text not match"
	}
	captcha.store.Del(key)
	return true, ""

	return false, "not reachable"
}

//GetImage will generate the binary image data
func (captcha *Captcha) GetImage(key string) (image.Image, error) {

	info := captcha.store.Get(key)
	if nil == info {
		return nil, errors.New("captcha info not found")
	}

	if info.CreateTime.Add(captcha.captchaConfig.CaptchaLifeTime).Before(time.Now()) {
		return nil, errors.New("captcha expires")
	}

	cimg := captcha.genImage(info.Text)
	return cimg, nil

}

func (captcha *Captcha) genImage(text string) *CImage {

	cimg := CreateCImage(captcha.imageConfig)
	cimg.drawString(text)

	filtermanager := CreateImageFilterManager()
	filtermanager.AddFilter(&ImageFilterStrike{StrikeLineNum: captcha.filterConfig.StrikeLineNum})
	filtermanager.AddFilter(&ImageFilterNoisePoint{NoisePointNum: captcha.filterConfig.NoisePointNum})
	filtermanager.AddFilter(&ImageFilterNoiseLine{NoiseLineNum: captcha.filterConfig.NoiseLineNum})

	for _, filter := range filtermanager.GetFilters() {
		filter.Proc(cimg)
	}

	return cimg
}
