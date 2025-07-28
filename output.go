package simple_bootstrap

import (
	"errors"
	"fmt"
	"net/http"
)

type Output struct {
	w   http.ResponseWriter
	err error
}

func NewOutput(w http.ResponseWriter) *Output {
	return &Output{w: w}
}

func NewOutputStart(w http.ResponseWriter, options ...StartOption) *Output {
	ret := NewOutput(w)
	ret.Start(options...)
	return ret
}

func (w *Output) Start(options ...StartOption) {
	var optns startOptions
	for _, opt := range options {
		opt(&optns)
	}

	w.w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.CheckError(Template(w.w, "layout_begin", optns.data))

	if optns.style != "" {
		w.Write(`
	<style>
	%s
	</style>
	`, optns.style)

	}
}

func (w *Output) End(data any) {
	w.CheckError(Template(w.w, "layout_end", data))
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
	_, err := w.w.Write([]byte(fmt.Sprintf(s, a...)))
	w.CheckError(err)
}

type StartOption func(options *startOptions)

func WithStartData(data any) StartOption {
	return func(options *startOptions) {
		options.data = data
	}
}

func WithStartStyle(style string) StartOption {
	return func(options *startOptions) {
		options.style = style
	}
}

type startOptions struct {
	data  any
	style string
}
