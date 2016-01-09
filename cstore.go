// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"crypto/md5"
	enc "encoding/gob"
	"encoding/hex"
	"fmt"
	"os"
	"sync"
	"time"
)

//CStore is the Captcha info store service
type CStore struct {
	engine        string
	mu            sync.RWMutex
	data          map[string]*CaptchaInfo
	expiresTime   time.Duration
	gcProbability int
	gcDivisor     int
}

//CreateCStore will create a new CStore
func CreateCStore(expiresTime time.Duration, gcProbability int, gcDivisor int) *CStore {
	store := new(CStore)
	store.engine = STORE_ENGINE_BUILDIN
	store.data = make(map[string]*CaptchaInfo)
	store.expiresTime = expiresTime
	store.gcProbability = gcProbability
	store.gcDivisor = gcDivisor

	return store
}

//Get captcha info by key
func (store *CStore) Get(key string) *CaptchaInfo {
	defer store.gcWrapper() //run gc after get
	store.mu.Lock()
	defer store.mu.Unlock()

	ret := store.data[key]
	return ret
}

//Add captcha info and get the auto generated key
func (store *CStore) Add(captcha *CaptchaInfo) string {
	store.mu.Lock()
	defer store.mu.Unlock()
	key := fmt.Sprintf("%s%s%x", captcha.Text, randStr(20), time.Now().UnixNano())
	key = hex.EncodeToString(md5.New().Sum([]byte(key)))
	key = key[:32]
	store.data[key] = captcha
	return key
}

//Update captcha info
func (store *CStore) Update(key string, captcha *CaptchaInfo) bool {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.data[key] = captcha
	return true
}

//Del captcha info by key
func (store *CStore) Del(key string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	delete(store.data, key)
}

//Destroy the whole store
func (store *CStore) Destroy() {

	for key := range store.data {
		store.Del(key)
	}
}

//OnConstruct load data
func (store *CStore) OnConstruct() {

}

//OnDestruct dump data
func (store *CStore) OnDestruct() {

}

//Dump the whole store
func (store *CStore) Dump(file string) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	pwd, err := os.Getwd()

	if (nil == err) && ("" != pwd) {
		f, err := os.Create(pwd + "/" + file)
		if nil == err {
			encoder := enc.NewEncoder(f)
			err := encoder.Encode(store.data)
			f.Close()

			if nil == err {
				return err
			} else {
				return nil
			}
		} else {
			return err
		}
	} else {
		return err
	}

	return nil
}

//LoadDumped file to store
func (store *CStore) LoadDumped(file string) error {
	data := &map[string]*CaptchaInfo{}
	pwd, err := os.Getwd()

	if (nil == err) && ("" != pwd) {
		f, err := os.Open(pwd + "/" + file)
		if nil == err {
			decoder := enc.NewDecoder(f)
			err := decoder.Decode(data)
			f.Close()

			if nil == err {
				store.data = *data
				return nil

			} else {
				return err
			}

		} else {
			return err
		}
	} else {
		return err
	}
	return err
}

func (store *CStore) gcWrapper() {
	//run PROBABILITY
	if rnd(0, store.gcDivisor) == store.gcProbability {
		go store.gc()
	}
}

func (store *CStore) gc() {
	for key, val := range store.data {
		if val.CreateTime.Add(store.expiresTime).Before(time.Now()) {
			store.Del(key)
		}
	}
}
