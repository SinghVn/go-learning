package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func samplecobol() {
	cobolCopybook := `
		01 CUSTOMER-RECORD.
		   05 CUSTOMER-ID     PIC X(10).
		   05 CUSTOMER-NAME   PIC X(50).
		   05 CUSTOMER-AMOUNT PIC 9(10)V99.
	`

	jsonSchema := ConvertCopybookToJSONSchema(cobolCopybook)
	fmt.Println(jsonSchema)
}

func ConvertCopybookToJSONSchema(cobolCopybook string) string {
	lines := strings.Split(cobolCopybook, "\n")
	jsonSchema := make(map[string]interface{})
	properties := make(map[string]interface{})
	jsonSchema["type"] = "object"
	jsonSchema["properties"] = properties

	stack := []map[string]interface{}{properties}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := strings.Fields(line)
		level, fieldName := getLevelAndFieldName(parts)
		if level == 1 {
			properties[fieldName] = map[string]interface{}{"type": "object", "properties": make(map[string]interface{})}
			stack = append(stack, properties[fieldName].(map[string]interface{})["properties"].(map[string]interface{}))
		} else {
			fieldType, _ := getFieldTypeAndSize(parts)
			if fieldType != "" {
				currentProperties := stack[len(stack)-1]
				currentProperties[fieldName] = map[string]interface{}{"type": mapFieldTypeToJSONType(fieldType), "description": fmt.Sprintf("COBOL PIC: %s", parts[len(parts)-1])}
				if fieldType == "array" {
					currentProperties[fieldName].(map[string]interface{})["items"] = map[string]interface{}{"type": mapFieldTypeToJSONType(fieldType), "description": fmt.Sprintf("COBOL PIC: %s", parts[len(parts)-1])}
				}
			}
		}
	}

	jsonSchemaStr, _ := json.MarshalIndent(jsonSchema, "", "  ")
	return string(jsonSchemaStr)
}

func getLevelAndFieldName(parts []string) (int, string) {
	if len(parts) >= 2 {
		level := 0
		fieldName := ""
		for _, part := range parts {
			if part == "01" {
				level++
			} else if level == 1 && part != "RECORD." {
				fieldName = part
			}
		}
		return level, fieldName
	}
	return 0, ""
}

func getFieldTypeAndSize(parts []string) (string, string) {
	if len(parts) >= 3 {
		fieldType := parts[1]
		fieldSize := parts[2]
		return fieldType, fieldSize
	}
	return "", ""
}

func mapFieldTypeToJSONType(fieldType string) string {
	switch fieldType {
	case "X", "X.":
		return "string"
	case "9", "9.":
		return "number"
	default:
		return "array"
	}
}
