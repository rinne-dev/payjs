package entity

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Response 响应结构
type Response struct {
	ReturnCode int    `json:"return_code"` // 1:请求成功，0:请求失败
	ReturnMsg  string `json:"return_msg"`  // 返回消息
	Sign       string `json:"sign"`        // 数据签名
	Data       string `json:"-"`           // 数据
}

// ParseResponse 解析响应
func ParseResponse(data string) *Response {
	res := gjson.Parse(data)
	if res.Get("return_code").Exists() {
		data, _ = sjson.Delete(data, "return_code")
	}
	if res.Get("return_msg").Exists() {
		data, _ = sjson.Delete(data, "return_msg")
	}
	if res.Get("sign").Exists() {
		data, _ = sjson.Delete(data, "sign")
	}
	return &Response{
		ReturnCode: int(res.Get("return_code").Int()),
		ReturnMsg:  res.Get("return_msg").String(),
		Sign:       res.Get("sign").String(),
		Data:       data,
	}
}

// Get 获取响应内容
func (r *Response) Get(key string) gjson.Result {
	return gjson.Get(r.Data, key)
}
