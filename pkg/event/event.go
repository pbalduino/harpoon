package event

import (
	"encoding/json"
	"log"
)

/*
string -> json -> filter -> connect -> send string
*/

// Process mimimi
func Process(byt []byte) {
	toJSON(byt)
}

func toJSON(byt []byte) ([]byte, map[string]interface{}) {
	var j map[string]interface{}

	if err := json.Unmarshal(byt, &j); err != nil {
		panic(err)
	}

	log.Printf("%s:%s@%s\n", j["Type"], j["Action"], j["id"])

	return byt, j
}
