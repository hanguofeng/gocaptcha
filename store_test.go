// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"testing"
	"time"
)

func TestStore(t *testing.T) {
	store := CreateCStore(10*time.Second, 1, 100) //10 s

	captcha := new(CaptchaInfo)
	captcha.text = "hello"
	captcha.createTime = time.Now()

	key := store.Add(captcha)
	ret_v := store.Get(key)
	if ret_v != captcha {
		t.Errorf("not equal")
	}

	store.Del(key)
	ret_v = store.Get(key)
	if nil != ret_v {
		t.Errorf("not del")
	}

}
