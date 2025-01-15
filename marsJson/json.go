package marsJson

import (
	"bytes"
	"encoding/json"
)

// Marshal json序列化
func Marshal(v any) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(marshal)
}

// UnMarshal 解析json字符串
func UnMarshal(str string, v any) error {
	return json.Unmarshal([]byte(str), v)
}

// PrettyString 格式化json字符串
func PrettyString(str string) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", " "); err != nil {
		return ""
	}
	return prettyJSON.String()
}
