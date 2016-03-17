package djson

import "testing"

func testDecodeBool(t *testing.T, input string, expected bool) {
	json, err := Decode([]byte(input))
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
		return
	}
}

func TestDecodeBool(t *testing.T) {
	testDecodeBool(t, `true`, true)
	testDecodeBool(t, `false`, false)
	testDecodeBool(t, `   true`, true)
	testDecodeBool(t, `true   `, true)
	testDecodeBool(t, `  true  `, true)
	testDecodeBool(t, `  false`, false)
	testDecodeBool(t, `false  `, false)
	testDecodeBool(t, `  false  `, false)
}
