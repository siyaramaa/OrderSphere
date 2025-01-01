package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

type CustomField struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	IsRequired bool   `json:"isRequired"`
}

type CustomFieldArray []CustomField


// MarshalCustomFieldSchemasData marshals the CustomFieldArray into a writer for GraphQL response
func MarshalCustomFieldArray(val CustomFieldArray) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err) // Replace with error logging if needed
		}
	})
}


// UnmarshalCustomFieldSchemasData unmarshals an array into CustomFieldArray from the given value
func UnmarshalCustomFieldArray(v any) (CustomFieldArray, error) {
	switch v := v.(type) {
	case []interface{}:
		var result CustomFieldArray
		for _, item := range v {
			m, ok := item.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("expected map for item, got %T", item)
			}

			field := CustomField{
				Name:       m["name"].(string),
				Type:       m["type"].(string),
				IsRequired: m["isRequired"].(bool),
			}
			result = append(result, field)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("expected []interface{}, got %T", v)
	}
}


func (c *CustomFieldArray) Scan(value interface{}) error {
	if value == nil {
		*c = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	var temp []CustomField
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return err
	}

	*c = temp
	return nil
}

// Value implements the driver.Valuer interface for CustomFieldJSON
func (j CustomFieldArray) Value() (driver.Value, error) {
    if j == nil {
        return nil, nil
    }
    return json.Marshal(j)
}

