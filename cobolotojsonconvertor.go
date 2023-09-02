package main

import (
	"fmt"
	"strings"
)

type CobolFields struct {
	Name     string
	Type     string
	Pic      string
	Occurs   int
	Children []CobolFields
}

func convertToJson() {
	// Example parsed COBOL copybook
	copybook := CobolFields{
		Name: "CUSTOMER-RECORD",
		Type: "RECORD",
		Children: []CobolFields{
			{
				Name: "CUSTOMER-ID",
				Type: "ALPHANUMERIC",
				Pic:  "X(10)",
			},
			{
				Name: "CUSTOMER-NAME",
				Type: "ALPHANUMERIC",
				Pic:  "X(50)",
			},
			{
				Name:   "CUSTOMER-AMOUNT",
				Type:   "NUMERIC",
				Pic:    "9(10)V99",
				Occurs: 0,
			},
		},
	}

	// Convert COBOL copybook to JSON schema
	jsonSchema := ConvertToJSONSchema(copybook)

	// Print the resulting JSON schema
	fmt.Println(jsonSchema)
}

func ConvertToJSONSchema(field CobolFields) map[string]interface{} {
	jsonSchema := make(map[string]interface{})
	jsonSchema["type"] = "object"
	properties := make(map[string]interface{})

	for _, child := range field.Children {
		properties[child.Name] = mapCobolTypeToJSONTypes(child.Type, child.Pic)
		if child.Occurs > 0 {
			properties[child.Name] = mapCobolTypeToJSONTypes("array", "")
			items := make(map[string]interface{})
			items["type"] = mapCobolTypeToJSONTypes(child.Type, child.Pic)
			properties[child.Name+".items"] = items
		}
	}

	jsonSchema["properties"] = properties

	return jsonSchema
}

func mapCobolTypeToJSONTypes(cobolType string, pic string) interface{} {
	// Implement your COBOL type to JSON type mapping logic here
	switch cobolType {
	case "ALPHANUMERIC":
		return "string"
	case "NUMERIC":
		if strings.Contains(pic, "V") {
			return "number"
		}
		return "integer"
	case "RECORD":
		return "object"
	case "array":
		return "array"
	default:
		return "string" // Default to string if type is unknown
	}
}
