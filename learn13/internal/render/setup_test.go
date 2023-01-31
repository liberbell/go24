package render

import (
	"os"
	"testing"

	"github.com/alexedwards/scs/v2"
	"github.com/tsawler/bookings/internal/config"
)

var session *scs.SessionManager
var testApp config.AppConfig

func testMain(m *testing.M) {

	os.Exit(m.Run())
}
