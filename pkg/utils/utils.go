package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(x); err != nil {
		return err
	}
	return nil
}
