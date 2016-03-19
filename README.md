# About 

Additional checker for gocheck library. This checker will try dereferencing each of passed argument and check for direct value equality.
Checker uses reflect.DeepEquals inside for equality checks.

# Example

```
package my_test
import (
	. "github.com/soider/gocheck-pointers/checkers"
	. "gopkg.in/check.v1"
	"testing"
)

func (s *MySuite) TestIntWithPointerInt(c *C) {
    v := 4
    v1 := &v
	c.Assert(v, DeepEqualsPointer, v1)
}

```

A bit more examples in examples/example_test.go
