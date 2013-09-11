***
View in [[English](README-en.md)][[中文](README.md)]
***
# gocaptcha
Captcha server writen in golang

[![Build Status](https://travis-ci.org/hanguofeng/gocaptcha.png?branch=master)](https://travis-ci.org/hanguofeng/gocaptcha)

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

Useage
------
**Install**

	go get github.com/hanguofeng/gocaptcha

**Quick Start**

See [captcha_test.go](captcha_test.go)

**Document**

[[captcha.go Wiki](https://github.com/hanguofeng/gocaptcha/wiki)]

TODO
----
* support deployment in stand-alone or cluster
* ops tools

LICENCE
-------
gocaptcha use [[MIT LICENSE](LICENSE)]

