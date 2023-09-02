package main

import (
	"encoding/json"
	"fmt"
	"testing"

	//"github.com/go-avro/avro"
	//"github.com/go-avro/avro"
	//"github.com/elodina/go-avro"

	"github.com/go-avro/avro"
)

func TestMain(t *testing.T) {

	avroSchema := `{
        "type": "record",
        "name": "test",
        "fields": [
            {"name": "foo", "type": "int"},
            {"name": "bar", "type": "string"}
        ]
    }`

	sc, err := avro.ParseSchema(avroSchema)
	if err != nil {
		fmt.Println("Failed to parse Avro schema:", err)
		return
	}

	// Convert the Avro schema to a JSON Schema

	if err != nil {
		fmt.Println("Failed to convert Avro schema to JSON Schema:", err)
		return
	}

	// Pretty-print the JSON for readability
	var schemaMap map[string]interface{}
	json.Unmarshal([]byte(sc.String()), &schemaMap)
	if err != nil {
		fmt.Println("Failed to pretty-print JSON:", err)
		return
	}

	jsonstr, err := json.MarshalIndent(schemaMap, "", "")

	fmt.Println(string(jsonstr))
}

func TestSampleCobol(t *testing.T) {
	//samplecobol()
	//cobolsample()
	finalcoboltojson()
}
