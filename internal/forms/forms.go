package forms

import (
	"net/http"
	"net/url"
)


type Form struct{
	url.Values
	Errors errors
}

//New initializes a form struct
func New(data url.Values) *Form{
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Has checks if fo(rm field is in post and non empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == ""{
		return false
	}
	return true
}