package simpletable

const (
	Left = iota
	Center
	Right
)

type Options struct {
	Align uint8
	Span  uint
	Width uint
}

type Option func(*Options)

func Align(align uint8) Option {
	return func(o *Options) {
		o.Align = align
	}
}

func Span(span uint) Option {
	return func(o *Options) {
		o.Span = span
	}
}

func Width(width uint) Option {
	return func(o *Options) {
		o.Width = width
	}
}
