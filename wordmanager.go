// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

//WordManager is a captcha word manage tool
type WordManager struct {
	words            []string
	isDataSingleChar bool
}

//Get a specifical length word
func (mgr *WordManager) Get(length int) string {

	rst := ""
	if true == mgr.isDataSingleChar {
		if len(mgr.words) < length {
			panic("dict words count is less than your length")
		}

		for {
			line := mgr.getLine()
			if false == strings.ContainsRune(rst, []rune(line)[0]) {
				rst = rst + line
			}
			if utf8.RuneCountInString(rst) >= length {
				break
			}
		}

		rstRune := []rune(rst)
		rst = string(rstRune[0:length])
	} else {
		rst = mgr.getLine()
	}

	return rst
}

func (mgr *WordManager) getLine() string {
	maxIndex := len(mgr.words) - 1
	rstIndex := rnd(0, maxIndex)
	rst := mgr.words[rstIndex]

	return rst
}

//LoadFromFile is the loader func of word manager
func (mgr *WordManager) LoadFromFile(filename string) error {
	mgr.words = []string{}
	f, err := os.Open(filename)
	if nil != err {
		panic("file not readable:" + err.Error())
	}
	defer f.Close()
	reader := csv.NewReader(f)
	mgr.isDataSingleChar = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if nil != err {
			return err
		}

		if 1 < len([]rune(record[0])) {
			mgr.isDataSingleChar = false
		}

		mgr.words = append(mgr.words, strings.TrimSpace(record[0]))
	}

	return nil
}
