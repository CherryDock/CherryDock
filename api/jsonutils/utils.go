package jsonutils

import (
	"encoding/json"
	"log"
)

func FormatToJson(data interface{}) []byte {
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Fail to encode to json", err)
	}
	return dataJson
}

func ParseJson(data []byte) map[string]interface{} {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		log.Fatal("Fail to parse json", err)
	}
	return jsonData
}
