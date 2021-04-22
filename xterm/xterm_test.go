package xterm_test

import (
	"testing"

	"github.com/alexeyco/simpletable/xterm"
	"github.com/stretchr/testify/assert"
)

func TestStrip(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		name     string
		source   string
		expected string
	}{
		{
			name:     "Escape color",
			source:   "\u001B[38;5;140mLorem ipsum dolor sit amet\u001B[0m",
			expected: "Lorem ipsum dolor sit amet",
		},
		{
			name:     "Escape color",
			source:   "\u001B[00;38;5;244m\u001B[m\u001B[00;38;5;33mLorem ipsum dolor sit amet\u001B[0m",
			expected: "Lorem ipsum dolor sit amet",
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			actual := xterm.Strip(testDatum.source)

			assert.Equal(t, testDatum.expected, actual)
		})
	}
}
