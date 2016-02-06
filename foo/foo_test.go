package foo

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

type FooSuite struct{}

var _ = check.Suite(&FooSuite{})

func (s *FooSuite) TestCreate(c *check.C) {
	f1 := New()
	c.Assert(f1.ID, check.Equals, 0)
	c.Assert(f1.Content, check.Equals, "")

	f2 := New()
	c.Assert(f2.ID, check.Equals, 1)
}

func (s *FooSuite) TestInit(c *check.C) {
	content := "Foo 1"
	f := New(content)
	c.Assert(f.Content, check.Equals, content)
}

func (s *FooSuite) TestGetters(c *check.C) {
	content := "wut"
	f := New(content)
	c.Assert(f.GetContent(), check.Equals, content)
}
