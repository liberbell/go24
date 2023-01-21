package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

func (f *Form) valid() bool {
	return len(f.Errors) == 0
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string, r http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}
