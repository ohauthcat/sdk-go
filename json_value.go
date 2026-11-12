package ohauthcat

import (
	"bytes"
	"errors"
)

// JSONValue stores a validated arbitrary JSON value without erasing typed models.
// Its zero value represents JSON null. Input and output bytes are copied.
type JSONValue struct {
	data []byte
}

// ParseJSONValue validates one JSON value and takes an independent copy.
func ParseJSONValue(data []byte) (JSONValue, error) {
	if err := unambiguousJSON(data); err != nil {
		return JSONValue{}, errors.New("invalid JSON value")
	}
	return JSONValue{data: bytes.Clone(data)}, nil
}

// MarshalJSON returns a copy of the validated representation.
func (v JSONValue) MarshalJSON() ([]byte, error) {
	if len(v.data) == 0 {
		return []byte("null"), nil
	}
	return bytes.Clone(v.data), nil
}

// UnmarshalJSON replaces the value only after validation succeeds.
func (v *JSONValue) UnmarshalJSON(data []byte) error {
	if v == nil {
		return errors.New("nil JSON value receiver")
	}
	parsed, err := ParseJSONValue(data)
	if err != nil {
		return err
	}
	*v = parsed
	return nil
}
