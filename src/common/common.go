package common

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

// ReadJSONFile : read json file from path string
func ReadJSONFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// Jsonify : stringtify
func Jsonify(m interface{}) string {
	b, _ := json.MarshalIndent(m, "", " ")

	return string(b)
}

// StringToBoolean : convert string to bool
func StringToBoolean(s string) (bool, error) {
	v := strings.ToLower(s)
	if v == "true" {
		return true, nil
	}
	if v == "false" {
		return false, nil
	}

	return false, errors.New("invalid boolean string")
}
