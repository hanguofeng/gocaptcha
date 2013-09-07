// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type WordManager struct {
	words            []string
	isDataSingleChar bool
}

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

		rst_rune := []rune(rst)
		rst = string(rst_rune[0:length])
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

	log.Printf("Load Complete,totle : %d,isDataSingleChar:%s", len(mgr.words), mgr.isDataSingleChar)

	return nil
}
