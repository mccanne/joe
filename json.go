//
// joe is a simple JSON object extractor (joe) for Go
//
// joe provides a simple API to access unstructured and ad hoc JSON objects
// that is parsed generically by json.Unmarshal.  When JSON inputs are
// unstructured, it can be difficult to define Go structs that map cleanly
// onto the JSON input.
//
// Joe provides all the reflection needed into the empty interface types
// that Unmarshall creates but with a javascripty set of method calls.
//
package joe

import (
	"encoding/json"
)

//
// JSON defines a generic type for a value extracted from JSON input and into go.
// Method calls on this type perform reflection and extract specific values.
//
type JSON struct {
	v interface{}
}

var Null = JSON{nil}
var Undefined = JSON{Null}

//
// Parse produces a JSON object from its input, which is a string encoded as JSON.
//
func Parse(in []byte) (JSON, error) {
	var result JSON
	err := json.Unmarshal(in, &result.v)
	return result, err
}

//
// UnmarshalJSON implements JSON's custom parser so anything declared as type joe.JSON
// will be decoded by go's json package into an object for this interface
//
func (p *JSON) UnmarshalJSON(in []byte) error {
	return json.Unmarshal(in, &p.v)
}

//
// Value returns the generic value of this JSON object
//
func (o JSON) Value() interface{} {
	return o.v
}

//
// IsNull returns true iff this object is a null object
//
func (o JSON) IsNull() bool {
	return o.v == nil
}

//
// IsUndefined returns true iff this object is the Undefined object that occurs when
// accessing Object keys or Array indices through this interface that do not exist.
//
func (o JSON) IsUndefined() bool {
	return o == Undefined
}

//
// Number returns two values: (1) this JSON object as a number and (2) an indication
// of success if it is a number in the underlying JSON representation.  Otherwise, it
// returns an indication of failure.
//
func (o JSON) Number() (float64, bool) {
	if v, ok := o.v.(float64); ok {
		return v, true
	}
	return 0.0, false
}

//
// IsNumber returns true iff this object is a number in the underlying
// JSON representation.
//
func (o JSON) IsNumber() bool {
	_, ok := o.v.(float64)
	return ok
}

//
// String returns two values: (1) this JSON object as a string and (2) an indication
// of success if it is a string in the underlying JSON representation.  Otherwise, it
// returns an indication of failure.
//
func (o JSON) String() (string, bool) {
	if v, ok := o.v.(string); ok {
		return v, true
	}
	return "", false
}

//
// IsString returns true iff this object is a string in the underlying
// JSON representation.
//
func (o JSON) IsString() bool {
	_, v := o.String()
	return v
}

//
// Boolean returns two values: (1) this JSON object as a string and (2) an indication
// of success if it is a string in the underlying JSON representation.  Otherwise, it
// returns an indication of failure.
//
func (o JSON) Boolean() (bool, bool) {
	if v, ok := o.v.(bool); ok {
		return v, true
	}
	return false, false
}

//
// IsBoolean returns true iff this object is a boolean in the underlying
// JSON representation.
//
func (o JSON) IsBoolean() bool {
	_, v := o.Boolean()
	return v
}

//
// IsArray returns true iff this object is an array in the underlying
// JSON representation.
//
func (o JSON) IsArray() bool {
	_, ok := o.v.([]interface{})
	return ok
}

//
// Index returns the element indicated by the offset argument as a JSON object
// presuming this object is an array.  If not, it returns the Undefinied JSON object.
//
func (o JSON) Index(offset int) JSON {
	if a, ok := o.v.([]interface{}); ok {
		if offset >= 0 && offset < len(a) {
			return JSON{a[offset]}
		}
	}
	return Undefined
}

//
// Len returns the length of the underlying array presuming this object is an array.
// If not, it returns -1.
//
func (o JSON) Len() int {
	if v, ok := o.v.([]interface{}); ok {
		return len(v)
	}
	return -1
}

//
// IsObject returns true iff this object is a JSON object in the underlying
// JSON representation.
//
func (o JSON) IsObject() bool {
	_, ok := o.v.(map[string]interface{})
	return ok
}

//
// Get returns the element indicated by the key argument as a JSON object
// presuming this object is an JSON object.  If not, it returns the Undefinied JSON object.
//
func (o JSON) Get(key string) JSON {
	if m, ok := o.v.(map[string]interface{}); ok {
		v, ok := m[key]
		if ok {
			return JSON{v}
		}
	}
	return Undefined
}
