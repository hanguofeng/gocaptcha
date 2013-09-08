// CServer project main.go
package main

import (
	"fmt"
	"github.com/hanguofeng/gocaptcha"
	"image/png"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	ccaptcha *gocaptcha.Captcha
)

func ShowImageHandler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	if len(key) >= 0 {
		cimg, err := ccaptcha.GetImage(key)
		log.Println("err", err)
		if nil == err {
			w.Header().Add("Content-Type", "image/png")
			png.Encode(w, cimg)
		} else {
			fmt.Fprintf(w, "ERR:%s", err)
		}
	}
}

func ShowKey(w http.ResponseWriter, r *http.Request) {
	key := ccaptcha.GetKey(4)
	w.Write([]byte(key))
}

func ShowPage(w http.ResponseWriter, r *http.Request) {
	key := ccaptcha.GetKey(4)
	page := "<html><body><form method=post action=/verify>key:<input name=key value=" + key + "><br>val:<input name=val><br><img src=/getimage?key=" + key + "><br><input type=submit value=go></form></body></html>"
	w.Write([]byte(page))
}

func DoVerify(w http.ResponseWriter, r *http.Request) {

	page := ""
	key := r.FormValue("key")
	val := r.FormValue("val")
	page = "key:" + key + "<br>"
	page = page + "val:" + val + "<br>"
	suc, msg := ccaptcha.Verify(key, val)
	page = page + fmt.Sprintf("%s,%s", suc, msg)

	w.Header().Add("content-type", "text/html")
	w.Write([]byte(page))
}

func main() {

	pwd, err := os.Getwd()
	if (nil != err) || "" == pwd {
		return
	}
	path := pwd + "/../../data/en_phrases"

	wordmgr := new(gocaptcha.WordManager)
	wordmgr.LoadFromFile(path)
	captchaConfig, imageConfig, filterConfig := loadConfig()
	ccaptcha = gocaptcha.CreateCaptcha(wordmgr, captchaConfig, imageConfig, filterConfig)

	http.HandleFunc("/getimage", ShowImageHandler)
	http.HandleFunc("/getkey", ShowKey)
	http.HandleFunc("/verify", DoVerify)
	http.HandleFunc("/", ShowPage)

	s := &http.Server{Addr: ":8080"}
	log.Fatal(s.ListenAndServe())
}
func loadConfig() (*gocaptcha.CaptchaConfig, *gocaptcha.ImageConfig, *gocaptcha.FilterConfig) {

	captchaConfig := new(gocaptcha.CaptchaConfig)
	captchaConfig.CaptchaLifeTime = time.Minute

	imageConfig := new(gocaptcha.ImageConfig)
	imageConfig.FontFiles = []string{
		"c:/windows/fonts/SIMLI.TTF",
		"c:/windows/fonts/simfang.ttf",
		"c:/windows/fonts/SIMYOU.TTF",
		"c:/windows/fonts/msyh.TTF",
		"c:/windows/fonts/simhei.ttf",
		"c:/windows/fonts/simkai.ttf"}
	imageConfig.FontSize = 26
	imageConfig.Height = 40
	imageConfig.Width = 120

	filterConfig := new(gocaptcha.FilterConfig)
	filterConfig.EnableNoiseLine = true
	filterConfig.EnableNoisePoint = true
	filterConfig.EnableStrike = true
	filterConfig.StrikeLineNum = 3
	filterConfig.NoisePointNum = 30
	filterConfig.NoiseLineNum = 10

	return captchaConfig, imageConfig, filterConfig
}
