// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"testing"
	"time"
)

func TestRedisStore(t *testing.T) {
	storeConfig := &StoreConfig{}
	storeConfig.CaptchaConfig.LifeTime = time.Second * 100
	storeConfig.Engine = "redis"
	storeConfig.Servers = []string{"127.0.0.1:6379"}
	storeConfig.Password = "123456"
	storeConfig.DB = 5
	storeConfig.NetWork = "tcp"
	store, _ := CreateCaptchaRedisStore(storeConfig)

	captcha := new(CaptchaInfo)
	captcha.Text = "hello"
	captcha.CreateTime = time.Now()

	//test add and get
	key := store.Add(captcha)
	retV := store.Get(key)

	if retV.Text != captcha.Text {
		t.Errorf("not equal,retV:%s,captcha:%s", retV, captcha)
	}

	retV.Text = "world"
	store.Update(key, retV)
	retV = store.Get(key)
	if retV.Text != "world" {
		t.Errorf("update not equal,retV:%s,captcha:%s", retV, captcha)
	}

	t.Logf("TestMCStore:get from redis:%s", retV)

	//test del
	store.Del(key)
	retV = store.Get(key)
	if nil != retV {
		t.Errorf("not del")
	}

}
