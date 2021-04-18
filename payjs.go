package payjs

import (
	"errors"
	"rinne.dev/payjs/config"
	"rinne.dev/payjs/entity"
	"rinne.dev/payjs/native"
	"rinne.dev/payjs/util"
)

// BaseURL API 地址
const BaseURL = "https://payjs.cn/api"

// PayJS 实例
type PayJS struct {
	Config  *config.Config
	Request *entity.Request
}

// New 新建 PayJS 实例
func New(merchID, key string) *PayJS {
	return &PayJS{
		Config: &config.Config{
			MerchID: merchID,
			Key:     key,
		},
		Request: nil,
	}
}

// Send 发送请求
func (p *PayJS) Send() (*entity.Response, error) {
	if p.Request == nil {
		return nil, errors.New("无请求")
	}

	// 发送请求
	err := p.Request.Signature(p.Config.Key)
	if err != nil {
		return nil, err
	}
	data, err := p.Request.JSON()
	if err != nil {
		return nil, err
	}
	resp, err := util.PostJSON(BaseURL+p.Request.URL, data)
	if err != nil {
		return nil, err
	}

	return entity.ParseResponse(resp), nil
}

// Native 创建扫码支付请求
func (p *PayJS) Native(req *native.Request) (*native.Response, error) {
	// 构造请求
	p.Request = entity.NewRequest("/native", p.Config.MerchID, req.JSON())

	// 发送请求
	resp, err := p.Send()
	if err != nil {
		return nil, err
	}

	// 返回响应
	return native.GetResponse(resp), nil
}
