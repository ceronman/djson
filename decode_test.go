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

func expectBool(t *testing.T, input string, expected bool) {
	json, err := djson.Decode([]byte(input))
	if err != nil {
		t.Error(err)
		return
	}

	value, err := json.Bool()
	if err != nil {
		t.Error(err)
		return
	}

	if value != expected {
		t.Errorf("Unexpected: %v != %v", value, expected)
	}
}

func TestDecodeBool(t *testing.T) {
	expectBool(t, `true`, true)
	expectBool(t, `false`, false)
	expectBool(t, `   true`, true)
	expectBool(t, `true   `, true)
	expectBool(t, `  true  `, true)
	expectBool(t, `  false`, false)
	expectBool(t, `false  `, false)
	expectBool(t, `  false  `, false)

	expectError(t, `False`)
	expectError(t, `FALSE`)
	expectError(t, ` False `)
	expectError(t, `True`)
	expectError(t, `TRUE`)
	expectError(t, ` True `)
}

func expectNumber(t *testing.T, input string, expected float64) {
	json, err := djson.Decode([]byte(input))
	if err != nil {
		t.Error(err)
		return
	}

	value, err := json.Number()
	if err != nil {
		t.Error(err)
		return
	}

	if value != expected {
		t.Errorf("Unexpected: %v != %v", value, expected)
		return
	}
}

func TestDecodeNumberInteger(t *testing.T) {
	expectNumber(t, `12345`, 12345)
	expectNumber(t, `1`, 1)
	expectNumber(t, `  2 `, 2)
	expectNumber(t, `-12345`, -12345)
	expectNumber(t, ` -3  `, -3)

	expectError(t, `0123`)
	expectError(t, `+123`)
	expectError(t, `--123`)
	expectError(t, `-0123`)
	expectError(t, `123A`)
	expectError(t, `0x1230`)
	expectError(t, `NaN`)
}

func TestDecodeNumberFloat(t *testing.T) {
	expectNumber(t, `123.456`, 123.456)
	expectNumber(t, `1.0`, 1.0)
	expectNumber(t, ` 2.0  `, 2.0)
	expectNumber(t, `-123.456`, -123.456)
	expectNumber(t, ` -3.1  `, -3.1)
	expectNumber(t, `0.9`, 0.9)
	expectNumber(t, `0.08`, 0.08)
	expectNumber(t, `1.5e10`, 1.5e+10)
	expectNumber(t, `1.5E10`, 1.5e+10)
	expectNumber(t, `-1.5e10`, -1.5e+10)
	expectNumber(t, `-1.5E10`, -1.5e+10)
	expectNumber(t, `1.5e+10`, 1.5e+10)
	expectNumber(t, `1.5E+10`, 1.5e+10)
	expectNumber(t, `-1.5e+10`, -1.5e+10)
	expectNumber(t, `-1.5E+10`, -1.5e+10)
	expectNumber(t, `1.5e-10`, 1.5e-10)
	expectNumber(t, `1.5E-10`, 1.5e-10)
	expectNumber(t, `-1.5e-10`, -1.5e-10)
	expectNumber(t, `-1.5E-10`, -1.5e-10)
	expectNumber(t, `-2E10`, -2e+10)
	expectNumber(t, `-2E-10`, -2e-10)
	expectNumber(t, `-2E+10`, -2e+10)
	expectNumber(t, `-2e10`, -2e+10)
	expectNumber(t, `-2e-10`, -2e-10)
	expectNumber(t, `-2e+10`, -2e+10)

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

func expectString(t *testing.T, input string, expected string) {
	json, err := djson.Decode([]byte(input))
	if err != nil {
		t.Error(err)
		return
	}

	value, err := json.String()
	if err != nil {
		t.Error(err)
		return
	}

	if value != expected {
		t.Errorf("Unexpected: %v != %v", value, expected)
		return
	}
}

func TestDecodeString(t *testing.T) {
	expectString(t, `""`, "")
	expectString(t, `"Hello World"`, "Hello World")
	expectString(t, `"  'single quotes' "`, "  'single quotes' ")
	expectString(t, `  "spaces"   `, "spaces")
	expectString(t, `"\""`, `"`)
	expectString(t, `"\\"`, `\`)
	expectString(t, `"\/"`, "/")
	expectString(t, `"\b"`, "\b")
	expectString(t, `"\f"`, "\f")
	expectString(t, `"\n"`, "\n")
	expectString(t, `"\t"`, "\t")
	expectString(t, `"\r"`, "\r")
	expectString(t, `"\u12e4"`, "\u12e4")
	expectString(t, `"\u12e4 \u12e5"`, "\u12e4 \u12e5")
	expectString(t, `   "\u12e4"   `, "\u12e4")
	expectString(t, `"  \u12e4  "`, "  \u12e4  ")
	expectString(t, `"\n\\\""`, "\n\\\"")

	expectError(t, `"`)
	expectError(t, `"""`)
	expectError(t, `"one"two"`)
	expectError(t, `'single quotes'`)
	expectError(t, `"\u12"`)
	expectError(t, `123"string"two`)
}
