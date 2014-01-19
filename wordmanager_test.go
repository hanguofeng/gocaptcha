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
	mgr := new(WordManager)
	for _, f := range singleCharDict {
		mgr.LoadFromFile(path + f)
		s := mgr.Get(length)
		if length != len([]rune(s)) {
			t.Errorf("get no equals length:" + f)
		}
	}
	for _, f := range phrasesDict {
		mgr.LoadFromFile(path + f)
		mgr.Get(length)
	}

}
