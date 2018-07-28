package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetJSONParams(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(data)
	switch {
	case err == io.EOF:
		return nil
	case err != nil:
		return fmt.Errorf("[GetJSONParams] failed to decode : {%+v}", err)
	}

	return nil
}
