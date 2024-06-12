package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/macropusgiganteus/covid-summarizer-api/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, "pong", w.Body.String())
}
