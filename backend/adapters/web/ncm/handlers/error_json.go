package handlers

import "encoding/json"

func jsonError(msg string) []byte {
	err := struct {
		Message string `json:"message"`
	}{
		msg,
	}
	r, errr := json.Marshal(err)
	if errr != nil {
		return []byte(errr.Error())
	}
	return r
}
