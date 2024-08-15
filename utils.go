package gchat_wh

import (
	"bytes"
	"encoding/json"
	"io"
)

func jsonStream[T any](data T) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
