package i18n

import (
	"encoding/json"
	"strings"
)

// Dict holds the internationalization entries for a specific locale.
type Dict struct {
	value   string
	entries map[string]*Dict
}

// NewDict instantiates a new dict object.
func NewDict() *Dict {
	return &Dict{
		entries: make(map[string]*Dict),
	}
}

// Add adds a new key value pair to the dictionary.
func (d *Dict) Add(key string, value any) {
	switch v := value.(type) {
	case string:
		d.entries[key] = &Dict{value: v}
	case map[string]any:
		nd := NewDict()
		for k, row := range v {
			nd.Add(k, row)
		}
		d.entries[key] = nd
	case *Dict:
		d.entries[key] = v
	default:
		// ignore
	}
}

// Value returns the dictionary value or an empty string
// if the dictionary is nil.
func (d *Dict) Value() string {
	if d == nil {
		return ""
	}
	return d.value
}

// Get recursively retrieves the dictionary at the provided key location.
func (d *Dict) Get(key string) *Dict {
	if d == nil {
		return nil
	}
	if key == "" {
		return nil
	}
	n := strings.SplitN(key, ".", 2)
	entry, ok := d.entries[n[0]]
	if !ok {
		return nil
	}
	if len(n) == 1 {
		return entry
	}
	return entry.Get(n[1])
}

// Has is a convenience method to check if a key exists in the dictionary
// recursively, and is the equivalent of calling `Get` and checking if
// the result is not nil.
func (d *Dict) Has(key string) bool {
	return d.Get(key) != nil
}

// Merge combines the entries of the second dictionary into this one. If a
// key is duplicated in the second diction, the original value takes priority.
func (d *Dict) Merge(d2 *Dict) {
	if d2 == nil {
		return
	}
	if d.entries == nil {
		d.entries = make(map[string]*Dict)
	}
	for k, v := range d2.entries {
		if d.entries[k] == nil {
			d.entries[k] = v
			continue
		}
		d.entries[k].Merge(v)
	}
}

// UnmarshalJSON attempts to load the dictionary data from a JSON byte slice.
func (d *Dict) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] == '"' {
		d.value = string(data[1 : len(data)-1])
		return nil
	}
	d.entries = make(map[string]*Dict)
	return json.Unmarshal(data, &d.entries)
}
