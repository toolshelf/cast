package internal

import "encoding/json"

// JsonStringToObject attempts to unmarshall a string as JSON into
// the object passed as pointer.
func JsonStringToObject(s string, v interface{}) error {
	data := []byte(s)
	return json.Unmarshal(data, v)
}
