package type_test

import (
	. "github.com/soider/gocheck-pointers/checkers"
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

// MySuite
type MySuite struct{}

var _ = Suite(&MySuite{})

func returnInt(v int) int {
	return v
}

func returnIntPointer(v int) *int {
	return &v
}

func returnStr(v string) string {
	return v
}

func returnStrPointer(v string) *string {
	return &v
}

type myStruct struct {
	Name string
}

type myStruct2 struct {
	Name string
}

func (s *MySuite) TestIntWithPointerInt(c *C) {
	c.Assert(returnInt(100), DeepEqualsPointer, returnIntPointer(100))
}

func (s *MySuite) TestStringWithPointerToString(c *C) {
	c.Assert(returnStr("HelloWorld"), DeepEqualsPointer, returnStrPointer("HelloWorld"))
}

func (s *MySuite) TestStructWithPointerToStruct(c *C) {
	c.Assert(myStruct{Name: "1"}, DeepEqualsPointer, &myStruct{Name: "1"})
}

func (s *MySuite) TestIntWithPointerIntFalse(c *C) {
	c.ExpectFailure("Should fail")
	c.Assert(returnInt(100), DeepEqualsPointer, returnIntPointer(101))
}

func (s *MySuite) TestStringWithPointerToStringFalse(c *C) {
	c.ExpectFailure("Should fail")
	c.Assert(returnStr("HelloWorld"), DeepEqualsPointer, returnStrPointer("HelloWorld1"))
}

func (s *MySuite) TestStructWithPointerToStructFalse(c *C) {
	c.ExpectFailure("Should fail")
	c.Assert(myStruct{Name: "1"}, DeepEqualsPointer, &myStruct2{Name: "1"})
}
