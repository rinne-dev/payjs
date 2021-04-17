package entity

import (
	"testing"
)

func TestNewRequest(t *testing.T) {
	req := NewRequest("merch_id", "")
	if req.MerchID != "merch_id" {
		t.Errorf("entity.NewRequest test failed\n")
	}
	if req.Data != "{}" {
		t.Errorf("entity.NewRequest test failed\n")
	}
}

func TestRequest_Signature(t *testing.T) {
	req := NewRequest("merch_id", `{"test": "haha"}`)
	_ = req.Signature("test_key")
	if req.Sign != "1278F033181E84E6A7DECBBD767B7084" {
		t.Errorf("entity.Request_Signature test failed\n")
	}
}

func TestRequest_JSON(t *testing.T) {
	req := NewRequest("merch_id", `{"test": "haha"}`)
	_ = req.Signature("test_key")
	d, _ := req.JSON()
	if d != "{\"test\": \"haha\",\"merch_id\":\"merch_id\",\"sign\":\"1278F033181E84E6A7DECBBD767B7084\"}" {
		t.Errorf("entity.Request_JSON test failed\n")
	}
}
