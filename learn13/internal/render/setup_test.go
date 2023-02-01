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

	gob.Register(models.Reservation{})
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}
