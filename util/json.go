package util

import (
	j "encoding/json"
	"log"
)

// UnmarshalInput is a wrapper for json.Unmarshal()
func UnmarshalInput(input []byte) {

	var output string

	err := j.Unmarshal(input, output)

	if err != nil {
		log.Fatal(err)
	}

}
