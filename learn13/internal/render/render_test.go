package render

import (
	"github.com/tsawler/bookings/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(error)
	}
	result := AddDefaultData(&td, r)

	if result == nil {
		t.Error("failed to add default data.")
	}
}
