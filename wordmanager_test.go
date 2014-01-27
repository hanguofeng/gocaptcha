// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"os"
	"testing"
)

func TestWordManager(t *testing.T) {
	pwd, err := os.Getwd()
	if (nil != err) || "" == pwd {
		return
	}
	path := pwd + "/data/"
	singleCharDict := []string{"cn_char", "en_char"}
	phrasesDict := []string{"cn_phrases", "en_phrases"}
	length := 6

	for _, f := range singleCharDict {
		mgr, err := CreateWordManagerFromDataFile(path + f)
		s, err := mgr.Get(length)
		if nil != err {
			t.Errorf(err.Error())
		} else if length != len([]rune(s)) {
			t.Errorf("get no equals length:" + f)
		}
	}
	for _, f := range phrasesDict {
		mgr, err := CreateWordManagerFromDataFile(path + f)
		s, err := mgr.Get(length)
		if nil != err {
			t.Errorf(err.Error())
		} else if 0 >= len([]rune(s)) {
			t.Errorf("not get:" + f)
		}
	}

}
