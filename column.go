package simpletable

type Column struct {
	content *Content
}

func (c *Column) Content() *Content {
	return c.content
}
