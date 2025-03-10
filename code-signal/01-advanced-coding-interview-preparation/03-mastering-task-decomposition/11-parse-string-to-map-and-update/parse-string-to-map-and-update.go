package main

import (
	"fmt"
	"strings"
)

func ParseStringAndUpdateValue(input, updateKey, newValue string) map[string]interface{} {
	// TODO: Your solution
	resultMap := make(map[string]interface{})

	var key string
	innerMap := make(map[string]interface{})
	inInnerMap := false
	i := 0

	for i < len(input) {
		switch input[i] {
		case '{':
			inInnerMap = true
			i++
		case '}':
			resultMap[key] = innerMap
			innerMap = make(map[string]interface{})
			inInnerMap = false
			i++
			if i < len(input) && input[i] == ',' {
				i++
			}
		default:
			if !inInnerMap {
				// TODO: we need the positions for "=" and ","
				equalPos := strings.Index(input[i:], "=")
				if equalPos == -1 {
					break // or handle the error
				}
				equalPos += i
				commaPos := strings.Index(input[equalPos:], ",")
				if commaPos != -1 {
					commaPos += equalPos // Adjust commaPos to be relative to the start of the string
				} else {
					commaPos = len(input) // If no comma is found, set it to the end of the string
				}

				key = input[i:equalPos]
				value := input[equalPos+1 : commaPos]

				if strings.Contains(value, "{") {
					key = input[i:equalPos]
					i = equalPos + 1
				} else {
					resultMap[key] = value
					i = commaPos + 1
				}
			}

			if inInnerMap {
				equalPos := strings.Index(input[i:], "=") + i
				commaPos := strings.Index(input[equalPos:], ",") + equalPos
				bracePos := strings.Index(input[equalPos:], "}") + equalPos

				endPos := len(input)
				if bracePos >= equalPos && (commaPos < 0 || bracePos < commaPos) {
					endPos = bracePos
				} else if commaPos >= equalPos {
					endPos = commaPos
				}

				innerKey := input[i:equalPos]
				innerValue := input[equalPos+1 : endPos]
				innerMap[innerKey] = innerValue

				i = endPos
				if i < len(input) && input[i] == ',' {
					i++
				}
			}
		}
	}

	updateMap(resultMap, updateKey, newValue)

	return resultMap
}

func updateMap(m map[string]interface{}, key, value string) {
	if val, exists := m[key]; exists {
		if _, ok := val.(string); ok {
			m[key] = value
		} else if innerMap, ok := val.(map[string]interface{}); ok {
			if _, exists := innerMap[key]; exists {
				innerMap[key] = value
			}
		}
		return
	}

	for _, v := range m {
		if innerMap, ok := v.(map[string]interface{}); ok {
			if _, exists := innerMap[key]; exists {
				innerMap[key] = value
				return
			}
		}
	}
}

// This function is provided to assist with result visualization
func PrettyPrint(dict map[string]interface{}, indent string) string {
	result := "{\n"
	for k, v := range dict {
		switch val := v.(type) {
		case map[string]interface{}:
			result += fmt.Sprintf("%s  %s: %s", indent, k, PrettyPrint(val, indent+"  "))
		default:
			result += fmt.Sprintf("%s  %s: '%s',\n", indent, k, v)
		}
	}
	result += fmt.Sprintf("%s}\n", indent)
	return result
}

func main() {
	input := "A1=B1,C1={D1=E1,F1=G1},I1=J1"
	updateKey := "F1"
	newValue := "newValue"

	updatedMap := ParseStringAndUpdateValue(input, updateKey, newValue)
	fmt.Println(PrettyPrint(updatedMap, ""))
}
