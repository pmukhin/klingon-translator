package http

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

const (
	canNotRead  = "can not read response from server: %s"
	canNotParse = "can not parse response from server: %s"
)

// ReadAsJson reads the whole response ignoring Content-Length
// as it might be set to 0 by some vendor APIs
//
// Then content is fed to json.Unmarshal with given pointer
func ReadAsJson(reader io.Reader, data interface{}) error {
	jsonBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return fmt.Errorf(canNotRead, err.Error())
	}
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return fmt.Errorf(canNotParse, err)
	}
	return nil
}
