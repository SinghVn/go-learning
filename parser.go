package main

// import (
// 	"fmt"
// )

// type Field interface {
// 	add(Field)
// }

// type FieldWrapper struct {
// 	children []Field
// }

// func (fw *FieldWrapper) add(f Field) {
// 	fw.children = append(fw.children, f)
// }

// type FieldArray struct {
// 	Type     string
// 	Name     string
// 	children []Field
// }

// func (fa *FieldArray) add(f Field) {
// 	fa.children = append(fa.children, f)
// }

// func parse() {
// 	bookStructure := map[string]interface{}{
// 		"MAIN-RECORD": map[string]interface{}{
// 			"FIELD1": map[string]interface{}{
// 				"type": "ALPHANUMERIC",
// 				"name": "Field1",
// 				"size": 10,
// 			},
// 			"FIELD2": map[string]interface{}{
// 				"type": "NUMERIC",
// 				"name": "Field2",
// 				"size": 5,
// 			},
// 			"ARRAY_FIELD": []interface{}{
// 				map[string]interface{}{
// 					"type": "ALPHANUMERIC",
// 					"name": "ArrayField1",
// 					"size": 20,
// 				},
// 				map[string]interface{}{
// 					"type": "NUMERIC",
// 					"name": "ArrayField2",
// 					"size": 3,
// 				},
// 			},
// 		},
// 	}

// 	parserType := "FLAT_ASCII"
// 	mainNode := &FieldWrapper{}
// 	parser := create(mainNode, bookStructure, parserType)

// 	result := build(parser)
// 	fmt.Println(result)
// }

// func build(parser Field) Field {
// 	return parser
// }

// func create(parent Field, bookStructure map[string]interface{}, parserType string) Field {
// 	for key, value := range bookStructure {
// 		if subMap, isMap := value.(map[string]interface{}); isMap {
// 			f := chooseInstance(subMap, parserType)
// 			parent.add(f)
// 		} else if subSlice, isSlice := value.([]interface{}); isSlice {
// 			arr := &FieldArray{
// 				Type: "ARRAY",
// 				Name: key,
// 			}
// 			parent.add(arr)
// 			for _, item := range subSlice {
// 				g := &FieldWrapper{}
// 				arr.add(g)
// 				create(g, item.(map[string]interface{}), parserType)
// 			}
// 		} else {
// 			panic(fmt.Sprintf("Invalid type %T", value))
// 		}
// 	}
// 	return parent
// }

// func chooseInstance(value map[string]interface{}, parserType string) Field {
// 	var f Field
// 	typ := value["type"].(string)
// 	name := value["name"].(string)
// 	switch typ {
// 	case "ALPHANUMERIC", "ALPHABETIC":
// 		if parserType == "FLAT_ASCII" {
// 			f = &FieldAlphanumericAscii{
// 				Type: typ,
// 				Name: name,
// 				Size: value["size"].(int),
// 			}
// 		} else if parserType == "BINARY_EBCDIC" {
// 			f = &FieldAlphanumericEbcdic{
// 				Type: typ,
// 				Name: name,
// 				Size: value["size"].(int),
// 			}
// 		} else {
// 			panic(fmt.Sprintf("Invalid parserType: %s", parserType))
// 		}
// 	case "NUMERIC":
// 		if parserType == "FLAT_ASCII" {
// 			f = &FieldNumericAscii{
// 				Type:     typ,
// 				Name:     name,
// 				Size:     value["size"].(int),
// 				Decimals: value["decimals"].(int),
// 			}
// 		} else if parserType == "BINARY_EBCDIC" {
// 			f = &FieldNumericEbcdic{
// 				Type:     typ,
// 				Name:     name,
// 				Size:     value["size"].(int),
// 				Decimals: value["decimals"].(int),
// 			}
// 		} else {
// 			panic(fmt.Sprintf("Invalid parserType: %s", parserType))
// 		}
// 	case "NUMERIC_COMP3":
// 		f = &FieldNumericComp3{
// 			Type:     typ,
// 			Name:     name,
// 			Size:     value["size"].(int),
// 			Decimals: value["decimals"].(int),
// 		}
// 	case "NUMERIC_BINARY":
// 		f = &FieldNumericBinary{
// 			Type:     typ,
// 			Name:     name,
// 			Size:     value["size"].(int),
// 			Decimals: value["decimals"].(int),
// 		}
// 	case "NUMERIC_MASKED":
// 		if parserType != "FLAT_ASCII" {
// 			panic(fmt.Sprintf("Invalid parserType: %s, can only be used with ParseType.FLAT_ASCII", parserType))
// 		}
// 		f = &FieldNumericMaskedAscii{
// 			Type:     typ,
// 			Name:     name,
// 			Size:     value["size"].(int),
// 			Decimals: value["decimals"].(int),
// 		}
// 	default:
// 		panic(fmt.Sprintf("Invalid format: %s", typ))
// 	}
// 	return f
// }

// type FieldAlphanumericAscii struct {
// 	Type string
// 	Name string
// 	Size int
// }

// type FieldAlphanumericEbcdic struct {
// 	Type string
// 	Name string
// 	Size int
// }

// type FieldNumericAscii struct {
// 	Type     string
// 	Name     string
// 	Size     int
// 	Decimals int
// }

// type FieldNumericBinary struct {
// 	Type     string
// 	Name     string
// 	Size     int
// 	Decimals int
// }

// type FieldNumericComp3 struct {
// 	Type     string
// 	Name     string
// 	Size     int
// 	Decimals int
// }

// type FieldNumericEbcdic struct {
// 	Type     string
// 	Name     string
// 	Size     int
// 	Decimals int
// }

// type FieldNumericMaskedAscii struct {
// 	Type     string
// 	Name     string
// 	Size     int
// 	Decimals int
// }
