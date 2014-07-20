// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"github.com/hanguofeng/freetype-go-mirror/freetype/truetype"
	"io/ioutil"
	"math/rand"
	"time"
)

//FontManager is a Font Manager the manage font list
type FontManager struct {
	fontFiles   []string
	fontObjects map[string]*truetype.Font
	randObject  *rand.Rand
}

//CreateFontManager will create a new Font Manager
func CreateFontManager() *FontManager {
	fm := new(FontManager)
	fm.fontFiles = []string{}
	fm.fontObjects = make(map[string]*truetype.Font)
	fm.randObject = rand.New(rand.NewSource(time.Now().UnixNano()))

	return fm
}

//AddFont will add a new font file to the list
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

//GetFont will return a Font struct by path
func (fm *FontManager) GetFont(pathToFontFile string) *truetype.Font {
	return fm.fontObjects[pathToFontFile]
}

//GetRandomFont will return a random Font struct
func (fm *FontManager) GetRandomFont() *truetype.Font {
	randomIndex := fm.randObject.Intn(len(fm.fontFiles))
	fontFile := fm.fontFiles[randomIndex]
	rst := fm.GetFont(fontFile)

	return rst
}
