package checkers

import (
	"gopkg.in/check.v1"
	"reflect"
)

type pointerValueChecker struct {
}

var DeepEqualsPointer = pointerValueChecker{}

func (c pointerValueChecker) Info() *check.CheckerInfo {
	return &check.CheckerInfo{Name: "PointerValueChecker", Params: []string{"obtained", "expected"}}
}

func (c pointerValueChecker) Check(params []interface{}, names []string) (result bool, error string) {
	p1 := reflect.Indirect(reflect.ValueOf(params[0]))
	p2 := reflect.Indirect(reflect.ValueOf(params[1]))
	if p1.Type() != p2.Type() {
		return false, "Arguments has different types"
	}
	return reflect.DeepEqual(p1.Interface(), p2.Interface()), ""
}
