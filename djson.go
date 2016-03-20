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

// Interface return the representation of the JSON value as an interface{}.
// Available for convenience.
func (v *Value) Interface() (interface{}, error) {
	switch v.kind {
	case kindBool:
		return v.boolContent, nil
	case kindNumber:
		return v.numberContent, nil
	case kindString:
		return v.stringContent, nil
	case kindArray:
		return v.arrayContent, nil
	case kindObject:
		return v.objectContent, nil
	default:
		return nil, errors.New("Unknown type of JSON value")
	}
}

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
	return nil, errors.New("JSON value is not an array")
}

// ArrayLen returns the lenght of an array if the JSON value is an array
// otherwise an error.
func (v *Value) ArrayLen() (int, error) {
	if v.kind == kindArray {
		return len(v.arrayContent), nil
	}
	return 0, errors.New("JSON value is not an array")
}

// InterfaceAt returns the element at a given index of an array as an
// interface{} if the JSON value is an array, otherwise returns an error.
// This method is added for convenience.
func (v *Value) InterfaceAt(i int) (interface{}, error) {
	if v.kind != kindArray {
		return false, errors.New("JSON value is not an array")
	}
	if i >= len(v.arrayContent) {
		return false, errors.New("Index is out of bounds of array")
	}
	switch v.arrayContent[i].kind {
	case kindBool:
		return v.arrayContent[i].boolContent, nil
	case kindNumber:
		return v.arrayContent[i].numberContent, nil
	case kindString:
		return v.arrayContent[i].stringContent, nil
	case kindArray:
		return v.arrayContent[i].arrayContent, nil
	case kindObject:
		return v.arrayContent[i].objectContent, nil
	default:
		return nil, errors.New("Unknown type of JSON value")
	}
}

// BoolAt returns the boolean in at a given index of an array if the JSON value
// is an array, otherwise an error.
func (v *Value) BoolAt(i int) (bool, error) {
	if v.kind != kindArray {
		return false, errors.New("JSON value is not an array")
	}
	if i >= len(v.arrayContent) {
		return false, errors.New("Index is out of bounds of array")
	}
	value, err := v.arrayContent[i].Bool()
	if err != nil {
		return false, err
	}
	return value, nil
}

// NumberAt returns the number in at a given index of an array if the JSON value
// is an array, otherwise an error.
func (v *Value) NumberAt(i int) (float64, error) {
	if v.kind != kindArray {
		return 0.0, errors.New("JSON value is not an array")
	}
	if i >= len(v.arrayContent) {
		return 0.0, errors.New("Index is out of bounds of array")
	}
	value, err := v.arrayContent[i].Number()
	if err != nil {
		return 0.0, err
	}
	return value, nil
}

// StringAt returns the string in at a given index of an array if the JSON value
// is an array, otherwise an error.
func (v *Value) StringAt(i int) (string, error) {
	if v.kind != kindArray {
		return "", errors.New("JSON value is not an array")
	}
	if i >= len(v.arrayContent) {
		return "", errors.New("Index is out of bounds of array")
	}
	value, err := v.arrayContent[i].String()
	if err != nil {
		return "", err
	}
	return value, nil
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
	return &Value{kind: kindArray, arrayContent: make([]Value, 0)}
}
