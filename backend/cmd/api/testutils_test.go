package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/BlokOfWood/CSK4-szamtan-gyakorlo/backend/internal/data"
)

func newTestApplication(t *testing.T) *application {
	var cfg = config{
		port: 8080,
		env:  "test",
		db: struct{ path string }{
			path: path.Join(t.TempDir(), "test.db"),
		},
	}

	db, err := initDB(cfg)
	if err != nil {
		t.Fatal(err)
	}

	return &application{
		config: cfg,
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
		models: data.NewModels(db),
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)
	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}

func (ts *testServer) post(t *testing.T, urlPath, body string) (int, http.Header, string) {
	rs, err := ts.Client().Post(ts.URL+urlPath, "application/json", bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	resBody, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	resBody = bytes.TrimSpace(resBody)

	return rs.StatusCode, rs.Header, string(resBody)
}
