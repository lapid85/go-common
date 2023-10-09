package utils

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

// HttpGet 获取http内容
func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	// 针对 UTF-8 BOM（字节顺序标记）的处理，通常不需要，而且可能会导致不必要的修改。如果有特殊需求来处理 BOM，可以单独处理。
	// body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	return string(body), err
}

// HttpPost 0 json 1 form
func HttpPost(url string, requestBody []byte, bodyType int) (string, error) {
	var bt string
	if bodyType == 0 {
		bt = "application/json"
	} else if bodyType == 1 {
		bt = "application/form"
	} else {
		return "", errors.New("不支持的body类型")
	}

	resp, err := http.Post(url, bt, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	return string(body), err
}
