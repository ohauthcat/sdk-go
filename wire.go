package ohauthcat

import (
	"bytes"
	"encoding/json"
	"math"
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type schema struct {
	Ref           string            `json:"$ref"`
	Type          json.RawMessage   `json:"type"`
	Enum          []JSONValue       `json:"enum"`
	Const         json.RawMessage   `json:"const"`
	Properties    map[string]schema `json:"properties"`
	Required      []string          `json:"required"`
	Additional    json.RawMessage   `json:"additionalProperties"`
	Items         *schema           `json:"items"`
	OneOf         []schema          `json:"oneOf"`
	AnyOf         []schema          `json:"anyOf"`
	AllOf         []schema          `json:"allOf"`
	MinLength     *int              `json:"minLength"`
	MaxLength     *int              `json:"maxLength"`
	MinItems      *int              `json:"minItems"`
	MaxItems      *int              `json:"maxItems"`
	UniqueItems   bool              `json:"uniqueItems"`
	RejectNull    bool              `json:"x-ohauthcat-reject-null"`
	MinProperties *int              `json:"minProperties"`
	MaxProperties *int              `json:"maxProperties"`
	Minimum       *float64          `json:"minimum"`
	Maximum       *float64          `json:"maximum"`
	Pattern       string            `json:"pattern"`
	Format        string            `json:"format"`
	Discriminator struct {
		PropertyName string `json:"propertyName"`
	} `json:"discriminator"`
}

// wireField names the property being matched, so a response can tell an open
// single-value enum from one that identifies a type.
type wireField struct {
	name          string
	discriminator string
}

func matchesJSON(data []byte, definition schema, definitions map[string]schema, response bool) bool {
	if err := unambiguousJSON(data); err != nil {
		return false
	}
	return matchesValue(bytes.TrimSpace(data), definition, definitions, response, 0, wireField{})
}

func matchesValue(data []byte, definition schema, definitions map[string]schema, response bool, depth int, field wireField) bool {
	if depth > 128 {
		return false
	}
	if len(data) == 0 || (definition.RejectNull && bytes.Equal(data, []byte("null"))) {
		return false
	}
	discriminator := field.discriminator
	if definition.Discriminator.PropertyName != "" {
		discriminator = definition.Discriminator.PropertyName
	}
	// child matches a nested value that is not itself a named property.
	child := func(value []byte, next schema) bool {
		return matchesValue(bytes.TrimSpace(value), next, definitions, response, depth+1, wireField{name: field.name, discriminator: discriminator})
	}
	property := func(name string, value []byte, next schema) bool {
		return matchesValue(bytes.TrimSpace(value), next, definitions, response, depth+1, wireField{name: name, discriminator: discriminator})
	}
	if definition.Ref != "" {
		const prefix = "#/components/schemas/"
		if !strings.HasPrefix(definition.Ref, prefix) {
			return false
		}
		next, exists := definitions[strings.TrimPrefix(definition.Ref, prefix)]
		if !exists || !child(data, next) {
			return false
		}
		// OpenAPI 3.1 reference siblings remain conjunctive constraints.
	}
	for _, next := range definition.AllOf {
		if !child(data, next) {
			return false
		}
	}
	if len(definition.OneOf) > 0 {
		matches := 0
		for _, next := range definition.OneOf {
			if child(data, next) {
				matches++
			}
		}
		if matches != 1 {
			return false
		}
	}
	if len(definition.AnyOf) > 0 {
		matched := false
		for _, next := range definition.AnyOf {
			if child(data, next) {
				matched = true
			}
		}
		if !matched {
			return false
		}
	}
	if len(definition.Const) > 0 && !equalJSON(data, definition.Const) {
		return false
	}
	if len(definition.Enum) > 0 && (!response || (len(definition.Enum) == 1 && !openEnum(definition.Enum, field, discriminator))) {
		found := false
		for _, value := range definition.Enum {
			encoded, _ := value.MarshalJSON()
			if equalJSON(data, encoded) {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	var kinds []string
	if len(definition.Type) > 0 {
		var single string
		if err := json.Unmarshal(definition.Type, &single); err == nil {
			kinds = []string{single}
		} else if err := json.Unmarshal(definition.Type, &kinds); err != nil {
			return false
		}
	}
	if len(kinds) > 0 {
		matched := false
		for _, kind := range kinds {
			matchedKind := matchesKind(data, kind)
			if kind == "integer" && definition.Format == "int64" {
				_, err := strconv.ParseInt(string(data), 10, 64)
				matchedKind = err == nil
			}
			if matchedKind {
				matched = true
			}
		}
		if !matched {
			return false
		}
	}
	if data[0] == '"' {
		var value string
		if json.Unmarshal(data, &value) != nil {
			return false
		}
		length := utf8.RuneCountInString(value)
		if definition.MinLength != nil && length < *definition.MinLength {
			return false
		}
		if definition.MaxLength != nil && length > *definition.MaxLength {
			return false
		}
		if definition.Pattern != "" {
			pattern, err := regexp.Compile(definition.Pattern)
			if err != nil || !pattern.MatchString(value) {
				return false
			}
		}
		if definition.Format == "date-time" {
			if _, err := time.Parse(time.RFC3339Nano, value); err != nil {
				return false
			}
		}
		if definition.Format == "uri" {
			parsed, err := url.Parse(value)
			if err != nil || parsed.Scheme == "" {
				return false
			}
		}
		if definition.Format == "email" {
			address, err := mail.ParseAddress(value)
			if err != nil || address.Address != value {
				return false
			}
		}
	}
	if matchesKind(data, "number") {
		number, _ := strconv.ParseFloat(string(data), 64)
		if definition.Minimum != nil && number < *definition.Minimum {
			return false
		}
		if definition.Maximum != nil && number > *definition.Maximum {
			return false
		}
	}
	if data[0] == '[' {
		var values []json.RawMessage
		if json.Unmarshal(data, &values) != nil {
			return false
		}
		if definition.MinItems != nil && len(values) < *definition.MinItems {
			return false
		}
		if definition.MaxItems != nil && len(values) > *definition.MaxItems {
			return false
		}
		if definition.UniqueItems {
			seen := make(map[string]bool, len(values))
			for _, raw := range values {
				var value string
				if len(raw) == 0 || raw[0] != '"' || json.Unmarshal(raw, &value) != nil || seen[value] {
					return false
				}
				seen[value] = true
			}
		}
		if definition.Items != nil {
			for _, value := range values {
				if !child(value, *definition.Items) {
					return false
				}
			}
		}
	}
	if data[0] == '{' {
		var fields map[string]json.RawMessage
		if json.Unmarshal(data, &fields) != nil {
			return false
		}
		if definition.MinProperties != nil && len(fields) < *definition.MinProperties {
			return false
		}
		if definition.MaxProperties != nil && len(fields) > *definition.MaxProperties {
			return false
		}
		for _, required := range definition.Required {
			if _, exists := fields[required]; !exists {
				return false
			}
		}
		for name, value := range fields {
			if next, exists := definition.Properties[name]; exists {
				if !property(name, value, next) {
					return false
				}
				continue
			}
			additional := bytes.TrimSpace(definition.Additional)
			if bytes.Equal(additional, []byte("false")) {
				return false
			}
			if len(additional) > 0 && additional[0] == '{' {
				var next schema
				if json.Unmarshal(additional, &next) != nil || !child(value, next) {
					return false
				}
			}
		}
	}
	return true
}

// openEnum reports whether a response may carry a value newer than this
// single-value enum: a string enum on a named property that is neither
// `object` nor the discriminator of an enclosing union.
func openEnum(values []JSONValue, field wireField, discriminator string) bool {
	if field.name == "" || field.name == "object" || field.name == discriminator {
		return false
	}
	for _, value := range values {
		encoded, _ := value.MarshalJSON()
		if len(encoded) == 0 || encoded[0] != '"' {
			return false
		}
	}
	return true
}

func matchesKind(data []byte, kind string) bool {
	switch kind {
	case "null":
		return bytes.Equal(data, []byte("null"))
	case "object":
		return data[0] == '{'
	case "array":
		return data[0] == '['
	case "string":
		return data[0] == '"'
	case "boolean":
		return bytes.Equal(data, []byte("true")) || bytes.Equal(data, []byte("false"))
	case "integer", "number":
		if data[0] != '-' && (data[0] < '0' || data[0] > '9') {
			return false
		}
		number, err := strconv.ParseFloat(string(data), 64)
		return err == nil && !math.IsNaN(number) && !math.IsInf(number, 0) && (kind == "number" || (number == math.Trunc(number) && math.Abs(number) <= 9007199254740991))
	default:
		return false
	}
}

func equalJSON(left, right []byte) bool {
	var l, r bytes.Buffer
	if json.Compact(&l, left) != nil || json.Compact(&r, right) != nil {
		return false
	}
	return bytes.Equal(l.Bytes(), r.Bytes())
}
