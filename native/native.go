package native

// CreateRequest 创建扫码支付请求
func CreateRequest(
	totalFee int64,
	outTradeNo string,
	body string,
	attach string,
	notifyURL string,
	apiType string,
	noCredit string,
	timeExpire string,
) *Request {
	return &Request{
		TotalFee:   totalFee,
		OutTradeNo: outTradeNo,
		Body:       body,
		Attach:     attach,
		NotifyURL:  notifyURL,
		Type:       apiType,
		NoCredit:   noCredit,
		TimeExpire: timeExpire,
	}
}
