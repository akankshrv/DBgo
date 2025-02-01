package dbgo

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	encoder := JSONEncoder{}

	inputStruct := map[string]interface{}{
		"name":  "Akanksh",
		"age":   22,
		"email": "akanksh@example.com",
	}
	expectedJSON := `{"age":22,"email":"akanksh@example.com","name":"Akanksh"}`

	response, err := encoder.Encode(inputStruct)
	if err != nil {
		t.Fatalf("Encode isnt working champ: %v\n", err)

	}

	if !bytes.Equal(response, []byte(expectedJSON)) {
		t.Errorf("error, doesnt match ")
	}

}

func TestDecode(t *testing.T) {
	decoder := JSONDecoder{}

	inputJSON := `{"age":22,"email":"akanksh@example.com","name":"Akanksh"}`

	expectedStruct := Map{
		"name":  "Akanksh",
		"age":   float64(22),
		"email": "akanksh@example.com",
	}

	var response Map

	err := decoder.Decode([]byte(inputJSON), &response)
	if err != nil {
		t.Fatalf("error in decoding: %v", err)
	}

	for k, v := range response {
		if expectedStruct[k] != v {
			t.Errorf("error")
		}
	}
}
