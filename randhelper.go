// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// rnd returns a non-crypto pseudorandom int in range [from, to].
func rnd(from, to int) int {
	return rand.Intn(to+1-from) + from
}

// rndf returns a non-crypto pseudorandom float64 in range [from, to].
func rndf(from, to float64) float64 {
	return (to-from)*rand.Float64() + from
}

func randStr(length int) string {
	rst := ""
	charset := "abcdefghijklmnopqrstuvwxyz1234567890"
	for i := 0; i < length; i++ {
		rst = rst + string(charset[rnd(0, len(charset)-1)])
	}
	return rst
}
