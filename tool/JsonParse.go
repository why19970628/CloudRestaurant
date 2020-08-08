package tool

import (
	"encoding/json"
	"io"
)

/**
	提供参数解析的一个工具类
**/
type JsonParse struct {
}

func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}
