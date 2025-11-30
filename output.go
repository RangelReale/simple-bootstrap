package simple_bootstrap

import (
	"errors"
	"fmt"
	"net/http"
)

type Output struct {
	w   http.ResponseWriter
	err error

	start     bool
	startData map[string]any

	createContainer bool
	containerClass  string
	bootstrapCSSURL string
	bootstrapJSURL  string

	style  string
	script string
}

func NewOutput(w http.ResponseWriter, options ...Option) *Output {
	ret := &Output{
		w:               w,
		createContainer: true,
		containerClass:  "container",
		bootstrapCSSURL: `https://cdn.jsdelivr.net/npm/bootstrap@5.3.7/dist/css/bootstrap.min.css`,
		bootstrapJSURL:  `https://cdn.jsdelivr.net/npm/bootstrap@5.3.7/dist/js/bootstrap.bundle.min.js`,
	}
	for _, option := range options {
		option(ret)
	}
	if ret.start {
		ret.Start(ret.startData)
	}
	return ret
}

func (w *Output) Start(data map[string]any) {
	beginData := map[string]any{
		"title":           "Title",
		"createContainer": w.createContainer,
		"containerClass":  w.containerClass,
		"bootstrapCSSURL": w.bootstrapCSSURL,
	}
	for k, v := range data {
		beginData[k] = v
	}

	w.w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.CheckError(Template(w.w, "layout_begin", beginData))

	if w.style != "" {
		w.Write(`
	<style>
	%s
	</style>
	`, w.style)
	}
	if w.script != "" {
		w.Write(`
	<script>
	%s
	</script>
	`, w.script)
	}
}

func (w *Output) End(data map[string]any) {
	endData := map[string]any{
		"bootstrapJSURL":  w.bootstrapJSURL,
		"createContainer": w.createContainer,
	}
	for k, v := range data {
		endData[k] = v
	}

	w.CheckError(Template(w.w, "layout_end", endData))
}

func (w *Output) Err() error {
	if w.err != nil {
		return w.err
	}
	return w.err
}

func (w *Output) CheckError(err error) {
	if err != nil {
		w.err = errors.Join(w.err, err)
	}
}

func (w *Output) Write(s string, a ...interface{}) {
	var err error
	if len(a) > 0 {
		_, err = w.w.Write([]byte(fmt.Sprintf(s, a...)))
	} else {
		_, err = w.w.Write([]byte(s))
	}
	w.CheckError(err)
}

type Option func(options *Output)

func WithStart(start bool, data map[string]any) Option {
	return func(options *Output) {
		options.start = start
		options.startData = data
	}
}

func WithCreateContainer(createContainer bool) Option {
	return func(options *Output) {
		options.createContainer = createContainer
	}
}

func WithContainerClass(containerClass string) Option {
	return func(options *Output) {
		options.containerClass = containerClass
	}
}

func WithStyle(style string) Option {
	return func(options *Output) {
		options.style = style
	}
}

func WithScript(script string) Option {
	return func(options *Output) {
		options.script = script
	}
}
