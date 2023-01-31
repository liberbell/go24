package render

import (
	"github.com/alexedwards/scs/v2"
	"github.com/tsawler/bookings/internal/config"
)

var session *scs.SessionManager
var testApp config.AppConfig
