package entity

import (
	"testing"
)

func TestNewRequest(t *testing.T) {
	req := NewRequest("url", "merch_id", "")
	if req.MerchID != "merch_id" {
		t.Errorf("entity.NewRequest test failed\n")
	}
	if req.Data != "{}" {
		t.Errorf("entity.NewRequest test failed\n")
	}
}

func TestRequest_Signature(t *testing.T) {
	req := NewRequest("url", "merch_id", `{"test": "haha"}`)
	_ = req.Signature("test_key")
	if req.Sign != "8C25F4D8CE652EE13397915BC2F8A825" {
		t.Errorf("entity.Request_Signature test failed\n")
	}
}

func TestRequest_JSON(t *testing.T) {
	req := NewRequest("url", "merch_id", `{"test": "haha"}`)
	_ = req.Signature("test_key")
	d, _ := req.JSON()
	if d != "{\"test\": \"haha\",\"mchid\":\"merch_id\",\"sign\":\"8C25F4D8CE652EE13397915BC2F8A825\"}" {
		t.Errorf("entity.Request_JSON test failed\n")
	}
}
