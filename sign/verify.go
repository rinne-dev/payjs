package sign

import (
	"errors"
	"github.com/tidwall/gjson"
)

// Verify 验签
func Verify(data, key string) error {
	// 提取 JSON 对象中的签名
	signature := gjson.Get(data, "sign")
	if !signature.Exists() {
		return errors.New("找不到待验证的签名字段")
	}

	// 本地生成签名
	localSignature, err := Sign(data, key)
	if err != nil {
		return err
	}

	// 校验签名
	if signature.String() == localSignature {
		return nil
	} else {
		return errors.New("签名校验不通过")
	}
}
