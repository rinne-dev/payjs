package entity

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"rinne.dev/payjs/sign"
)

// Request 请求结构
type Request struct {
	MerchID string `json:"merch_id"` // 商户号
	Sign    string `json:"sign"`     // 数据签名
	Data    string `json:"-"`        // 数据
}

// NewRequest 创建请求
func NewRequest(merchID, data string) *Request {
	if !gjson.Parse(data).IsObject() {
		data = "{}"
	}
	return &Request{
		MerchID: merchID,
		Sign:    "",
		Data:    data,
	}
}

// Signature 计算签名
func (r *Request) Signature(key string) error {
	data, err := sjson.Set(r.Data, "merch_id", r.MerchID)
	if err != nil {
		return err
	}
	signature, err := sign.Sign(data, key)
	if err != nil {
		return err
	}
	r.Sign = signature
	return nil
}

// JSON 组合请求对象
func (r *Request) JSON() (string, error) {
	data, err := sjson.Set(r.Data, "merch_id", r.MerchID)
	if err != nil {
		return "", err
	}
	data, err = sjson.Set(data, "sign", r.Sign)
	if err != nil {
		return "", err
	}
	return data, nil
}
