// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"os"
	"testing"
	"time"
)

func TestCaptcha(t *testing.T) {

	captcha, err := getCaptcha()
	if nil != err {
		t.Fatalf("getCaptcha Error:%s", err.Error())
	}
	key, err := captcha.GetKey(4)
	if nil != err {
		t.Fatalf("GetKey Error:%s", err.Error())
	}
	captcha.GetImage(key)
	captcha.Verify(key, "test")

}

func BenchmarkCaptcha(t *testing.B) {
	captcha, err := getCaptcha()
	if nil != err {
		t.Fatalf("getCaptcha Error:%s", err.Error())
	}

	for i := 0; i < t.N; i++ {

		s, _ := captcha.GetKey(4)
		captcha.GetImage(s)
		captcha.Verify(s, "ssss")
	}
}

func BenchmarkCaptchaInternalAPI(t *testing.B) {
	captcha, err := getCaptcha()
	if nil != err {
		t.Fatalf("getCaptcha Error:%s", err.Error())
	}
	for i := 0; i < t.N; i++ {

		s, _ := captcha.GetKey(4)
		captcha.Verify(s, "ssss")
	}
}

func BenchmarkCaptchaDrawImage(t *testing.B) {
	captcha, err := getCaptcha()
	if nil != err {
		t.Fatalf("getCaptcha Error:%s", err.Error())
	}
	for i := 0; i < t.N; i++ {

		s, _ := captcha.GetKey(4)
		captcha.GetImage(s)
	}
}

func getCaptcha() (*Captcha, error) {
	wordDict, captchaConfig, imageConfig, filterConfig, storeConfig := loadConfig()

	wordmgr, err := CreateWordManagerFromDataFile(wordDict)
	captcha, err := CreateCaptcha(wordmgr, captchaConfig, imageConfig, filterConfig, storeConfig)
	return captcha, err
}

func loadConfig() (string, *CaptchaConfig, *ImageConfig, *FilterConfig, *StoreConfig) {

	pwd, _ := os.Getwd()
	data_path := pwd + "/data/"

	wordDict := data_path + "en_phrases"

	captchaConfig := new(CaptchaConfig)
	captchaConfig.LifeTime = 10 * time.Second

	imageConfig := new(ImageConfig)
	imageConfig.FontFiles = []string{data_path + "zpix.ttf"}
	imageConfig.FontSize = 26
	imageConfig.Height = 40
	imageConfig.Width = 120

	filterConfig := new(FilterConfig)
	filterConfig.Init()
	filterConfig.Filters = []string{"ImageFilterNoiseLine", "ImageFilterNoisePoint", "ImageFilterStrike"}

	var filterConfigGroup *FilterConfigGroup
	filterConfigGroup = new(FilterConfigGroup)
	filterConfigGroup.Init()
	filterConfigGroup.SetItem("Num", "5")
	filterConfig.SetGroup("ImageFilterNoiseLine", filterConfigGroup)
	filterConfigGroup = new(FilterConfigGroup)
	filterConfigGroup.Init()
	filterConfigGroup.SetItem("Num", "10")
	filterConfig.SetGroup("ImageFilterNoisePoint", filterConfigGroup)
	filterConfigGroup = new(FilterConfigGroup)
	filterConfigGroup.Init()
	filterConfigGroup.SetItem("Num", "3")
	filterConfig.SetGroup("ImageFilterStrike", filterConfigGroup)

	storeConfig := new(StoreConfig)
	storeConfig.Engine = STORE_ENGINE_BUILDIN
	storeConfig.GcDivisor = 100
	storeConfig.GcProbability = 1

	return wordDict, captchaConfig, imageConfig, filterConfig, storeConfig
}
