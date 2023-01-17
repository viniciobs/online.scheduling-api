package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	resp, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Header().Add("content-type", "application/json")
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

func GetUUID(r *http.Request) (*uuid.UUID, error) {
	id, err := uuid.Parse(mux.Vars(r)["Id"])
	if err != nil {
		return nil, err
	}

	return &id, nil
}
