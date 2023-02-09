package u_str

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"strings"
)

// Str2Int64 字符串转int64
func Str2Int64(v string) int64 {
	valInt, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		valInt = 0
	}
	return valInt
}

// FirstStr 获取首个不为空的字符串
func FirstStr(vLst ...interface{}) string {
	for _, v := range vLst {
		if &v != nil && len(v.(string)) > 0 {
			return v.(string)
		}
	}
	return ""
}

// TrimNewlineSpace 去除所有换行跟空格
func TrimNewlineSpace(s string) string {
	if len(s) == 0 {
		return ""
	}
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, " ", "")
	return s
}

// Unicode2Str unicode字符串转string
func Unicode2Str(form string, defaultValue string) (to string) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return defaultValue
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return to
}
