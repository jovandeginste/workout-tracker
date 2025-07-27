package partials

type PageOptions struct {
	Scripts, Styles []string
}

func NewPageOptions() *PageOptions {
	return &PageOptions{Scripts: nil, Styles: nil}
}

func (o *PageOptions) WithScripts(scripts ...string) *PageOptions {
	return &PageOptions{
		Scripts: append(o.Scripts, scripts...),
		Styles:  o.Styles,
	}
}

func (o *PageOptions) WithStyles(styles ...string) *PageOptions {
	return &PageOptions{
		Scripts: o.Scripts,
		Styles:  append(o.Styles, styles...),
	}
}

func (o *PageOptions) WithSharing() *PageOptions {
	return o.
		WithScripts("/vendor/shareon/shareon.iife.js").
		WithStyles("/vendor/shareon/shareon.min.css")
}
