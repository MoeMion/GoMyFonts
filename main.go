package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

func cssHandeler(c *gin.Context, link string) {
	url := "https://fonts.googleapis.com/css?family=" + c.Query("family")
	client := &http.Client{}

	reqest, err := http.NewRequest("GET", url, nil)

	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:94.0) Gecko/20100101 Firefox/94.0")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Upgrade-Insecure-Requests", "1")
	reqest.Header.Add("Cache-Control", "max-age=0")
	reqest.Header.Add("Upgrade-Insecure-Requests", "1")

	if err != nil {
		panic(err)
	}

	response, _ := client.Do(reqest)
	cssBody, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	replacedCSS := strings.Replace(string(cssBody), "https://fonts.gstatic.com/", link, -1)
	c.String(200, "%s", replacedCSS)
}

func fontHandeler(c *gin.Context) {
	url := "https://fonts.gstatic.com/s/" + c.Param("res")
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	defer reader.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func usage() {
	fmt.Fprintf(os.Stderr, `GoMyFonts version: 1.0
Usage: gomyfonts [-h] [-p :port] [-l link] [-t title] [-c timeout of cache]
Github: https://github.com/MoeMion/GoMyFonts
Author: Mion
Options:
`)
	flag.PrintDefaults()
}

func main() {
	port := flag.String("p", ":2333", "Bind TCP Port.")
	link := flag.String("l", "http://127.0.0.1/", "The url of your mirror site.")
	title := flag.String("t", "GoMyFont", "The title of your mirror site.")
	cacheTime := flag.Int("c", 10, "Expiration of cache,Unit:Minute.")
	flag.Usage = usage
	flag.Parse()

	web := gin.Default()
	cacheDuration := time.Duration(*cacheTime) * time.Minute
	memoryStore := persist.NewMemoryStore(cacheDuration)

	web.GET("/css", func(c *gin.Context) {
		cssHandeler(c, *link)
	})
	web.GET("/s/*res", cache.CacheByRequestURI(memoryStore, cacheDuration), fontHandeler)
	web.StaticFile("/favicon.ico", "./dist/favicon.ico")
	web.LoadHTMLFiles("dist/index.html", "dist/404.html")
	web.GET("/", cache.CacheByRequestURI(memoryStore, cacheDuration), func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": *title,
			"link":  *link,
		})
	})
	web.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{
			"title": "404",
			"link":  *link,
		})
	})
	web.Run(*port)
}
