// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

//WordManager is a captcha word manage tool
type WordManager struct {
	words            []string
	isDataSingleChar bool
	isValid          bool
}

//CreateWordManagerFromDataFile will create a entity from a dictionary file
func CreateWordManagerFromDataFile(filename string) (*WordManager, error) {
	mgr := &WordManager{}
	mgr.words = []string{}
	mgr.isValid = false

	f, err := os.Open(filename)
	if nil != err {
		return mgr, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	mgr.isDataSingleChar = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if nil != err {
			return mgr, err
		}

		if 1 < len([]rune(record[0])) {
			mgr.isDataSingleChar = false
		}

		mgr.words = append(mgr.words, strings.TrimSpace(record[0]))
	}

	mgr.isValid = true
	return mgr, nil
}

//Get a specifical length word
func (mgr *WordManager) Get(length int) (string, error) {
	var retErr error
	rst := ""
	if mgr.isValid {
		if true == mgr.isDataSingleChar {
			if len(mgr.words) < length {
				return "", errors.New("dict words count is less than your length")
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
	} else {
		retErr = errors.New("WordManager is invalid")
	}

	return rst, retErr
}

func (mgr *WordManager) SetWords(words []string) {
	mgr.words = words
	mgr.isValid = len(words) > 0
	mgr.isDataSingleChar = true
	for _, s := range words {
		if len([]rune(s)) > 1 {
			mgr.isDataSingleChar = false
		}
	}
}

func (mgr *WordManager) getLine() string {
	maxIndex := len(mgr.words) - 1
	rstIndex := rnd(0, maxIndex)
	rst := mgr.words[rstIndex]

	return rst
}
