package copyheaders

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCopy(t *testing.T) {
	cfg := CreateConfig()
	cfg.Headers = append(cfg.Headers, HeaderConfig{
		From:   "from",
		To:     "to",
		Prefix: "prefix ",
	})

	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	handler, err := New(context.Background(), next, cfg, "demo-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 1. SIMULATE THE INCOMING HEADER
	req.Header.Set("from", "abc")

	// 2. RUN THE MIDDLEWARE
	handler.ServeHTTP(recorder, req)

	// 3. ASSERT AGAINST THE REQUEST (req), NOT THE RECORDER
	expected := "prefix abc"
	actual := req.Header.Get("to")
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
