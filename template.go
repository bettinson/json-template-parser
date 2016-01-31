package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatalf("Incorrect arg number")
	}
	replacementString := args[0]
	jsonString := args[1]
	jsonMap := stringToString(jsonString)
	replacedString := replaceString(jsonMap, replacementString)
	fmt.Println(replacedString)

}

func stringToString(j string) map[string]string {
	keyValues := map[string]string{}
	err := json.Unmarshal([]byte(j), &keyValues)
	if err != nil {
		log.Fatal(err)
	}
	return keyValues
}

func replaceString(m map[string]string, s string) string {
	appendString := make([]byte, 0)
	varAccum := make([]byte, 0)
	inVariable := false

	for i := 0; i < len(s); i++ {
		if !inVariable {
			if s[i] == '{' {
				inVariable = true
			} else {
				appendString = append(appendString, s[i])
			}
		} else {
			if s[i] == '}' {
				value, ok := m[string(varAccum)]
				if !ok {
					log.Fatalf("Cannot find key %s", string(varAccum))
				}
				appendString = append(appendString, []byte(value)...)
				inVariable = false
				varAccum = make([]byte, 0)
			} else {
				varAccum = append(varAccum, s[i])
			}
		}
	}
	return string(appendString)
}
