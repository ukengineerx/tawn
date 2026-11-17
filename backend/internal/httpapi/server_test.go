package httpapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	"tawn/backend/internal/notes"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func TestHealthDoesNotNeedDatabase(t *testing.T) {
	response := request(t, New(nil), http.MethodGet, "/health", nil)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d", response.Code)
	}
}

func TestReadyWithoutDatabase(t *testing.T) {
	response := request(t, New(nil), http.MethodGet, "/ready", nil)
	if response.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d", response.Code)
	}
}

func TestNotesRoundTrip(t *testing.T) {
	router := New(notes.NewMemory())

	created := request(t, router, http.MethodPost, "/v1/notes", []byte(`{"body":"hello"}`))
	if created.Code != http.StatusCreated {
		t.Fatalf("create status = %d body = %s", created.Code, created.Body.String())
	}

	listed := request(t, router, http.MethodGet, "/v1/notes", nil)
	if listed.Code != http.StatusOK {
		t.Fatalf("list status = %d", listed.Code)
	}
	var payload struct {
		Notes []notes.Note `json:"notes"`
	}
	if err := json.Unmarshal(listed.Body.Bytes(), &payload); err != nil {
		t.Fatal(err)
	}
	if len(payload.Notes) != 1 || payload.Notes[0].Body != "hello" {
		t.Fatalf("notes = %#v", payload.Notes)
	}
}

func TestRejectsEmptyNote(t *testing.T) {
	response := request(t, New(notes.NewMemory()), http.MethodPost, "/v1/notes", []byte(`{"body":"  "}`))
	if response.Code != http.StatusBadRequest {
		t.Fatalf("status = %d", response.Code)
	}
}

func request(t *testing.T, router http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(method, path, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}
