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
)

const (
	GC_PROBABILITY = 1
	GC_DIVISOR     = 100
)

type CStore struct {
	mu          sync.RWMutex
	data        map[string]*CaptchaInfo
	expiresTime time.Duration
}

func CreateCStore(expiresTime time.Duration) *CStore {
	store := new(CStore)
	store.data = make(map[string]*CaptchaInfo)
	store.expiresTime = expiresTime

	return store
}

func (store *CStore) Get(key string) *CaptchaInfo {
	defer store.gcWrapper() //run gc after get
	store.mu.Lock()
	defer store.mu.Unlock()

	return store.data[key]
}

func (store *CStore) Add(captcha *CaptchaInfo) string {
	store.mu.Lock()
	defer store.mu.Unlock()
	key := fmt.Sprintf("%s%s%x", captcha.text, randStr(20), time.Now().UnixNano())
	key = hex.EncodeToString(md5.New().Sum([]byte(key)))
	store.data[key] = captcha
	return key
}
func (store *CStore) Del(key string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	delete(store.data, key)
}

func (store *CStore) Destroy(key string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	for key, _ := range store.data {
		store.Del(key)
	}
}

func (store *CStore) gcWrapper() {
	//run PROBABILITY
	if rnd(0, GC_DIVISOR) == GC_PROBABILITY {
		go store.gc()
	}
}

func (store *CStore) gc() {
	for key, val := range store.data {
		if val.createTime.Add(store.expiresTime).Before(time.Now()) {
			log.Println("collecting:" + key)
			store.Del(key)
		}
	}
}
