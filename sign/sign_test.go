package sign

import "testing"

const TestSignData = `{"test2": "hehe", "test": 123}`
const TestSignResult = `B16319F0C3F2BBDA5CB8B4FFC9048A81`

func TestSign(t *testing.T) {
	res, _ := Sign(TestSignData, "test_key")
	if res != TestSignResult {
		t.Errorf("sign.Sign expected be %s, but %s got\n", TestSignResult, res)
	}
}
