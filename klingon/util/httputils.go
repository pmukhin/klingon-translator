package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	canNotRead  = "can not read response from server: %s"
	canNotParse = "can not parse response from server: %s"
)

func ReadResponse(resp *http.Response, data interface{}) error {
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf(canNotRead, err.Error())
	}
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return fmt.Errorf(canNotParse, err)
	}
	return nil
}
