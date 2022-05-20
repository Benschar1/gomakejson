package gomakejson

import (
	"bytes"
	"encoding/json"
)

func (jo Object) MarshalJSON() ([]byte, error) {
	fieldVals := make([][]byte, len(jo.Fields))

	for i, v := range jo.Fields {
		nameBytes, _ := json.Marshal(v.Name)
		valBytes, _ := json.Marshal(v.Val)
		bytes := append(nameBytes, []byte(":")...)
		fieldVals[i] = append(bytes, valBytes...)
	}

	bytes := append([]byte("{"), bytes.Join(fieldVals, []byte(","))...)

	return append(bytes, []byte("}")...), nil
}

func (ja Array) MarshalJSON() ([]byte, error) {
	arrVals := make([][]byte, len(ja.Arr))

	for i, v := range ja.Arr {
		jstr, _ := json.Marshal(v)
		arrVals[i] = jstr
	}

	bytes := append([]byte("["), bytes.Join(arrVals, []byte(","))...)

	return append(bytes, []byte("]")...), nil
}

func (js String) MarshalJSON() ([]byte, error) {
	return json.Marshal(js.String)
}

func (jn Number[N]) MarshalJSON() ([]byte, error) {
	return json.Marshal(jn.Number)
}

func (jb Bool) MarshalJSON() ([]byte, error) {
	boolstr := "false"
	if jb.Boolean {
		boolstr = "true"
	}
	return []byte(boolstr), nil
}

func (jn Null) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}
