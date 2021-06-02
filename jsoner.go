package jsoner

import (
	"encoding/json"
)

type buf []byte

func (b *buf) Write(p []byte) (n int, err error) {
	*b = append(*b, p...)
	return len(p), nil
}

func ToJSON(input map[string]interface{}) (output string, err error) {
	b := &buf{}
	enc := json.NewEncoder(b)
	if err = enc.Encode(&input); err != nil {
		return
	}
	output = string(*b)
	return
}
