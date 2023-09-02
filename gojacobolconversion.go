package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	"github.com/dop251/goja"
// )

// func gojaconversion() {
// 	copybook := `
// 		01 EMPLOYEE-RECORD.
// 		   05 EMPLOYEE-ID     PIC 9(5).
// 		   05 EMPLOYEE-NAME   PIC X(30).
// 		   05 EMPLOYEE-SALARY PIC 9(6)V99.
// 	`

// 	// Parse the COBOL copybook
// 	record, err := parseCopybook(copybook)
// 	if err != nil {
// 		log.Fatalf("Error parsing copybook: %v", err)
// 	}

// 	// Sample data (replace with your data)
// 	cobolData := "12345John Doe    005000" // Example COBOL data

// 	// Parse COBOL data using the copybook structure
// 	parsedData, err := parseData(record, cobolData)
// 	if err != nil {
// 		log.Fatalf("Error parsing data: %v", err)
// 	}

// 	// Convert to JSON
// 	jsonData, err := json.Marshal(parsedData)
// 	if err != nil {
// 		log.Fatalf("Error converting to JSON: %v", err)
// 	}

// 	fmt.Println(string(jsonData))
// }

// func parseCopybook(copybook string) (*goja.Object, error) {
// 	vm := goja.New()
// 	_, err := vm.RunString("var copybook = " + "`" + copybook + "`")
// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err = vm.RunString(`
// 		var parser = require('copybook-parser');
// 		var record = parser.parse(copybook);
// 	`)

// 	if err != nil {
// 		return nil, err
// 	}

// 	record := vm.Get("record").ToObject(vm)
// 	return record, nil
// }

// func parseData(record *goja.Object, data string) (map[string]interface{}, error) {
// 	vm := goja.New()
// 	_, err := vm.RunString("var data = " + "`" + data + "`")
// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err = vm.RunString(`
// 		var parser = require('copybook-parser');
// 		var parsedData = parser.parseData(record, data);
// 	`)

// 	if err != nil {
// 		return nil, err
// 	}

// 	parsedData := vm.Get("parsedData").ToObject(vm)
// 	return convertObject(parsedData), nil
// }

// func convertObject(obj *goja.Object) map[string]interface{} {
// 	result := make(map[string]interface{})

// 	return result
// }
