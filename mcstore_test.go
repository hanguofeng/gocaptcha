// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"testing"
	"time"
)

func TestMCStore(t *testing.T) {
	return //skip this test

	memcacheServers := []string{"127.0.0.1:11211"}
	store := CreateMCStore(10*time.Second, memcacheServers) //10 s

	captcha := new(CaptchaInfo)
	captcha.Text = "hello"
	captcha.CreateTime = time.Now()

	//test add and get
	key := store.Add(captcha)
	retV := store.Get(key)

	if retV.Text != captcha.Text {
		t.Errorf("not equal,retV:%s,captcha:%s", retV, captcha)
	}

	//test del
	store.Del(key)
	retV = store.Get(key)
	if nil != retV {
		t.Errorf("not del")
	}

}
