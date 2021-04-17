package sign

import "testing"

const TestSignData = `{"test2": "hehe", "test": 123}`
const TestSignResult = `B16319F0C3F2BBDA5CB8B4FFC9048A81`

func TestSign(t *testing.T) {
	res, err := Sign(TestSignData, "test_key")
	if err != nil {
		t.Errorf("Sign error: %s\n", err.Error())
	}
	if res != TestSignResult {
		t.Errorf("Sign expected be %s, but %s got\n", TestSignResult, res)
	}
}
