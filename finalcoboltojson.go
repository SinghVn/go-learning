package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func finalcoboltojson() {
	copybook := `
		01 EMPLOYEE-RECORD.
		   05 EMPLOYEE-ID     PIC 9(5).
		   05 EMPLOYEE-NAME   PIC X(30).
		   05 EMPLOYEE-SALARY PIC 9(6)V99.
		   05 EMPLOYEE-ADDRESS.
		      10 STREET      PIC X(50).
		      10 CITY        PIC X(30).
		      10 STATE       PIC X(2).
		      10 ZIP-CODE    PIC 9(5).
		01 DEPARTMENT-RECORD.
		   05 DEPT-ID         PIC 9(3).
		   05 DEPT-NAME       PIC X(20).
		   05 EMPLOYEES OCCURS 10 TIMES.
		      10 EMPLOYEE-INFO.
		         15 EMPLOYEE-ID     PIC 9(5).
		         15 EMPLOYEE-NAME   PIC X(30).
	`

	jsonSchema, err := ConvertCopybookToJSONSchemas(copybook)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(jsonSchema)
}

func ConvertCopybookToJSONSchemas(copybook string) (string, error) {
	lines := strings.Split(copybook, "\n")
	var jsonSchemaMap map[string]interface{}
	jsonSchemaMap = make(map[string]interface{})
	currentSchema := jsonSchemaMap
	var stack []map[string]interface{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Detect the hierarchy level based on the indentation
		level := strings.Count(line, " ")

		// Remove any level-specific indentation
		line = strings.TrimLeft(line, " ")

		// Check if it's a field declaration or a group declaration
		if strings.HasPrefix(line, "01 ") {
			// Field declaration
			parts := strings.Fields(line)
			fieldName := parts[1]
			//cobolType := parts[2]
			// jsonType, err := MapCobolTypeToJSONTypes(cobolType)
			// if err != nil {
			// 	return "", err
			// }

			currentSchema[fieldName] = map[string]interface{}{
				"type": "object",
			}
		} else if strings.HasPrefix(line, "05 ") {
			// Field within a group
			parts := strings.Fields(line)
			if len(parts) < 3 {
				fieldName := parts[1]
				currentSchema[fieldName] = map[string]interface{}{
					"type": map[string]interface{}{},
				}
			} else {
				fieldName := parts[1]
				cobolType := parts[2] + " " + parts[3]
				jsonType, err := MapCobolTypeToJSONTypes(cobolType)
				if err != nil {
					return "", err
				}

				currentSchema[fieldName] = map[string]interface{}{
					"type": jsonType,
				}
			}
		} else if strings.HasPrefix(line, "10 ") {
			// Nested field within a group
			var jsonType = ""
			parts := strings.Fields(line)
			fieldName := parts[1]
			if len(parts) < 3 {

				currentSchema[fieldName] = map[string]interface{}{
					"type": map[string]interface{}{},
				}
			} else {

				cobolType := parts[2] + " " + parts[3]
				_, err := MapCobolTypeToJSONTypes(cobolType)
				if err != nil {
					return "", err
				}
			}

			// Check if the currentSchema is a group or an array
			if currentSchema["_array"] != nil {
				// Handle array elements
				arraySchema := currentSchema["_array"].(map[string]interface{})
				arraySchema[fieldName] = map[string]interface{}{
					"type": jsonType,
				}
			} else {
				// Handle group elements
				groupSchema := currentSchema[fieldName].(map[string]interface{})
				groupSchema["type"] = "object"
				groupSchema["properties"] = map[string]interface{}{
					fieldName: map[string]interface{}{
						"type": jsonType,
					},
				}
			}
		} else {
			// Assume it's a group declaration
			parts := strings.Fields(line)
			groupName := parts[1]
			groupSchema := map[string]interface{}{
				"type":       "object",
				"properties": make(map[string]interface{}),
			}

			// Check if the currentSchema is an array
			if currentSchema["_array"] != nil {
				arraySchema := currentSchema["_array"].(map[string]interface{})
				arraySchema[groupName] = groupSchema
				stack = append(stack, arraySchema)
			} else {
				// Add the group to the current schema
				currentSchema[groupName] = groupSchema
				stack = append(stack, currentSchema)
			}

			// Set the current schema to the group schema
			currentSchema = groupSchema
		}

		// Pop the schema stack if the indentation level decreases
		for len(stack) > level {
			currentSchema = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	// Convert the JSON schema map to a JSON string
	jsonSchemaBytes, err := json.MarshalIndent(jsonSchemaMap, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonSchemaBytes), nil
}

func MapCobolTypeToJSONTypes(cobolType string) (string, error) {
	// Map COBOL data types to JSON schema types
	switch value := cobolType; {
	case strings.Contains(value, "PIC X") || strings.Contains(value, "PIC A"):
		return "string", nil
	case strings.Contains(value, "PIC 9"):

		return "number", nil
	default:
		return "", fmt.Errorf("Unsupported COBOL data type: %s", cobolType)
	}
}
