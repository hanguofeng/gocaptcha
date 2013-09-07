***
View in [[English](README-en.md)][[中文](README.md)]
***
# gocaptcha
go语言验证码服务器

[![Build Status](https://travis-ci.org/hanguofeng/gocaptcha.png?branch=master)](https://travis-ci.org/hanguofeng/gocaptcha)

Feature
-------
* 支持中文验证码
* 支持自定义词库、字库
* 支持滤镜机制，通过滤镜来增加干扰，加大识别难度
* 当前的滤镜包括：
	* 支持干扰点
	* 支持干扰线
	* 支持其他模式的干扰

Useage
------
**安装**

	go get github.com/hanguofeng/gocaptcha

**Quick Start**

参考 [captcha_test.go](captcha_test.go)

**文档**

[[captcha.go Wiki](https://github.com/hanguofeng/gocaptcha/wiki)]

TODO
----
* 支持自定义滤镜机制
* 支持单机、集群部署
* 运维管理工具

LICENCE
-------
gocaptcha使用[[MIT许可协议](LICENSE)]

