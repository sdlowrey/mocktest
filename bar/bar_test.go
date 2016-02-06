package bar

import (
	"fmt"
	"testing"

	"gopkg.in/check.v1"
)

// Test registers the goCheck test harness with the go test framework.
func Test(t *testing.T) {
	check.TestingT(t)
}

// TestError is a dummy error type used for testing.
type TestError struct {
	Message string // Optional message to include in dummy message.
}

// Error satisfies the error interface for Error.
func (e *TestError) Error() string {
	return fmt.Sprintf("Error: %v", e.Message)
}

type BarSuite struct{}

var _ = check.Suite(&BarSuite{})

func (s *BarSuite) TestCreate(c *check.C) {
	b1 := New("test1")
	c.Assert(b1.Name, check.Equals, "test1")
	c.Assert(b1.GetContent(), check.Equals, "test1")
}
