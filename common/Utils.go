package common

import (
	"bytes"
	"log"
	"net/http"
)

func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return ""
	}

	defer resp.Body.Close()

	if 200 == resp.StatusCode {
		buf := bytes.NewBuffer(make([]byte, 0, 512))
		length, _ := buf.ReadFrom(resp.Body)

		if len(buf.Bytes()) == int(length) {
			return string(buf.Bytes())
		}
	}

	return ""
}
