package main

import (
	"fmt"
	"strings"
)

func cobolsample() {
	copybook := `
	01 EMPLOYEE-RECORD.
	05 EMPLOYEE-ID     PIC 9(5).
	05 EMPLOYEE-NAME   PIC X(30).
	05 EMPLOYEE-SALARY PIC 9(6)V99.

01 DEPARTMENT-RECORD.
	05 DEPT-ID         PIC 9(3).
	05 DEPT-NAME       PIC X(20).

01 PAYROLL-RECORD.
	05 EMPLOYEE-INFO   OCCURS 10 TIMES.
		10              GROUP.
			15 EMPLOYEE-DATA REDEFINES EMPLOYEE-RECORD.
			15 DEPARTMENT-DATA REDEFINES DEPARTMENT-RECORD.
		10 EMPLOYEE-HIRED  PIC 9(8).

	`

	jsonSchema, err := MapCopybookToJSONSchema(copybook)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(jsonSchema)
}

func MapCopybookToJSONSchema(copybook string) (string, error) {
	lines := strings.Split(copybook, "\n")
	jsonSchema := "{\n"
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue
			//return "", fmt.Errorf("Invalid COBOL line: %s", line)
		}

		fieldName := parts[1]
		cobolType := parts[2] + " " + parts[3]
		jsonType, err := MapCobolTypeToJSONType(cobolType)
		if err != nil {
			return "", err
		}

		jsonSchema += fmt.Sprintf(`  "%s": {
    "type": "%s"
  },`, fieldName, jsonType)
	}

	// Remove the trailing comma
	jsonSchema = strings.TrimSuffix(jsonSchema, ",")

	jsonSchema += "\n}"

	return jsonSchema, nil
}

func MapCobolTypeToJSONType(cobolType string) (string, error) {
	// Map COBOL data types to JSON schema types
	switch cobolType {
	case "PIC X", "PIC X(N)", "PIC A", "PIC A(N)":
		return "string", nil
	case "PIC 9", "PIC 9(N)", "PIC 9(5).":
		return "integer", nil
	case "PIC 9V9", "PIC 9.9", "PIC 9.9(N)":
		return "number", nil
	default:
		return "", fmt.Errorf("Unsupported COBOL data type: %s", cobolType)
	}
}
