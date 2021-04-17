package native

import (
	"rinne.dev/payjs/entity"
	"testing"
)

func TestRequest_JSON(t *testing.T) {
	req := &Request{
		TotalFee:   123,
		OutTradeNo: "out_trade_no",
		Body:       "body",
		Attach:     "attach",
		NotifyURL:  "notify_url",
		Type:       "type",
		NoCredit:   "no_credit",
		TimeExpire: "time_expire",
	}
	d := req.JSON()
	if d != `{"total_fee":123,"out_trade_no":"out_trade_no","body":"body","attach":"attach","notify_url":"notify_url","type":"type","no_credit":"no_credit","time_expire":"time_expire"}` {
		t.Errorf("native.Request_JSON test failed\n")
	}
}

func TestGetResponse(t *testing.T) {
	resp := &entity.Response{
		ReturnCode: 0,
		ReturnMsg:  "",
		Sign:       "",
		Data:       `{"payjs_order_id": "1"}`,
	}
	ans := GetResponse(resp)
	if ans.OrderID != "1" {
		t.Errorf("native.Request_JSON test failed\n")
	}
}
