// CServer project main.go
package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"net/http"

	"github.com/hanguofeng/gocaptcha"
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
	ccaptcha = gocaptcha.CreateCaptchaFromConfigFile(*configFile)

	http.HandleFunc("/getimage", ShowImageHandler)
	http.HandleFunc("/getkey", ShowKey)
	http.HandleFunc("/verify", DoVerify)
	http.HandleFunc("/", ShowPage)

	s := &http.Server{Addr: ":8080"}
	log.Fatal(s.ListenAndServe())
}
