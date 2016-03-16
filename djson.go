package djson

import "errors"

// Value represents any posible JSON value. It's a union-like struct
// that can be any of the posible JSON values: number, string, bool,
// array, object or null.
type Value struct {
	kind          kind
	boolContent   bool
	numberContent float64
	stringContent string
	arrayContent  []Value
	objectContent map[string]Value
}

type kind int8

const (
	kindNull kind = iota
	kindBool
	kindNumber
	kindString
	kindObject
	kindArray
)

// Bool returns the boolean representation of a JSON value if the value
// is actually a boolean otherwise returns an error.
func (v *Value) Bool() (bool, error) {
	if v.kind == kindBool {
		return v.boolContent, nil
	}
	return false, nil
}

// Number returns the number representation as a float64 of a JSON value if
// that value is actually a number, otherwise returns an error.
func (v *Value) Number() (float64, error) {
	if v.kind == kindNumber {
		return v.numberContent, nil
	}
	return 0.0, errors.New("JSON value is not a number")
}

// String returns the string representation of a JSON value if that value
// is actually a string, otherwise returns an error.
func (v *Value) String() (string, error) {
	if v.kind == kindString {
		return v.stringContent, nil
	}
	return "", errors.New("JSON value is not a string")
}

// Object returns the object represenation of the JSON value as a map if
// the value is actually an object, otherwise returns an error.
func (v *Value) Object() (map[string]Value, error) {
	if v.kind == kindObject {
		return v.objectContent, nil
	}
	return nil, errors.New("JSON value is not an object")
}

// Array returns the object represenation of the JSON value as a map if
// the value is actually an object, otherwise returns an error.
func (v *Value) Array() ([]Value, error) {
	if v.kind == kindArray {
		return v.arrayContent, nil
	}
	return nil, errors.New("JSON value is not an object")
}

// NewBool creates a new JSON Value object representing a boolean.
func NewBool(value bool) *Value {
	return &Value{kind: kindBool, boolContent: value}
}

// NewNumber creates a new JSON Value object representing a number.
func NewNumber(value float64) *Value {
	return &Value{kind: kindNumber, numberContent: value}
}

// NewString creates a new JSON Value object representing a string.
func NewString(value string) *Value {
	return &Value{kind: kindString, stringContent: value}
}

// NewObject creates a new JSON Value object representing an object.
func NewObject() *Value {
	return &Value{kind: kindObject, objectContent: make(map[string]Value)}
}

// NewArray creates a new JSON Value object representing an array.
func NewArray() *Value {
	// TODO: think about size and capacity defaults
	return &Value{kind: kindBool, arrayContent: make([]Value, 0, 1)}
}
