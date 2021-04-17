package entity

import (
	"testing"
)

func TestParseResponse(t *testing.T) {
	data := `{"return_code": 1, "return_msg": "msg", "sign": "test_sign", "test": "haha"}`
	resp := ParseResponse(data)
	if resp.ReturnCode != 1 {
		t.Errorf("entity.ParseResponse test failed\n")
	}
	if resp.ReturnMsg != "msg" {
		t.Errorf("entity.ParseResponse test failed\n")
	}
	if resp.Sign != "test_sign" {
		t.Errorf("entity.ParseResponse test failed\n")
	}
}

func TestResponse_Get(t *testing.T) {
	data := `{"return_code": 1, "return_msg": "msg", "sign": "test_sign", "test": "haha"}`
	resp := ParseResponse(data)
	if resp.Get("test").String() != "haha" {
		t.Errorf("entity.ParseResponse test failed\n")
	}
}
