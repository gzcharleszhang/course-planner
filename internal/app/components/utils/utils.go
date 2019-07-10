package utils

import (
	"encoding/json"
	"log"
)

func ToJson(v interface{}) string {
	marshalled, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("Cannot marshal %v", v)
	}
	return string(marshalled)
}
