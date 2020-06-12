package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Str2Float32(f string) float32 {
	value, err := strconv.ParseFloat(f, 32)
	if err != nil {
		fmt.Println("some error")
	}
	return float32(value)
}

func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return ""
	}

	defer resp.Body.Close()

	if 200 == resp.StatusCode {
		buf, _ := ioutil.ReadAll(resp.Body)
		return string(buf)
	}

	return ""
}

func HttpPost(url string, contentType string, jsonParams string) string {
	resp, err := http.Post(url, contentType, strings.NewReader(jsonParams))
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()

	if 200 == resp.StatusCode {
		buf, _ := ioutil.ReadAll(resp.Body)
		return string(buf)
	}

	return ""
}
