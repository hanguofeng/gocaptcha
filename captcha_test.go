// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"image/png"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestCaptcha(t *testing.T) {

	var err error

	if "windows" != runtime.GOOS {
		return
	}

	pwd, err := os.Getwd()
	if (nil != err) || "" == pwd {
		return
	}

	path := pwd + "/data/cn_phrases"
	wordmgr := new(WordManager)
	wordmgr.LoadFromFile(path)

	captchaConfig, imageConfig, filterConfig := loadConfig()

	captcha := CreateCaptcha(wordmgr, captchaConfig, imageConfig, filterConfig)
	key := captcha.GetKey(4)
	img, err := captcha.GetImage(key)
	captcha.Verify(key, "用户输入")

	err = os.Mkdir(pwd+"/_test_files/", os.ModeDir)
	f, err := os.Create(pwd + "/_test_files/hanguofeng.png")
	defer f.Close()
	png.Encode(f, img)
	f.Close()

}

func BenchmarkCaptcha(t *testing.B) {
	var err error

	if "windows" != runtime.GOOS {
		return
	}

	pwd, err := os.Getwd()
	if (nil != err) || "" == pwd {
		return
	}

	path := pwd + "/data/cn_phrases"
	wordmgr := new(WordManager)
	wordmgr.LoadFromFile(path)

	captchaConfig, imageConfig, filterConfig := loadConfig()

	for i := 0; i < t.N; i++ {
		ccaptcha := CreateCaptcha(wordmgr, captchaConfig, imageConfig, filterConfig)
		s := ccaptcha.GetKey(4)
		ccaptcha.GetImage(s)
		ccaptcha.Verify(s, "ssss")
	}
}

func BenchmarkCaptchaInternalAPI(t *testing.B) {
	var err error

	if "windows" != runtime.GOOS {
		return
	}

	pwd, err := os.Getwd()
	if (nil != err) || "" == pwd {
		return
	}

	path := pwd + "/data/cn_phrases"
	wordmgr := new(WordManager)
	wordmgr.LoadFromFile(path)

	captchaConfig, imageConfig, filterConfig := loadConfig()

	for i := 0; i < t.N; i++ {
		ccaptcha := CreateCaptcha(wordmgr, captchaConfig, imageConfig, filterConfig)
		s := ccaptcha.GetKey(4)
		ccaptcha.Verify(s, "ssss")
	}
}

func BenchmarkCaptchaDrawImage(t *testing.B) {
	var err error

	if "windows" != runtime.GOOS {
		return
	}

	pwd, err := os.Getwd()
	if (nil != err) || "" == pwd {
		return
	}

	path := pwd + "/data/cn_phrases"
	wordmgr := new(WordManager)
	wordmgr.LoadFromFile(path)

	captchaConfig, imageConfig, filterConfig := loadConfig()

	for i := 0; i < t.N; i++ {
		ccaptcha := CreateCaptcha(wordmgr, captchaConfig, imageConfig, filterConfig)
		s := ccaptcha.GetKey(4)
		ccaptcha.GetImage(s)
	}
}

func loadConfig() (*CaptchaConfig, *ImageConfig, *FilterConfig) {

	captchaConfig := new(CaptchaConfig)
	captchaConfig.CaptchaLifeTime = 10 * time.Second
	captchaConfig.GcProbability = 1
	captchaConfig.GcDivisor = 100

	imageConfig := new(ImageConfig)
	imageConfig.FontFiles = []string{
		"c:/windows/fonts/SIMLI.TTF",
		"c:/windows/fonts/simfang.ttf",
		"c:/windows/fonts/SIMYOU.TTF",
		"c:/windows/fonts/msyh.TTF",
		"c:/windows/fonts/simhei.ttf",
		"c:/windows/fonts/simkai.ttf"}
	imageConfig.FontSize = 26
	imageConfig.Height = 40
	imageConfig.Width = 120

	filterConfig := new(FilterConfig)
	filterConfig.EnableNoiseLine = true
	filterConfig.EnableNoisePoint = true
	filterConfig.EnableStrike = true
	filterConfig.StrikeLineNum = 3
	filterConfig.NoisePointNum = 30
	filterConfig.NoiseLineNum = 10

	return captchaConfig, imageConfig, filterConfig
}
