package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

type CustomFieldJSON map[string]interface{}

// MarshalCustomFieldJSON marshals the CustomFieldJSON into a writer (typically for GraphQL response)
func MarshalCustomFieldJSON(val CustomFieldJSON) graphql.Marshaler {
    return graphql.WriterFunc(func(w io.Writer) {
        err := json.NewEncoder(w).Encode(val)
        if err != nil {
            panic(err)
        }
    })
}

// UnmarshalCustomFieldJSON unmarshals a map into CustomFieldJSON from the given value
func UnmarshalCustomFieldJSON(v any) (CustomFieldJSON, error) {
    if m, ok := v.(map[string]any); ok {
        return CustomFieldJSON(m), nil
    }

    return nil, fmt.Errorf("%T is not a map", v)
}

// Scan implements the sql.Scanner interface for CustomFieldJSON
func (j *CustomFieldJSON) Scan(value interface{}) error {
    if value == nil {
        *j = nil
        return nil
    }
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    var temp map[string]interface{}
    if err := json.Unmarshal(bytes, &temp); err != nil {
        return err
    }
    *j = temp
    return nil
}

// Value implements the driver.Valuer interface for CustomFieldJSON
func (j CustomFieldJSON) Value() (driver.Value, error) {
    if j == nil {
        return nil, nil
    }
    return json.Marshal(j)
}

