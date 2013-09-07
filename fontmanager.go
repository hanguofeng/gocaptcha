// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"code.google.com/p/freetype-go/freetype/truetype"
	"io/ioutil"
	"math/rand"
	"time"
)

type FontManager struct {
	fontFiles   []string
	fontObjects map[string]*truetype.Font
	randObject  *rand.Rand
}

func CreateFontManager() *FontManager {
	fm := new(FontManager)
	fm.fontFiles = []string{}
	fm.fontObjects = make(map[string]*truetype.Font)
	fm.randObject = rand.New(rand.NewSource(time.Now().UnixNano()))

	return fm
}

func (fm *FontManager) AddFont(pathToFontFile string) error {
	fontBytes, err := ioutil.ReadFile(pathToFontFile)
	if err != nil {
		return err
	}

	font, err := truetype.Parse(fontBytes)

	if err != nil {
		return err
	}
	fm.fontFiles = append(fm.fontFiles, pathToFontFile)
	fm.fontObjects[pathToFontFile] = font

	return nil
}

func (fm *FontManager) GetFont(pathToFontFile string) *truetype.Font {
	return fm.fontObjects[pathToFontFile]
}

func (fm *FontManager) GetRandomFont() *truetype.Font {
	randomIndex := fm.randObject.Intn(len(fm.fontFiles))
	fontFile := fm.fontFiles[randomIndex]
	rst := fm.GetFont(fontFile)

	return rst
}
