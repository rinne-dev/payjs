package sign

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"rinne.dev/payjs/util"
	"strings"
)

// Sign 签名
func Sign(data, key string) (string, error) {
	var err error

	// 去除 sign 字段
	if gjson.Get(data, "sign").Exists() {
		data, err = sjson.Delete(data, "sign")
		if err != nil {
			return "", err
		}
	}

	// 获取 query string
	queryString, err := util.QueryString(data)
	if err != nil {
		return "", err
	}
	queryString += "&key="+key

	// 计算签名
	h := md5.New()
	h.Write([]byte(queryString))

	// 返回生成的签名
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil))), nil
}
