package util

import (
	"errors"
	"github.com/tidwall/gjson"
	"sort"
	"strconv"
	"strings"
)

// QueryString 根据 JSON 字符串构造 QueryString
func QueryString(data string) (string, error) {
	// 判断 JSON 字符串是否合法
	if !gjson.Valid(data) {
		return "", errors.New("传入的 JSON 字符串不合法")
	}

	// 解析 JSON 字符串
	result := parse(data)

	// 按照 Key 对结果排序
	keys := make([]string, 0)
	for k, _ := range result {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构造结果
	resultArr := make([]string, 0)
	for _, k := range keys {
		resultArr = append(resultArr, k+"="+result[k])
	}

	// 返回结果
	return strings.Join(resultArr, "&"), nil
}

func parse(data string) map[string]string {
	// 初始化结果串
	result := make(map[string]string, 0)
	// 解析 JSON 字符串
	obj := gjson.Parse(data)
	// 遍历所有对象
	obj.ForEach(func(key, value gjson.Result) bool {
		itemKey := key.String()
		if key.Type == gjson.Null {
			// 数组的 Key 是 Null，需要单独处理
			itemKey = strconv.Itoa(len(result))
		}
		if value.Type == gjson.JSON {
			// 值类型是 JSON 对象或数组
			res := parse(value.String())
			for k, v := range res {
				// 可能存在多级 Key，存在的话直接放到后面即可
				if strings.Contains(k, "[") {
					idx := strings.Index(k, "[")
					result[itemKey+"["+k[:idx]+"]"+k[idx:]] = v
				} else {
					result[itemKey+"["+k+"]"] = v
				}
			}
		}
		if value.Type == gjson.String || value.Type == gjson.Number {
			// 值类型是字符串、整数或浮点数
			result[itemKey] = value.String()
		}
		if value.Type == gjson.True || value.Type == gjson.False {
			// 值类型是布尔
			result[itemKey] = strconv.Itoa(int(value.Int()))
		}
		if value.Type == gjson.Null {
			// 值类型是空，直接跳过
		}
		return true
	})

	// 返回结果
	return result
}
