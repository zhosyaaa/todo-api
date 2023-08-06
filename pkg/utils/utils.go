package utils

import (
	"encoding/json"
	"net/http"
)

//func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
//	response, err := json.Marshal(payload)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write([]byte(err.Error()))
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(status)
//	w.Write([]byte(response))
//}
//
//func respondError(w http.ResponseWriter, code int, message string) {
//	respondJSON(w, code, map[string]string{"error": message})
//}

func ParseBody(r *http.Request, x interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(x); err != nil {
		return err
	}
	return nil
}
