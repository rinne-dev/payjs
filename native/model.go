package native

import (
	"encoding/json"
	"rinne.dev/payjs/entity"
)

// Request 扫码支付请求参数
type Request struct {
	TotalFee   int64  `json:"total_fee"`             // Y 金额。单位：分
	OutTradeNo string `json:"out_trade_no"`          // Y 用户端自主生成的订单号
	Body       string `json:"body,omitempty"`        // N 订单标题
	Attach     string `json:"attach,omitempty"`      // N 用户自定义数据，在 notify 的时候会原样返回
	NotifyURL  string `json:"notify_url,omitempty"`  // N 接收异步通知的回调地址
	Type       string `json:"type,omitempty"`        // N 支付宝交易传值：alipay，微信支付无需此字段
	NoCredit   string `json:"no_credit,omitempty"`   // N 禁止使用信用卡支付请传值：1
	TimeExpire string `json:"time_expire,omitempty"` // N 订单有效期。格式：20200826080808
}

// Response 扫码支付响应参数
type Response struct {
	OrderID    string `json:"payjs_order_id"` // Y PayJS 平台订单号
	OutTradeNo string `json:"out_trade_no"`   // Y 用户生成的订单号原样返回
	TotalFee   int64  `json:"total_fee"`      // Y 金额。单位：分
	CodeURL    string `json:"code_url"`       // Y 二维码内容（有效期2小时）
	QRCode     string `json:"qrcode"`         // Y 二维码图片地址
}

// JSON 生成请求 JSON 对象
func (r *Request) JSON() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// GetResponse 取得响应
func GetResponse(resp *entity.Response) *Response {
	if resp.ReturnCode != 1 {
		return nil
	}
	r := Response{}
	_ = json.Unmarshal([]byte(resp.Data), &r)
	return &r
}
