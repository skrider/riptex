package buffer

import (
	"fmt"
	"strings"
)

// this represents the internal Gap Buffer data structure. Before
type Buffer struct {
	before []rune // contains [:cursor]
	edit   []rune // contains temporary text being edited by the cursor
	after  []rune // contains [cursor:]
	cursor int
}

func NewBuffer() Buffer {
	return Buffer{
		before: make([]rune, 0),
		edit:   make([]rune, 0),
		after:  make([]rune, 0),
		cursor: 0,
	}
}

func (b *Buffer) Insert(s []rune) {
	b.edit = append(b.edit, s...)
	b.cursor = b.cursor + len(s)
}

func (b *Buffer) Delete(n int) {
	if b.cursor-n >= 0 {
		if n > len(b.edit) {
			b.edit = make([]rune, 0)
			b.before = b.before[:len(b.before)+len(b.edit)-n]
		} else {
			b.edit = b.edit[:len(b.edit)-n]
		}
		b.cursor = b.cursor - n
	}
}

func (b *Buffer) MoveCursor(n int) {
	if n < 0 {
		n = 0
	} else if n > b.Len() {
		n = b.Len()
	}
	var before, after []rune
	if n < len(b.before) {
		after = make([]rune, 0)
		after = append(after, b.before[n:]...)
		after = append(after, b.edit...)
		after = append(after, b.after...)
		before = b.before[:n]
	} else if n >= len(b.before) && n < len(b.before)+len(b.edit) {
		before = append(b.before, b.edit[:n-len(b.before)]...)
		after = append(b.edit[n-len(b.before):], b.after...)
	} else {
		before = make([]rune, 0)
		before = append(before, b.before...)
		before = append(before, b.edit...)
		before = append(before, b.after[:n-len(b.before)-len(b.edit)]...)
		after = b.after[n-len(b.before)-len(b.edit):]
	}
	b.cursor = n
	b.before = before
	b.after = after
	b.edit = make([]rune, 0)
}

func (b *Buffer) Len() int {
	return len(b.before) + len(b.after) + len(b.edit)
}

func (b *Buffer) MoveCursorRelative(n int) {
	b.MoveCursor(b.cursor + n)
}

func (b *Buffer) Cursor() int {
	return b.cursor
}

func (b *Buffer) String() string {
	var s strings.Builder
	fmt.Fprint(&s, string(b.before))
	fmt.Fprint(&s, string(b.edit))
	fmt.Fprint(&s, string(b.after))
	return s.String()
}
