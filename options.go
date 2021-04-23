package simpletable

const (
	Left = iota
	Center
	Right
)

type Options struct {
	WithoutMargin bool
	Align         uint8
	Span          int
	Width         int
}

type Option func(*Options)

func WithoutMargin(o *Options) {
	o.WithoutMargin = true
}

func Align(align uint8) Option {
	return func(o *Options) {
		o.Align = align
	}
}

func Span(span uint) Option {
	return func(o *Options) {
		o.Span = int(span)
	}
}

func Width(width uint) Option {
	return func(o *Options) {
		o.Width = int(width)
	}
}
