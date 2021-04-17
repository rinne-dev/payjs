package util

import (
	"testing"
)

const TestQueryStringData = `{
"null": null,
"str": "str test",
"int": 2,
"float": 1.23,
"bool": true,
"bool2": false,
"arr": ["test", 123, 1.23],
"obj": {"aaa": "a", "bbb": 123, "ccc": 1.23},
"complex": {"test111": {"test222": {"hehe": "haha"}}, "test222": "haha"}
}`
const TestQueryStringResult = `arr[0]=test&arr[1]=123&arr[2]=1.23&bool=1&bool2=0&complex[test111][test222][hehe]=haha&complex[test222]=haha&float=1.23&int=2&obj[aaa]=a&obj[bbb]=123&obj[ccc]=1.23&str=str test`

func TestQueryString(t *testing.T) {
	res, _ := QueryString(TestQueryStringData)
	if res != TestQueryStringResult {
		t.Errorf("util.QueryString expected be %s, but %s got\n", TestQueryStringResult, res)
	}
}
