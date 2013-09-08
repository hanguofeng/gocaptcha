// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"sync"
	"time"

	enc "encoding/gob"
	"os"
)

//CStore is the Captcha info store service
type CStore struct {
	mu            sync.RWMutex
	data          map[string]*CaptchaInfo
	expiresTime   time.Duration
	gcProbability int
	gcDivisor     int
}

//CreateCStore will create a new CStore
func CreateCStore(expiresTime time.Duration, gcProbability int, gcDivisor int) *CStore {
	store := new(CStore)
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
	store.data[key] = captcha
	return key
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

//Dump the whole store
func (store *CStore) Dump(file string) {
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
				log.Println("Dump encoded", store)
			} else {
				log.Fatalln("[Store Dump Err]", err)
			}
		} else {
			log.Fatalln("[Store Dump Err]", err)
		}
	} else {
		log.Fatalln("[Store Dump Err]", err)
	}
}

//LoadDumped file to store
func (store *CStore) LoadDumped(file string) {
	data := &map[string]*CaptchaInfo{}
	pwd, err := os.Getwd()

	if (nil == err) && ("" != pwd) {
		f, err := os.Open(pwd + "/" + file)
		if nil == err {
			decoder := enc.NewDecoder(f)
			err := decoder.Decode(data)
			f.Close()

			if nil == err {
				log.Println("LoadDumped decode", *data)
				store.data = *data
			} else {
				log.Fatalln("[Store LoadDumped Err]", err)
			}

		} else {
			log.Fatalln("[Store LoadDumped Err]", err)
		}
	} else {
		log.Fatalln("[Store LoadDumped Err]", err)
	}
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
			log.Println("collecting:" + key)
			store.Del(key)
		}
	}
}
