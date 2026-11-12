package ohauthcat

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// objectFields refuses ambiguous duplicate keys, including nested objects.
func objectFields(data []byte) (map[string]json.RawMessage, error) {
	if err := unambiguousJSON(data); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	opening, err := decoder.Token()
	if err != nil || opening != json.Delim('{') {
		return nil, errors.New("invalid JSON object")
	}
	fields := make(map[string]json.RawMessage)
	for decoder.More() {
		token, tokenErr := decoder.Token()
		if tokenErr != nil {
			return nil, errors.New("invalid JSON object key")
		}
		key, ok := token.(string)
		if !ok {
			return nil, errors.New("invalid JSON object key")
		}
		if _, exists := fields[key]; exists {
			return nil, errors.New("duplicate JSON object key")
		}
		var value json.RawMessage
		if decodeErr := decoder.Decode(&value); decodeErr != nil {
			return nil, errors.New("invalid JSON object value")
		}
		fields[key] = value
	}
	closing, err := decoder.Token()
	if err != nil || closing != json.Delim('}') {
		return nil, errors.New("invalid JSON object")
	}
	if _, err := decoder.Token(); err != io.EOF {
		return nil, errors.New("trailing JSON value")
	}
	return fields, nil
}

func unambiguousJSON(data []byte) error {
	if len(data) > 1024*1024 || !json.Valid(data) {
		return errors.New("invalid JSON value")
	}
	type frame struct {
		keys map[string]bool
		key  bool
	}
	var stack []frame
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.New("invalid JSON value")
		}
		if delimiter, ok := token.(json.Delim); ok {
			if delimiter == '}' || delimiter == ']' {
				stack = stack[:len(stack)-1]
				continue
			}
			if len(stack) == 64 {
				return errors.New("JSON nesting limit exceeded")
			}
			if len(stack) > 0 {
				stack[len(stack)-1].key = true
			}
			next := frame{}
			if delimiter == '{' {
				next.keys = make(map[string]bool)
				next.key = true
			}
			stack = append(stack, next)
			continue
		}
		if len(stack) == 0 {
			continue
		}
		current := &stack[len(stack)-1]
		if current.keys != nil && current.key {
			key, ok := token.(string)
			if !ok || current.keys[key] {
				return errors.New("duplicate or invalid JSON object key")
			}
			current.keys[key] = true
			current.key = false
		} else {
			current.key = true
		}
	}
	return nil
}

func additionalFields(fields map[string]json.RawMessage, known []string) (map[string]JSONValue, error) {
	result := make(map[string]JSONValue)
	for name, data := range fields {
		isKnown := false
		for _, key := range known {
			if key == name {
				isKnown = true
				break
			}
		}
		if isKnown {
			continue
		}
		value, err := ParseJSONValue(data)
		if err != nil {
			return nil, err
		}
		result[name] = value
	}
	return result, nil
}

func mergeAdditional(data []byte, additional map[string]JSONValue, known []string) ([]byte, error) {
	fields, err := objectFields(data)
	if err != nil {
		return nil, err
	}
	for name, value := range additional {
		for _, key := range known {
			if key == name {
				return nil, errors.New("additional field collides with typed field")
			}
		}
		encoded, err := value.MarshalJSON()
		if err != nil {
			return nil, err
		}
		fields[name] = encoded
	}
	return json.Marshal(fields)
}
