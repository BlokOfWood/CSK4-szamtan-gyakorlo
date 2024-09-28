package main

import (
	"net/http"
	"testing"

	"github.com/BlokOfWood/CSK4-szamtan-gyakorlo/backend/internal/assert"
)

func TestEnableCORS(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.enableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})))
	defer ts.Close()

	statusCode, headers, _ := ts.get(t, "/v1/healthcheck")

	assert.Equal(t, statusCode, http.StatusOK)
	assert.Equal(t, headers.Get("Access-Control-Allow-Origin"), "*")
}
