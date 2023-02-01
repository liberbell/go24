package render

import (
	"github.com/tsawler/bookings/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r.err = getSession()
	if err != nil {
		t.Error(err)
	}

	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(error)
	}
	result := AddDefaultData(&td, r)

	if result == nil {
		t.Error("failed to add default data.")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}
