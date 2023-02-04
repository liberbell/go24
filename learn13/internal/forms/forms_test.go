package forms

import (
	"net/http/httptest"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	is_valid := form.Valid()
	if !is_valid {
		t.Error("got invalid when shoud have benn valid")
	}
}
