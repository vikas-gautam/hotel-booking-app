package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/vikas-gautam/hotel-booking-app/internal/config"
)

func TestRoutes(t *testing.T) {

	var app config.AppConfig

	r := routes(&app)

	switch v := r.(type) {
	case *chi.Mux:
		//do nothing, test passed
	default:
		t.Errorf(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}

}
