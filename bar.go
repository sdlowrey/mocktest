package bar

import "github.com/sdlowrey/mocktest/foo"

type NamedRecord struct {
	Name  string
	Entry foo.Foo
}

func New(name string) *NamedRecord {
	r := NamedRecord{}
	r.Name = name
	r.Entry = foo.New(name)
	return &r
}

func (r NamedRecord) GetContent() string {
	return r.Entry.GetContent()
}
