package helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	resp, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}

func GetUUID(r *http.Request) (*uuid.UUID, error) {
	id, err := uuid.Parse(mux.Vars(r)["Id"])
	if err != nil {
		return nil, err
	}

	return &id, nil
}
