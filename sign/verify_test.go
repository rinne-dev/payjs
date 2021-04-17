package sign

import (
	"github.com/tidwall/sjson"
	"testing"
)

const TestVerifyData = `{"test2": "hehe", "test": 123, "sign": "B16319F0C3F2BBDA5CB8B4FFC9048A81"}`

func TestVerify(t *testing.T) {
	err := Verify(TestVerifyData, "test_key")
	if err != nil {
		t.Errorf("sign.Verify expected passed\n")
	}
}

func TestVerify2(t *testing.T) {
	newJSON, _ := sjson.Set(TestSignData, "sign", "test")
	err := Verify(newJSON, "test_key")
	if err == nil {
		t.Errorf("sign.Verify expected not passed\n")
	}
}

func TestVerify3(t *testing.T) {
	newJSON, _ := sjson.Delete(TestSignData, "sign")
	err := Verify(newJSON, "test_key")
	if err == nil {
		t.Errorf("sign.Verify expected not passed\n")
	}
}
