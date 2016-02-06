package foo

var nextID = 0

// Foo is a type that contains a Content string
type Foo struct {
	ID      int
	Content string
}

// New returns a Foo with an empty Content string
func New(content ...string) *Foo {
	f := Foo{ID: nextID}
	for _, c := range content {
		f.Content = c
	}
	nextID++
	return &f
}
