package kgo

import (
	"encoding/json"
	"math"
	"regexp"
	"strconv"
	"unicode"
)

// IsBinary 字符串是否二进制
func (ks *LkkString) IsBinary(s string) bool {
	for _, b := range s {
		if 0 == b {
			return true
		}
	}
	return false
}

// IsLetter 字符串是否字母
func (ks *LkkString) IsLetter(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// IsNumeric 字符串是否数值
func (ks *LkkString) IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsNumeric 字符串是否整数
func (ks *LkkString) IsInt(s string) bool {
	if s == "" {
		return false
	}
	_, err := strconv.Atoi(s)
	return err == nil
}

// HasChinese 字符串是否含有中文
func (ks *LkkString) HasChinese(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}

	return false
}

// IsChinese 字符串是否全部中文
func (ks *LkkString) IsChinese(s string) bool {
	if s == "" {
		return false
	}

	return regexp.MustCompile("^[\u4e00-\u9fa5]+$").MatchString(s)
}

// IsJSON 字符串是否合法的json格式
func (ks *LkkString) IsJSON(str string) bool {
	if str == "" {
		return false
	} else if str[0] != '{' || str[len(str)-1] != '}' {
		return false
	}

	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsArrayOrSlice 检查变量是否数组或切片;chkType检查类型,枚举值有(1仅数组,2仅切片,3数组或切片);结果为-1表示非,>=0表示是
func (ka *LkkArray) IsArrayOrSlice(data interface{}, chkType uint8) int {
	return isArrayOrSlice(data, chkType)
}

// IsMap 检查变量是否字典
func (ka *LkkArray) IsMap(data interface{}) bool {
	return isMap(data)
}

// IsNan 是否为“非数值”
func (kn *LkkNumber) IsNan(val float64) bool {
	return math.IsNaN(val)
}
