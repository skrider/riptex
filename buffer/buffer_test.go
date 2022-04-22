package buffer

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test inserts <phrase> lines, <n> total times
func doTestInsert(t *testing.T, b *Buffer, phrase string, n int) {
	assert := assert.New(t)
	pre := []rune(b.String())
	cursor := b.cursor
	var ref strings.Builder
	fmt.Fprint(&ref, string(pre[:cursor]))
	runes := []rune(phrase)
	// a
	for i := 0; i < n; i++ {
		fmt.Fprint(&ref, phrase)
		b.Insert(runes)
	}
	fmt.Fprint(&ref, string(pre[cursor:]))
	final_ref := ref.String()
	final_buf := b.String()
	assert.Equal(strings.Compare(final_ref, final_buf), 0, "buffer should contain inserted content")
}

func TestInsertUnicode(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "Hello ⾗♒Ⲣ⹊⡮╫ \n", 10000)
}

func TestInsertOne(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 10000)
}

func TestInsertMany(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "aaa", 10000)
}

func doTestDelete(t *testing.T, b *Buffer, n int) {
	assert := assert.New(t)
	pre := []rune(b.String())
	cursor := b.cursor
	var ref strings.Builder
	if n > b.cursor {
		fmt.Fprint(&ref, string(pre))
	} else {
		fmt.Fprint(&ref, string(pre[:cursor-n]))
		fmt.Fprint(&ref, string(pre[cursor:]))
	}
	b.Delete(n)
	final_buf := b.String()
	final_ref := ref.String()
	assert.Equal(strings.Compare(final_buf, final_ref), 0, "n bytes should be gone")
}

func TestDelete(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 11)
	doTestDelete(t, &b, 10)
}

func TestDeleteAll(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 11)
	doTestDelete(t, &b, 11)
}

func TestDeleteTooMany(t *testing.T) {
	assert := assert.New(t)
	b := NewBuffer()
	doTestInsert(t, &b, "a", 1)
	pre := b.String()
	b.Delete(2)
	post := b.String()
	assert.Equal(strings.Compare(pre, post), 0)
}

func TestDeleteTwice(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 11)
	doTestDelete(t, &b, 3)
	doTestDelete(t, &b, 3)
}

func doTestMove(t *testing.T, b *Buffer, n int) {
	assert := assert.New(t)
	pre := b.String()
	b.MoveCursorRelative(n)
	post := b.String()
	assert.Equal(strings.Compare(pre, post), 0, "moving cursor should not affect buffer")
}

func TestMoveBack(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 10)
	doTestMove(t, &b, -5)
}

func TestMoveBackTwice(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 10)
	doTestMove(t, &b, -5)
	doTestMove(t, &b, -5)
}

func TestMoveBackInsert(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 10)
	doTestMove(t, &b, -5)
	doTestInsert(t, &b, "b", 10)
}

func TestMoveBackDelete(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 10)
	doTestMove(t, &b, -5)
	doTestDelete(t, &b, 5)
}

func TestMoveForward(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 20)
	doTestMove(t, &b, -10)
	doTestMove(t, &b, 5)
}

func TestMoveForwardInsert(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 20)
	doTestMove(t, &b, -10)
	doTestMove(t, &b, 5)
	doTestInsert(t, &b, "b", 10)
}

func TestMoveForwardDelete(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 20)
	doTestMove(t, &b, -10)
	doTestMove(t, &b, 5)
	doTestDelete(t, &b, 10)
}

func TestMoveOutOfBounds(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 20)
	doTestMove(t, &b, -10)
	doTestMove(t, &b, -11)
	doTestMove(t, &b, 11)
}

func TestIntegration(t *testing.T) {
	b := NewBuffer()
	doTestInsert(t, &b, "a", 30)
	doTestMove(t, &b, -30)
	doTestDelete(t, &b, 1)
	doTestInsert(t, &b, "b", 10)
	doTestMove(t, &b, -5)
	doTestInsert(t, &b, "c", 5)
	doTestDelete(t, &b, 5)
	doTestMove(t, &b, 35)
	doTestDelete(t, &b, 40)
	doTestInsert(t, &b, "d", 30)
	doTestMove(t, &b, -5)
	doTestDelete(t, &b, 10)
}
