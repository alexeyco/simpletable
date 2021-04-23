package simpletable_test

import (
	"testing"

	st "github.com/alexeyco/simpletable"
	"github.com/stretchr/testify/assert"
)

func TestWithoutMargin(t *testing.T) {
	t.Parallel()

	var o st.Options
	st.WithoutMargin(&o)

	assert.True(t, o.WithoutMargin)
}

func TestAlign(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		name  string
		align uint8
	}{
		{
			name:  "Left",
			align: st.Left,
		},
		{
			name:  "Center",
			align: st.Center,
		},
		{
			name:  "Right",
			align: st.Right,
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			var o st.Options
			st.Align(testDatum.align)(&o)

			assert.Equal(t, testDatum.align, o.Align)
		})
	}
}

func TestSpan(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		name string
		span uint
	}{
		{
			name: "1",
			span: 1,
		},
		{
			name: "12",
			span: 12,
		},
		{
			name: "123",
			span: 123,
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			var o st.Options
			st.Span(testDatum.span)(&o)

			assert.Equal(t, int(testDatum.span), o.Span)
		})
	}
}

func TestWidth(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		name  string
		width uint
	}{
		{
			name:  "1",
			width: 1,
		},
		{
			name:  "12",
			width: 12,
		},
		{
			name:  "123",
			width: 123,
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			var o st.Options
			st.Width(testDatum.width)(&o)

			assert.Equal(t, int(testDatum.width), o.Width)
		})
	}
}
