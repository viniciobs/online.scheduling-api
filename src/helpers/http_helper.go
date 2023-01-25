package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/online.scheduling-api/src/shared"
)

func JSONResponseError(w http.ResponseWriter, status int, err error) {
	var _err any
	if err != nil {
		_err = NewError(err.Error())
	}

	JSONResponse(w, status, _err)
}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		resp, _ := json.Marshal(data)
		w.Write(resp)
	}
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

func GetErrorStatusCodeFrom(responseCode shared.Code) int {
	switch responseCode {
	case shared.DuplicatedRecord:
		return http.StatusUnprocessableEntity
	case shared.NonExistentRecord:
		return http.StatusNotFound
	case shared.ThirdPartyFail:
		return http.StatusBadGateway
	default:
		return http.StatusInternalServerError
	}
}
