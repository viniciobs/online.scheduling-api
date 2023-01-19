package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	resp, _ := json.Marshal(data)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(resp)
}

func ReadJSONBody(r *http.Request, data interface{}) error {
	body, err := ReadBody(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}

func ReadBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	return body, nil
}
