package utils

import (
	"encoding/json"
	"log"
)

type M map[string]interface{}

func ToJson(v interface{}) string {
	marshalled, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("Cannot marshal %v", v)
	}
	return string(marshalled)
}

func ToRawJson(v interface{}) []byte {
	marshalled, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("Cannot marshal %v", v)
	}
	return marshalled
}
