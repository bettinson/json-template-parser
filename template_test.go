package main

import "testing"

func TestStringToString(t *testing.T) {
	jsonString := `{"Matt": "Cool guy", "Jake": "Pretty cool"}`
	testJson := map[string]string{
		"Matt": "Cool guy",
		"Jake": "Pretty cool",
	}
	result := stringToString(jsonString)
	for k, v := range testJson {
		val, ok := result[k]
		if !ok || (v != val) {
			t.Fatal("Strngs are not equal")
		}
	}
}

func TestReplaceString(t *testing.T) {
	testJson := map[string]string{
		"Matt": "Cool guy",
		"Jake": "Pretty cool",
	}
	given := "I think that Matt is a {Matt} and Jake is only {Jake}"
	expected := "I think that Matt is a Cool guy and Jake is only Pretty cool"

	test := replaceString(testJson, given)
	if expected != test {
		t.Fatalf("Expected: %s \n Got: %s", expected, test)
	}

}
