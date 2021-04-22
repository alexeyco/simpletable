package simpletable

const (
	Left = iota
	Center
	Right
)

type Options struct {
	Align uint8
	Span  int
	Width int
}

type Option func(*Options)

func Align(align uint8) Option {
	return func(o *Options) {
		o.Align = align
	}
}

func Width(width uint) Option {
	return func(o *Options) {
		o.Width = int(width)
	}
}
