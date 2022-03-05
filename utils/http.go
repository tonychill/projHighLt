package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeHttpBody(req http.Request, v interface{}) error {
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}
	return nil
}
