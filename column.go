package simpletable

import (
	"fmt"
	"unicode/utf8"
)

const (
	AlignLeft   = -1
	AlignCenter = 0
	AlignRight  = 1
)

type Column struct {
	Align   int
	Span    int
	Content interface{}
}

func (c *Column) Width() int {
	return utf8.RuneCountInString(c.String()) / c.Span
}

func (c *Column) String() string {
	switch v := c.Content.(type) {
	case string:
		return v
	case int, int32, int64, uint, uint32, uint64:
		return fmt.Sprintf("%d", v)
	}

	return fmt.Sprintf("%v", c.Content)
}
