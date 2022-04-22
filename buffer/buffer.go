package buffer

import (
	"fmt"
	"strings"
)

// this represents the internal Gap Buffer data structure. Before
type Buffer struct {
	before []rune
	edit   []rune
	after  []rune
	cursor int
}

func NewBuffer() Buffer {
	return Buffer{
		before: make([]rune, 0),
		edit:   make([]rune, 0),
		after:  make([]rune, 0),
	}
}

func (b *Buffer) Insert(s []rune) {
	b.edit = append(b.edit, s...)
}

func (b *Buffer) String() string {
	var s strings.Builder
	fmt.Fprint(&s, string(b.before))
	fmt.Fprint(&s, string(b.edit))
	fmt.Fprint(&s, string(b.after))
	return s.String()
}
