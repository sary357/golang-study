package main

import "fmt"

func decideAllType(s []interface{}) []string {
	result := []string{}
	for _, i := range s {
		result = append(result, decideSingleType(i))
	}
	return result
}

func decideSingleType(v interface{}) string {
	switch v.(type) {
	case string:
		return "string"
	case int, int32, int64:
		return "int"
	case float32, float64:
		return "float"
	case bool:
		return "bool"
	default:
		return "unknown"
	}

}

type mytype struct {
	field1 string
}

func main() {
	sampleMyType := mytype{}
	sampleTest := []interface{}{1, 3.14, "hello", true, []interface{}{}, sampleMyType}
	typeSlice := decideAllType(sampleTest)
	for i := 0; i < len(sampleTest); i++ {
		fmt.Printf("%v is %v.\n", sampleTest[i], typeSlice[i])
	}

}
