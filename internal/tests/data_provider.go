package tests

import (
	"context"
	"encoding/json"
)

type dataProvider struct {
	data map[string]json.RawMessage
}

// Get implemnts the permission.DataProvider interface by returning the values
// from data.
func (t dataProvider) Get(ctx context.Context, fqfields ...string) ([]json.RawMessage, error) {
	data := make([]json.RawMessage, len(fqfields))
	for i, field := range fqfields {
		value, ok := t.data[field]
		if !ok {
			data[i] = nil
			continue
		}

		data[i] = json.RawMessage(value)
	}
	return data, nil
}
