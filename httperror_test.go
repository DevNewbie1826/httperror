package httperror

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestRespond tests the Respond function and SetErrorHandler.
func TestRespond(t *testing.T) {
	// Reset handler to default before test
	SetErrorHandler(nil)

	t.Run("default handler", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		err := New(http.StatusBadRequest, "Bad Request")

		Respond(rr, req, err)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
		}
		if rr.Header().Get("Content-Type") != "application/json; charset=utf-8" {
			t.Errorf("expected Content-Type application/json; charset=utf-8, got %s", rr.Header().Get("Content-Type"))
		}
	})

	t.Run("custom handler", func(t *testing.T) {
		customHandler := func(w http.ResponseWriter, r *http.Request, err error) {
			w.WriteHeader(http.StatusTeapot)
			w.Write([]byte("I am a teapot"))
		}
		SetErrorHandler(customHandler)
		defer SetErrorHandler(nil) // Cleanup

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		err := errors.New("some error")

		Respond(rr, req, err)

		if rr.Code != http.StatusTeapot {
			t.Errorf("expected status %d, got %d", http.StatusTeapot, rr.Code)
		}
		if rr.Body.String() != "I am a teapot" {
			t.Errorf("expected body 'I am a teapot', got '%s'", rr.Body.String())
		}
	})
}

// TestDefaultErrorHandler tests the DefaultErrorHandler function.
func TestDefaultErrorHandler(t *testing.T) {
	t.Run("with HttpError and JSON default", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		err := &HttpError{Status: http.StatusNotFound, Message: "Not Found"}

		DefaultErrorHandler(rr, req, err)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, rr.Code)
		}
		if rr.Header().Get("Content-Type") != "application/json; charset=utf-8" {
			t.Errorf("expected content type application/json, got %s", rr.Header().Get("Content-Type"))
		}

		var body HttpError
		if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
			t.Fatalf("could not decode response body: %v", err)
		}
		if body.Status != err.Status || body.Message != err.Message {
			t.Errorf("expected body %+v, got %+v", err, body)
		}
	})

	t.Run("with HttpError and HTML accept", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Accept", "text/html")
		err := &HttpError{Status: http.StatusForbidden, Message: "Forbidden"}

		DefaultErrorHandler(rr, req, err)

		if rr.Code != http.StatusForbidden {
			t.Errorf("expected status %d, got %d", http.StatusForbidden, rr.Code)
		}
		if rr.Header().Get("Content-Type") != "text/html; charset=utf-8" {
			t.Errorf("expected content type text/html, got %s", rr.Header().Get("Content-Type"))
		}
		expectedBody := `<div class="http-error">Forbidden</div>`
		if rr.Body.String() != expectedBody {
			t.Errorf("expected body '%s', got '%s'", expectedBody, rr.Body.String())
		}
	})

	t.Run("with generic error", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		err := errors.New("some internal error")

		DefaultErrorHandler(rr, req, err)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rr.Code)
		}

		var body HttpError
		if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
			t.Fatalf("could not decode response body: %v", err)
		}
		expectedErr := InternalServerErrorError()
		if body.Status != expectedErr.Status {
			t.Errorf("expected status %d, got %d", expectedErr.Status, body.Status)
		}
	})
}

func TestSetErrorHandler(t *testing.T) {
	// Ensure it resets to default when nil is passed
	SetErrorHandler(nil)
	// Check by address if possible, or behavior. Go doesn't allow comparing func pointers easily.
	// We'll test behavior.
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	Respond(rr, req, New(http.StatusOK, "ok"))
	if rr.Header().Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("Default handler should set JSON content type")
	}

	// Set custom
	SetErrorHandler(func(w http.ResponseWriter, r *http.Request, err error) {
		w.Write([]byte("custom"))
	})
	rr = httptest.NewRecorder()
	Respond(rr, req, errors.New("err"))
	if rr.Body.String() != "custom" {
		t.Error("Custom handler didn't work")
	}

	// Reset to default
	SetErrorHandler(nil)
	rr = httptest.NewRecorder()
	Respond(rr, req, New(http.StatusOK, "ok"))
	if rr.Header().Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("Should revert to default handler")
	}
}