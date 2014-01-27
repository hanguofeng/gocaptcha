// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"os"
	"testing"
	"time"
)

func TestStore(t *testing.T) {
	store := CreateCStore(10*time.Second, 1, 100) //10 s

	captcha := new(CaptchaInfo)
	captcha.Text = "hello"
	captcha.CreateTime = time.Now()

	//test add and get
	key := store.Add(captcha)
	retV := store.Get(key)
	if retV != captcha {
		t.Errorf("not equal")
	}

	//test dump,destroy and loaddumped
	store.Dump("data/data.dat")
	store.Destroy()
	retV = store.Get(key)
	if nil != retV {
		t.Errorf("Destroy error")
	}
	store.LoadDumped("data/data.dat")
	retV = store.Get(key)
	if captcha.Text != retV.Text {
		t.Errorf("LoadDumped error")
	}

	os.Remove("data/data.dat")

	//test del
	store.Del(key)
	retV = store.Get(key)
	if nil != retV {
		t.Errorf("not del")
	}

}
