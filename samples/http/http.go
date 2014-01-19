// CServer project main.go
package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hanguofeng/config"
	"github.com/hanguofeng/gocaptcha"
)

const (
	DEFAULT_LIFE_TIME = time.Minute * 5
	DEFAULT_FONT_SIZE = 26
	DEFAULT_WIDTH     = 120
	DEFAULT_HEIGHT    = 40
)

var (
	ccaptcha *gocaptcha.Captcha
)

var configFile = flag.String("c", "gocaptcha.conf", "the config file")

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
	page = page + "<br>"
	page = page + fmt.Sprintf("%s", ccaptcha)

	w.Header().Add("content-type", "text/html")
	w.Write([]byte(page))
}

func main() {

	log.Printf("configFile:%s", *configFile)

	pwd, err := os.Getwd()
	if (nil != err) || "" == pwd {
		return
	}
	path := pwd + "/../../data/en_phrases"

	wordmgr := new(gocaptcha.WordManager)
	wordmgr.LoadFromFile(path)
	captchaConfig, imageConfig, filterConfig := loadConfigFromFile()
	ccaptcha = gocaptcha.CreateCaptcha(wordmgr, captchaConfig, imageConfig, filterConfig)

	http.HandleFunc("/getimage", ShowImageHandler)
	http.HandleFunc("/getkey", ShowKey)
	http.HandleFunc("/verify", DoVerify)
	http.HandleFunc("/", ShowPage)

	s := &http.Server{Addr: ":8080"}
	log.Fatal(s.ListenAndServe())
}

func loadConfigFromFile() (*gocaptcha.CaptchaConfig, *gocaptcha.ImageConfig, *gocaptcha.FilterConfig) {
	c, err := config.ReadDefault(*configFile)
	log.Printf("config:%s,err:%s", c, err)

	//captchaConfig
	captchaConfig := new(gocaptcha.CaptchaConfig)
	var lifeTime time.Duration

	cfgLifeTime, err := c.String("captcha", "life_time")
	if nil == err {
		lifeTime, err = time.ParseDuration(cfgLifeTime)
		if nil != err {
			lifeTime = DEFAULT_LIFE_TIME
			log.Printf("time.ParseDuration of config file failed,using default")
		}
	} else {
		lifeTime = DEFAULT_LIFE_TIME
	}

	captchaConfig.CaptchaLifeTime = lifeTime

	//imageConfig
	imageConfig := new(gocaptcha.ImageConfig)
	var fontFiles []string
	cfgFontFiles, err := c.StringMuti("image", "font_files")
	if nil != err {
		log.Fatalf("read config fail,font files:%s", err.Error())
	} else {
		fontFiles = cfgFontFiles
	}

	imageConfig.FontFiles = fontFiles

	var fontSize float64
	cfgFontSize, err := c.Int("image", "font_size")
	if nil != err {
		fontSize = DEFAULT_FONT_SIZE
	} else {
		fontSize = float64(cfgFontSize)
	}

	imageConfig.FontSize = fontSize

	var width int
	cfgWidth, err := c.Int("image", "width")
	if nil != err {
		width = DEFAULT_WIDTH
	} else {
		width = int(cfgWidth)
	}
	imageConfig.Width = width

	var height int
	cfgHeight, err := c.Int("image", "height")
	if nil != err {
		height = DEFAULT_HEIGHT
	} else {
		height = int(cfgHeight)
	}
	imageConfig.Height = height

	//filterConfig
	filterConfig := new(gocaptcha.FilterConfig)
	filterConfig.EnableNoiseLine = true
	filterConfig.EnableNoisePoint = true
	filterConfig.EnableStrike = true
	filterConfig.StrikeLineNum = 3
	filterConfig.NoisePointNum = 30
	filterConfig.NoiseLineNum = 10

	log.Printf("imageConfig:%s", imageConfig)

	return captchaConfig, imageConfig, filterConfig
}
