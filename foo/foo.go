package foo

var nextID = 0

type Foo interface {
	GetContent() string
}

// Foo is a type that contains a Content string
type SimpleFoo struct {
	ID      int
	Content string
}

// New returns a Foo with an empty Content string
func New(content ...string) *SimpleFoo {
	f := SimpleFoo{ID: nextID}
	for _, c := range content {
		f.Content = c
	}
	nextID++
	return &f
}

func (f SimpleFoo) GetContent() string {
	return f.Content
}
