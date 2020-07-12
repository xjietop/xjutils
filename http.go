package xjutils

import (
	"io/ioutil"
	"net/http"
)

func HttpGetStr(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	//fmt.Println(string(body))
	return string(body), nil
}
