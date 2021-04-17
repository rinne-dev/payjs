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
	request := entity.NewRequest("merch_id", string(data))
	_ = request.Signature("test_key")
	o, _ := request.JSON()
	if gjson.Get(o, "sign").String() != "929001000398AEEFC3429512BF1E8426" {
		t.Errorf("native.CreateRequest test failed\n")
	}
}
