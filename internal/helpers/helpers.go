package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/sinakovs/bookings/internal/config"
)

var app *config.AppConfig

// NewHalpers sets up app config for helpers
func NewHalpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
