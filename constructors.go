package gomakejson

func (jo Object) Add(name string, val JsonVal) Object {
	jo.Fields = append(jo.Fields, Field{name, val})
	return jo
}

func MakeArr(jsonVals ...JsonVal) Array {
	return Array{jsonVals}
}
