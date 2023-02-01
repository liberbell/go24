package render

import (
	"github.com/tsawler/bookings/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := http.NewRequest("GET", "/some-url", nil)
}
