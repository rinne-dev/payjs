package native

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"rinne.dev/payjs/entity"
	"testing"
)

func TestCreateRequest(t *testing.T) {
	nativeRequest := CreateRequest(
		123,
		"out_trade_no",
		"body",
		"attach",
		"notify_url",
		"type",
		"no_credit",
		"time_expire",
	)
	data, _ := json.Marshal(nativeRequest)
	request := entity.NewRequest("url", "merch_id", string(data))
	_ = request.Signature("test_key")
	o, _ := request.JSON()
	if gjson.Get(o, "sign").String() != "9C491E4A5A84F4ABB0C9B182E81951A6" {
		t.Errorf("native.CreateRequest test failed\n")
	}
}
