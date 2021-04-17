package util

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// PostJSON 发送 JSON Post 请求
func PostJSON(u, data string) (string, error) {
	// 发送请求
	requestData := bytes.NewBuffer([]byte(data))
	resp, err := http.Post(u, "application/json; charset=utf-8", requestData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 判断请求是否成功
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("HTTP 请求错误："+strconv.Itoa(resp.StatusCode))
	}

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 返回
	return string(body), nil
}
