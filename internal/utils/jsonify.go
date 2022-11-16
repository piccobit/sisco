package utils

import (
	"encoding/json"
	"log"
)

func JSONify(in any, makePretty bool) string {
	var err error
	var jsonBlob []byte

	if makePretty {
		jsonBlob, err = json.MarshalIndent(in, "", "  ")
	} else {
		jsonBlob, err = json.Marshal(in)
	}
	if err != nil {
		log.Fatalf("error in JSONify: %v", err)
	}

	return string(jsonBlob)
}
