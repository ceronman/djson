package djson_test

import (
	"testing"

	"github.com/ceronman/djson"
)

func expectError(t *testing.T, input string) {
	_, err := djson.Decode([]byte(input))
	if err == nil {
		t.Error("Expecting decode error but got nil")
	}
}

func expect(t *testing.T, input string, expected interface{}) {
	json, err := djson.Decode([]byte(input))
	if err != nil {
		t.Error(err)
		return
	}

	value, err := json.Interface()
	if err != nil {
		t.Error(err)
		return
	}

	// If value is int, make things easier to test by converting to float
	intValue, ok := expected.(int)
	if ok {
		expected = float64(intValue)
	}

	if value != expected {
		t.Errorf("Unexpected: %v != %v", value, expected)
	}
}

func TestDecodeBool(t *testing.T) {
	expect(t, `true`, true)
	expect(t, `false`, false)
	expect(t, `   true`, true)
	expect(t, `true   `, true)
	expect(t, `  true  `, true)
	expect(t, `  false`, false)
	expect(t, `false  `, false)
	expect(t, `  false  `, false)

	expectError(t, `False`)
	expectError(t, `FALSE`)
	expectError(t, ` False `)
	expectError(t, `True`)
	expectError(t, `TRUE`)
	expectError(t, ` True `)
}

func TestDecodeNumberInteger(t *testing.T) {
	expect(t, `12345`, 12345)
	expect(t, `1`, 1)
	expect(t, `  2 `, 2)
	expect(t, `-12345`, -12345)
	expect(t, ` -3  `, -3)

	expectError(t, `0123`)
	expectError(t, `+123`)
	expectError(t, `--123`)
	expectError(t, `-0123`)
	expectError(t, `123A`)
	expectError(t, `0x1230`)
	expectError(t, `NaN`)
}

func TestDecodeNumberFloat(t *testing.T) {
	expect(t, `123.456`, 123.456)
	expect(t, `1.0`, 1.0)
	expect(t, ` 2.0  `, 2.0)
	expect(t, `-123.456`, -123.456)
	expect(t, ` -3.1  `, -3.1)
	expect(t, `0.9`, 0.9)
	expect(t, `0.08`, 0.08)
	expect(t, `1.5e10`, 1.5e+10)
	expect(t, `1.5E10`, 1.5e+10)
	expect(t, `-1.5e10`, -1.5e+10)
	expect(t, `-1.5E10`, -1.5e+10)
	expect(t, `1.5e+10`, 1.5e+10)
	expect(t, `1.5E+10`, 1.5e+10)
	expect(t, `-1.5e+10`, -1.5e+10)
	expect(t, `-1.5E+10`, -1.5e+10)
	expect(t, `1.5e-10`, 1.5e-10)
	expect(t, `1.5E-10`, 1.5e-10)
	expect(t, `-1.5e-10`, -1.5e-10)
	expect(t, `-1.5E-10`, -1.5e-10)
	expect(t, `-2E10`, -2e+10)
	expect(t, `-2E-10`, -2e-10)
	expect(t, `-2E+10`, -2e+10)
	expect(t, `-2e10`, -2e+10)
	expect(t, `-2e-10`, -2e-10)
	expect(t, `-2e+10`, -2e+10)

	expectError(t, `0123.1`)
	expectError(t, `+123.3`)
	expectError(t, `--123.5`)
	expectError(t, `-0123.5`)
	expectError(t, `123A.1`)
	expectError(t, `0x1230.5`)
	expectError(t, `11.-`)
	expectError(t, `12.+`)
	expectError(t, `13.e+10`)
	expectError(t, `13.e+10`)
	expectError(t, `13.2e+10.1`)
	expectError(t, `NaN`)
}

func TestDecodeString(t *testing.T) {
	expect(t, `""`, "")
	expect(t, `"Hello World"`, "Hello World")
	expect(t, `"  'single quotes' "`, "  'single quotes' ")
	expect(t, `  "spaces"   `, "spaces")
	expect(t, `"\""`, `"`)
	expect(t, `"\\"`, `\`)
	expect(t, `"\/"`, "/")
	expect(t, `"\b"`, "\b")
	expect(t, `"\f"`, "\f")
	expect(t, `"\n"`, "\n")
	expect(t, `"\t"`, "\t")
	expect(t, `"\r"`, "\r")
	expect(t, `"\u12e4"`, "\u12e4")
	expect(t, `"\u12e4 \u12e5"`, "\u12e4 \u12e5")
	expect(t, `   "\u12e4"   `, "\u12e4")
	expect(t, `"  \u12e4  "`, "  \u12e4  ")
	expect(t, `"\n\\\""`, "\n\\\"")

	expectError(t, `"`)
	expectError(t, `"""`)
	expectError(t, `"one"two"`)
	expectError(t, `'single quotes'`)
	expectError(t, `"\u12"`)
	expectError(t, `123"string"two`)
}

func expectArray(t *testing.T, input string, expected []interface{}) {
	json, err := djson.Decode([]byte(input))
	if err != nil {
		t.Error(err)
		return
	}

	length, err := json.ArrayLen()
	if err != nil {
		t.Error(err)
		return
	}
	if length != len(expected) {
		t.Error("Arrays have different length")
		return
	}

	for i, item := range expected {
		value, err := json.InterfaceAt(i)
		if err != nil {
			t.Error(err)
			return
		}
		// If value is int, make things easier to test by converting to float
		intValue, ok := item.(int)
		if ok {
			item = float64(intValue)
		}
		if item != value {
			t.Errorf("Unexpected item in array: %v != %v", value, item)
			return
		}
	}
}

func TestArrayFlat(t *testing.T) {
	expectArray(t, `[]`, []interface{}{})
	expectArray(t, `[1, 2, 3]`, []interface{}{1, 2, 3})
	expectArray(t, `[1,2,3]`, []interface{}{1, 2, 3})
	expectArray(t, `[1, 0.5, 1.25e+12]`, []interface{}{1, 0.5, 1.25e+12})
	expectArray(t, `[true, false]`, []interface{}{true, false})
	expectArray(t, `["hello", "world"]`, []interface{}{"hello", "world"})
	expectArray(t, `[1, "hello", false]`, []interface{}{1, "hello", false})
	expectArray(t, `  []   `, []interface{}{})
	expectArray(t, ` [1, "one", "two"]  `, []interface{}{1, "one", "two"})

	expectError(t, `[`)
	expectError(t, `[1, 2, 3`)
	expectError(t, `]`)
	expectError(t, `1, 2]`)
	expectError(t, `[1, 2, 3,]`)
	expectError(t, `[, 2, 3]`)
	expectError(t, `[1, 2, 3]]`)
	expectError(t, `[[1, 2, 3]`)
}
