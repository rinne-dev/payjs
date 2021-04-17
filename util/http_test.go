package util

import (
	"github.com/tidwall/gjson"
	"testing"
)

func TestPostJSON(t *testing.T) {
	resp, err := PostJSON("https://httpbin.org/post", `{"test": "haha"}`)
	if err != nil {
		t.Errorf("util.PostJSON failed: %s\n", err.Error())
	}
	if !gjson.Get(resp, "json.test").Exists() {
		t.Errorf("util.PostJSON failed\n")
	}
	if gjson.Get(resp, "json.test").String() != "haha" {
		t.Errorf("util.PostJSON failed\n")
	}
}
