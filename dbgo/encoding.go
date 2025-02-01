package dbgo

import "encoding/json"

type DataEncoder interface {
	Encode(Map) ([]byte, error)
}

type DataDecoder interface {
	Decode([]byte, any) error
}

type JSONEncoder struct{}

// Go struct to JSON
func (JSONEncoder) Encode(data Map) ([]byte, error) {
	return json.Marshal(data)
}

type JSONDecoder struct{}

// JSON( in form of byte slice) to Go struct
func (JSONDecoder) Decode(b []byte, v any) error {
	return json.Unmarshal(b, &v)
}
