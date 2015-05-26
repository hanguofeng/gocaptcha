***
View in [[English](README-en.md)][[中文](README.md)]
***
# gocaptcha
Captcha server writen in golang

[![Build Status](https://travis-ci.org/hanguofeng/gocaptcha.png?branch=master)](https://travis-ci.org/hanguofeng/gocaptcha)  [![Build Status](https://drone.io/github.com/hanguofeng/gocaptcha/status.png)](https://drone.io/github.com/hanguofeng/gocaptcha/latest)   [![Coverage Status](https://coveralls.io/repos/hanguofeng/gocaptcha/badge.png)](https://coveralls.io/r/hanguofeng/gocaptcha)  

Feature
-------
* supports captcha char in Chinese
* supports self-define word/char dictionary
* supports filter plugin
* filters：
	* noise point
	* noise line
	* other type of noise
	* plugin
* supports extensible store engine
	* build-in store engine
	* memcache
	* redis (from https://github.com/dtynn/gocaptcha)
	* implement your own by implement the StoreInterface

Useage
------
**Install**

	go get github.com/hanguofeng/gocaptcha

**Quick Start**

See [captcha_test.go](captcha_test.go)

See [samples/gocaptcha-server](samples/gocaptcha-server)

[Demo](http://hanguofeng-gocaptcha.daoapp.io/)

**Document**

[[captcha.go Wiki](https://github.com/hanguofeng/gocaptcha/wiki)]

TODO
----
* ops tools

LICENCE
-------
gocaptcha use [[MIT LICENSE](LICENSE)]

Thanks:

* https://github.com/dchest/captcha
* https://code.google.com/p/freetype-go/freetype
* https://github.com/bradfitz/gomemcache
* https://code.google.com/p/zpix/
