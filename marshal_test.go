package main

import (
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {

	tests := map[string]struct {
		input JsonVal
		want  string
	}{
		"empty object": {Object{}, "{}"},
		"object": {
			Object{}.
				Add("first name", String{"John"}).
				Add("last name", String{"Doe"}).
				Add("age", Number[int]{25}),
			"{\"first name\":\"John\",\"last name\":\"Doe\",\"age\":25}",
		},
		"object with duplicate fields": {
			Object{}.
				Add("field", String{}).
				Add("field", String{}).
				Add("field", String{}),
			"{\"field\":\"\",\"field\":\"\",\"field\":\"\"}",
		},
		"array": {
			MakeArr(
				String{"John"},
				String{"Doe"},
				Number[int]{25},
			),
			"[\"John\",\"Doe\",25]"},
		"empty array":  {Array{}, "[]"},
		"string":       {String{"abc123"}, "\"abc123\""},
		"empty string": {String{""}, "\"\""},
		"int":          {Number[int]{3}, "3"},
		"float":        {Number[float64]{3.12}, "3.12"},
		"bool true":    {Bool{true}, "true"},
		"bool false":   {Bool{false}, "false"},
		"null":         {Null{}, "null"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			bytes, err := tc.input.MarshalJSON()
			str := string(bytes)

			if err != nil {
				t.Fatalf("expected: %v, got error: %v", tc.want, err)
			}

			if !reflect.DeepEqual(tc.want, str) {
				t.Fatalf("expected: %v, got: %v", tc.want, str)
			}
		})
	}
}
