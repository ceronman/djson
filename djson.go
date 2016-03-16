package djson

// Value represents any posible JSON value. It's a union-like struct
// that can be any of the posible JSON values: number, string, true, false,
// array, object or null
// It has an interal `kind` flag that indicates the type
type Value struct {
	kind          valueKind
	numberContent float64
	stringContent string
	arrayContent  []Value
	objectContent map[string]Value
}

type valueKind int8

const (
	nullKind   = iota
	numberKind = iota
	stringKind = iota
	trueKind   = iota
	falseKind  = iota
	objectKind = iota
	arrayKind  = iota
)
