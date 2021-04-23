package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollection_Append(t *testing.T) {
	t.Parallel()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		c := &Collection{}
		c.Append(&Cell{Row: 1, Col: 1})

		assert.Equal(t, &Collection{&Cell{Row: 1, Col: 1}}, c)
	})

	t.Run("OkWithTheSame", func(t *testing.T) {
		t.Parallel()

		c := &Collection{}
		c.Append(&Cell{Row: 1, Col: 1})
		c.Append(&Cell{Row: 1, Col: 1})

		assert.Equal(t, &Collection{&Cell{Row: 1, Col: 1}}, c)
	})
}

func TestCollection_Exists(t *testing.T) {
	t.Parallel()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		c := &Collection{&Cell{Row: 1, Col: 1}}

		assert.True(t, c.Exists(&Cell{Row: 1, Col: 1}))
	})

	t.Run("DoesNotExist", func(t *testing.T) {
		t.Parallel()

		c := &Collection{&Cell{Row: 1, Col: 1}}

		assert.False(t, c.Exists(&Cell{Row: 1, Col: 2}))
	})
}
