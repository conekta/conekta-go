package conekta

import "encoding/json"

// ParamsConverter is an interface to handle Conekta params conversions in requestor
type ParamsConverter interface {
	Bytes() []byte
}

func paramsToBytes(v interface{}) []byte {
	r, err := json.Marshal(v)
	if err != nil {
		return []byte{}
	}
	return r
}
