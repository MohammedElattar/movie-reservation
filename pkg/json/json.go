// Package json
package json

import "github.com/bytedance/sonic"

func Marshal(v any) ([]byte, error) {
	return sonic.Marshal(v)
}

func MarshalString(v any) (string, error) {
	return sonic.MarshalString(v)
}

func Unmarshal(buf []byte, val any) error {
	return sonic.Unmarshal(buf, val)
}

func UnmarshalString(buf string, val any) error {
	return sonic.UnmarshalString(buf, val)
}
