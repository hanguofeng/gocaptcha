// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	enc "encoding/gob"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	MC_KEY_PREFIX = "gocaptcha-"
)

//MCStore is the Captcha info store service
type MCStore struct {
	engine string
	mc     *memcache.Client
}

//CreateCStore will create a new CStore
func CreateMCStore(expiresTime time.Duration, servers []string) *MCStore {
	store := new(MCStore)

	store.mc = memcache.New(servers...)
	store.mc.Timeout = expiresTime

	return store
}

//Get captcha info by key
func (store *MCStore) Get(key string) *CaptchaInfo {
	item, err := store.mc.Get(MC_KEY_PREFIX + key)
	if nil != err {
		return nil
	}
	ret := store.decodeValue(item.Value)
	return ret
}

//Add captcha info and get the auto generated key
func (store *MCStore) Add(captcha *CaptchaInfo) string {

	key := fmt.Sprintf("%s%s%x", captcha.Text, randStr(20), time.Now().UnixNano())
	key = hex.EncodeToString(md5.New().Sum([]byte(key)))
	key = key[:32]

	item := new(memcache.Item)
	item.Key = MC_KEY_PREFIX + key
	item.Value = store.encodeValue(captcha)

	err := store.mc.Add(item)

	if nil != err {
		log.Printf("add key in memcache err:%s", err)
	}

	return key
}

//Update captcha info
func (store *MCStore) Update(key string, captcha *CaptchaInfo) bool {
	item := new(memcache.Item)
	item.Key = MC_KEY_PREFIX + key
	item.Value = store.encodeValue(captcha)

	err := store.mc.Set(item)

	if nil != err {
		log.Printf("update key in memcache err:%s", err)
		return false
	} else {
		return true
	}
}

//Del captcha info by key
func (store *MCStore) Del(key string) {
	store.mc.Delete(MC_KEY_PREFIX + key)
}

//Destroy the whole store
func (store *MCStore) Destroy() {

}

//OnConstruct load data
func (store *MCStore) OnConstruct() {

}

//OnDestruct dump data
func (store *MCStore) OnDestruct() {

}

func (store *MCStore) encodeValue(val *CaptchaInfo) []byte {

	buf := bytes.NewBufferString("")
	encoder := enc.NewEncoder(buf)
	err := encoder.Encode(val)
	if nil != err {
		return nil
	}
	return buf.Bytes()
}

func (store *MCStore) decodeValue(value []byte) *CaptchaInfo {
	data := &CaptchaInfo{}
	buf := bytes.NewBuffer(value)
	decoder := enc.NewDecoder(buf)
	decoder.Decode(data)

	return data
}
