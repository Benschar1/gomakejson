package main

// JsonVal is the interface representing a Json value.
// It implements the Marshaler interface making it
// compatible with the encoding/json package.
//
// There are six types implementing JsonVal:
//   - Object
//   - Array
//   - String
//   - Number
//   - Bool
//   - Null
//
// Because JsonVal has the unexported method isJsonVal(),
// it cannot be implemented by other other types.
//
type JsonVal interface {
	isJsonVal()
	MarshalJSON() ([]byte, error)
}

func (x Object) isJsonVal()    {}
func (x Array) isJsonVal()     {}
func (x String) isJsonVal()    {}
func (x Number[N]) isJsonVal() {}
func (x Bool) isJsonVal()      {}
func (x Null) isJsonVal()      {}

type Object struct {
	Fields []Field
}

type Array struct {
	Arr []JsonVal
}

type String struct {
	String string
}

type Number[N NumberConstraint] struct {
	Number N
}

type Bool struct {
	Boolean bool
}

type Null struct{}

// utils

type Field struct {
	Name string
	Val  JsonVal
}

type NumberConstraint interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}
