package bnrjson

import (
	"encoding/json"
)

func UnmarshalString(jsonString string, output interface{}) error {
	err := json.Unmarshal([]byte(jsonString), output)
	if err != nil {
		return err
	}
	return nil
}
