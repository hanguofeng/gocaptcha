// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"errors"
	"image"
	"time"
)

type Captcha struct {
	store       *CStore
	wordManager *WordManager

	captchaConfig *CaptchaConfig
	imageConfig   *ImageConfig
	filterConfig  *FilterConfig
}

func CreateCaptcha(wordManager *WordManager, captchaConfig *CaptchaConfig, imageConfig *ImageConfig, filterConfig *FilterConfig) *Captcha {
	captcha := new(Captcha)
	captcha.store = CreateCStore(captchaConfig.CaptchaLifeTime, captchaConfig.GcProbability, captchaConfig.GcDivisor)
	captcha.wordManager = wordManager
	captcha.captchaConfig = captchaConfig
	captcha.imageConfig = imageConfig
	captcha.filterConfig = filterConfig

	return captcha
}

func (captcha *Captcha) GetKey(length int) string {
	text := captcha.wordManager.Get(length)
	info := new(CaptchaInfo)
	info.text = text
	info.createTime = time.Now()

	rst := captcha.store.Add(info)
	return rst
}

func (captcha *Captcha) Verify(key, textToVerify string) (bool, string) {
	info := captcha.store.Get(key)
	if nil == info {
		return false, "captcha info not found"
	}

	if info.createTime.Add(captcha.captchaConfig.CaptchaLifeTime).Before(time.Now()) {
		return false, "captcha expires"
	}

	if info.text != textToVerify {
		return false, "captcha text not match"
	} else {
		captcha.store.Del(key)
		return true, ""
	}

	return false, "not reachable"
}

func (captcha *Captcha) GetImage(key string) (image.Image, error) {

	info := captcha.store.Get(key)
	if nil == info {
		return nil, errors.New("captcha info not found")
	}

	if info.createTime.Add(captcha.captchaConfig.CaptchaLifeTime).Before(time.Now()) {
		return nil, errors.New("captcha expires")
	}

	cimg := captcha.genImage(info.text)
	return cimg, nil

}

func (captcha *Captcha) genImage(text string) *CImage {

	cimg := CreateCImage(captcha.imageConfig)
	cimg.drawString(text)

	if captcha.filterConfig.EnableStrike {
		cimg.strikeThrough(captcha.filterConfig.StrikeLineNum)
	}

	if captcha.filterConfig.EnableNoisePoint {
		for i := 0; i < captcha.filterConfig.NoisePointNum; i++ {
			cimg.drawCircle(rnd(0, captcha.imageConfig.Width), rnd(0, captcha.imageConfig.Height), rnd(0, 2), uint8(rnd(1, colorCount)))
		}
	}

	if captcha.filterConfig.EnableNoiseLine {
		for i := 0; i < captcha.filterConfig.NoiseLineNum; i++ {
			x := rnd(0, captcha.imageConfig.Width)
			cimg.drawHorizLine(int(float32(x)/1.5), x, rnd(0, captcha.imageConfig.Height), uint8(rnd(1, colorCount)))
		}
	}

	return cimg
}
