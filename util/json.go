package util

import "encoding/json"

func MustJsonString(v interface{}) string {
	jsonBytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
