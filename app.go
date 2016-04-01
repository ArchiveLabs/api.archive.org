package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var url = "https://archive.org"

// Loads server config settings such as port and, in the future, TLS/SSL
// See: http://web.archive.org/web/20160320143231/http://www.giantflyingsaucer.com/blog/?p=5157
// The two minute guide to using Viper – configuration management with Go (golang)
func loadSettings() {
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No configuration file loaded - using defaults")
	}
	viper.SetDefault("port", ":8000")
	flag.Set("bind", viper.GetString("port"))
}

func search(c web.C, w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	rows := r.URL.Query().Get("rows")
	page := r.URL.Query().Get("page")
	url := fmt.Sprintf("%s/advancedsearch.php?q=%s&output=json&rows=%s&page=%s", url, q, rows, page)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(``)))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	io.WriteString(w, string(body))
}

func main() {
	loadSettings()
	goji.Get("/search", search)
	goji.Serve()
}
