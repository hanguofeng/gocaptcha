***
View in [[English](README-en.md)][[中文](README.md)]
***
# gocaptcha
go语言验证码服务器

Feature
-------
* 支持中文验证码
* 支持自定义词库、字库
* 支持自定义滤镜机制，通过滤镜来增加干扰，加大识别难度
* 当前的滤镜包括：
	* 支持干扰点
	* 支持干扰线
	* 支持其他模式的干扰
	* 更多模式，可实现imagefilter接口来扩展
* 支持自定义存储引擎，存储引擎可扩展
* 目前支持的存储引擎包括:
	* 内置(buildin)
	* memcache
	* redis (from https://github.com/dtynn/gocaptcha)
	* 如需扩展存储引擎，可实现StoreInterface接口

Useage
------
**安装**

	go get github.com/hanguofeng/gocaptcha

**Quick Start**

参考 [captcha_test.go](captcha_test.go)

参考 [samples/gocaptcha-server](samples/gocaptcha-server)

[Demo](http://hanguofeng-gocaptcha.daoapp.io/)

**文档**

[[captcha.go Wiki](https://github.com/hanguofeng/gocaptcha/wiki)]

TODO
----
* 运维管理工具

LICENCE
-------
gocaptcha使用[[MIT许可协议](LICENSE)]


使用的开源软件列表，表示感谢

* https://github.com/dchest/captcha
* https://github.com/golang/freetype
* https://github.com/bradfitz/gomemcache
* https://code.google.com/p/zpix/
