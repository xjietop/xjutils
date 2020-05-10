package xjutils

import (
	"io/ioutil"
	"net/http"
)

func HttpGetStr(url string) string{
	resp, err :=   http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	//fmt.Println(string(body))
	return string(body)
}
