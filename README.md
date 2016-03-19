# About 

Additional checkers collection for gocheck library. 
Currently has only one checker - DeepEqualsPointer. 

This checker acts like default DeepEquals but will try dereferencing each of passed argument and check for direct value equality instead of passed.
Checker uses reflect.DeepEquals inside for equality checks as original.

# Example

```
package my_test
import (
	. "github.com/soider/gocheck-checkers/checkers"
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
