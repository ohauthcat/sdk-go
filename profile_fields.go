package ohauthcat

import (
	"bytes"
	"encoding/json"
	"errors"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

// normalizeProfileFields only follows pinned schema references, never field names alone.
func normalizeProfileFields(data []byte, definition schema, definitions map[string]schema) ([]byte, error) {
	if err := unambiguousJSON(data); err != nil {
		return nil, err
	}
	return normalizeProfileValue(data, definition, definitions, 0)
}

func normalizeProfileValue(data []byte, definition schema, definitions map[string]schema, depth int) ([]byte, error) {
	if depth > 128 {
		return nil, errors.New("invalid response depth")
	}
	if definition.Ref != "" {
		name := strings.TrimPrefix(definition.Ref, "#/components/schemas/")
		next, ok := definitions[name]
		if !ok {
			return data, nil
		}
		if name == "CustomAttributeValues" && !bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
			var values CustomAttributeValues
			if err := json.Unmarshal(data, &values); err != nil {
				return nil, err
			}
			return json.Marshal(values)
		}
		if name == "Event" {
			var err error
			data, err = normalizeDirectoryEvent(data)
			if err != nil {
				return nil, err
			}
		}
		return normalizeProfileValue(data, next, definitions, depth+1)
	}
	for _, next := range definition.AllOf {
		var err error
		data, err = normalizeProfileValue(data, next, definitions, depth+1)
		if err != nil {
			return nil, err
		}
	}
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return nil, errors.New("invalid response value")
	}
	if trimmed[0] == '{' && len(definition.Properties) != 0 {
		fields, err := objectFields(data)
		if err != nil {
			return nil, err
		}
		for name, child := range definition.Properties {
			raw, exists := fields[name]
			if !exists {
				continue
			}
			fields[name], err = normalizeProfileValue(raw, child, definitions, depth+1)
			if err != nil {
				return nil, err
			}
		}
		return json.Marshal(fields)
	}
	if trimmed[0] == '[' && definition.Items != nil {
		var items []json.RawMessage
		if err := json.Unmarshal(data, &items); err != nil {
			return nil, err
		}
		for i, raw := range items {
			var err error
			items[i], err = normalizeProfileValue(raw, *definition.Items, definitions, depth+1)
			if err != nil {
				return nil, err
			}
		}
		return json.Marshal(items)
	}
	return data, nil
}

func normalizeDirectoryEvent(data []byte) ([]byte, error) {
	fields, err := objectFields(data)
	if err != nil {
		return nil, err
	}
	var kind string
	if json.Unmarshal(fields["type"], &kind) != nil {
		return data, nil
	}
	switch kind {
	case "dsync.user.created", "dsync.user.updated", "dsync.user.deleted":
	default:
		return data, nil
	}
	raw, exists := fields["data"]
	if !exists {
		return data, nil
	}
	payload, err := objectFields(raw)
	if err != nil {
		return data, nil
	}
	var object string
	if json.Unmarshal(payload["object"], &object) != nil || object != "directory_user" {
		return data, nil
	}
	attributes, exists := payload["custom_attributes"]
	if !exists || bytes.Equal(bytes.TrimSpace(attributes), []byte("null")) {
		return data, nil
	}
	var values CustomAttributeValues
	if decodeErr := json.Unmarshal(attributes, &values); decodeErr != nil {
		return nil, decodeErr
	}
	payload["custom_attributes"], err = json.Marshal(values)
	if err != nil {
		return nil, err
	}
	fields["data"], err = json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return json.Marshal(fields)
}

// validCustomAttributeJSON applies D96 byte and Unicode bounds before decoding can replace malformed text.
func validCustomAttributeJSON(data []byte) bool {
	if !utf8.Valid(data) || len(data) > 16*1024 {
		return false
	}
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return false
	}
	if data[0] == '[' {
		var items []json.RawMessage
		if json.Unmarshal(data, &items) != nil {
			return false
		}
		for _, item := range items {
			if !validCustomAttributeJSON(item) {
				return false
			}
		}
		return true
	}
	if data[0] != '"' {
		return true
	}
	for i := 1; i < len(data)-1; i++ {
		if data[i] != '\\' {
			continue
		}
		i++
		if i >= len(data)-1 {
			return false
		}
		if data[i] != 'u' {
			continue
		}
		if i+4 >= len(data) {
			return false
		}
		value, err := strconv.ParseUint(string(data[i+1:i+5]), 16, 16)
		if err != nil {
			return false
		}
		i += 4
		if value >= 0xDC00 && value <= 0xDFFF {
			return false
		}
		if value >= 0xD800 && value <= 0xDBFF {
			if i+6 >= len(data) || data[i+1] != '\\' || data[i+2] != 'u' {
				return false
			}
			low, err := strconv.ParseUint(string(data[i+3:i+7]), 16, 16)
			if err != nil || low < 0xDC00 || low > 0xDFFF {
				return false
			}
			i += 6
		}
	}
	var value string
	return json.Unmarshal(data, &value) == nil && value != "" && len(value) <= 1024
}

var customAttributeKey = regexp.MustCompile(`^[a-z][a-z0-9_]{0,63}$`)

func trimCustomAttributeValues(values CustomAttributeValues) {
	keys := make([]string, 0, len(values))
	for key := range values {
		if !customAttributeKey.MatchString(key) {
			delete(values, key)
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for len(keys) > 50 {
		delete(values, keys[len(keys)-1])
		keys = keys[:len(keys)-1]
	}
	for len(keys) > 0 {
		encoded, err := json.Marshal(values)
		if err == nil && len(encoded) <= 32*1024 {
			return
		}
		delete(values, keys[len(keys)-1])
		keys = keys[:len(keys)-1]
	}
}
