package httperror

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestReportError tests the ReportError function.
func TestReportError(t *testing.T) {
	t.Run("with reporter", func(t *testing.T) {
		var reportedError error
		reporter := func(err error) {
			reportedError = err
		}

		req := httptest.NewRequest("GET", "/", nil)
		ctx := req.Context()
		ctx = context.WithValue(ctx, errorKey, reporter)
		req = req.WithContext(ctx)

		err := errors.New("test error")
		ReportError(req, err)

		if reportedError == nil {
			t.Error("expected error to be reported, but it was nil")
		}
		if reportedError.Error() != "test error" {
			t.Errorf("expected error message 'test error', got '%s'", reportedError.Error())
		}
	})

	t.Run("without reporter", func(t *testing.T) {
		// This test ensures that ReportError does not panic when no reporter is in the context.
		req := httptest.NewRequest("GET", "/", nil)
		err := errors.New("test error")
		// We just call it. If it doesn't panic, the test passes.
		ReportError(req, err)
	})
}

// TestDefaultErrorHandler tests the DefaultErrorHandler function.
func TestDefaultErrorHandler(t *testing.T) {
	t.Run("with HttpError", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		err := &HttpError{Status: http.StatusNotFound, Message: "Not Found"}

		DefaultErrorHandler(rr, req, err)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, rr.Code)
		}

		var body HttpError
		if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
			t.Fatalf("could not decode response body: %v", err)
		}
		if body.Status != err.Status || body.Message != err.Message {
			t.Errorf("expected body %+v, got %+v", err, body)
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
		expectedErr := InternalServerError()
		if body.Status != expectedErr.Status || body.Message != expectedErr.Message {
			t.Errorf("expected body %+v, got %+v", expectedErr, body)
		}
	})
}

// TestNewErrorReporterMiddleware tests the NewErrorReporterMiddleware.
func TestNewErrorReporterMiddleware(t *testing.T) {
	// Handler that reports an error but does NOT write to the response.
	// This should trigger the error handler.
	errorOnlyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ReportError(r, errors.New("handler error"))
	})

	t.Run("with default handler", func(t *testing.T) {
		middleware := NewErrorReporterMiddleware(nil)
		testHandler := middleware(errorOnlyHandler)

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		testHandler.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rr.Code)
		}
	})

	t.Run("with custom handler", func(t *testing.T) {
		customHandler := func(w http.ResponseWriter, r *http.Request, err error) {
			w.WriteHeader(http.StatusTeapot)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}
		middleware := NewErrorReporterMiddleware(customHandler)
		testHandler := middleware(errorOnlyHandler)

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		testHandler.ServeHTTP(rr, req)

		if rr.Code != http.StatusTeapot {
			t.Errorf("expected status %d, got %d", http.StatusTeapot, rr.Code)
		}
	})

	t.Run("no error reported", func(t *testing.T) {
		emptyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		middleware := NewErrorReporterMiddleware(nil)
		testHandler := middleware(emptyHandler)

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		testHandler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("error reported after response committed", func(t *testing.T) {
		// Handler that writes to the response AND reports an error.
		// Since we are now in pass-through mode, the error should be IGNORED (or logged),
		// and the original response should be preserved.
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
			ReportError(r, errors.New("handler error"))
		})

		middleware := NewErrorReporterMiddleware(nil)
		testHandler := middleware(handler)

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		testHandler.ServeHTTP(rr, req)

		// Should NOT be 500 Internal Server Error
		if rr.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
		}
		if rr.Body.String() != "ok" {
			t.Errorf("expected body 'ok', got '%s'", rr.Body.String())
		}
	})
}

// TestErrorReporterMiddleware tests the backward-compatible ErrorReporterMiddleware.
func TestErrorReporterMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ReportError(r, BadRequest("test bad request"))
		// Do NOT write to w here, so the middleware can handle the error.
	})

	testHandler := ErrorReporterMiddleware(handler)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	testHandler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}

	var body HttpError
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}
	if body.Message != "test bad request" {
		t.Errorf("expected message 'test bad request', got '%s'", body.Message)
	}
}

func TestNilHttpErrorPanic(t *testing.T) {
	t.Run("Middleware", func(t *testing.T) {
		// Create a handler that reports a typed nil error
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var err *HttpError = nil
			ReportError(r, err)
		})

		// Wrap with middleware
		middleware := ErrorReporterMiddleware(handler)

		// Create a request
		req := httptest.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()

		// This should not panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("The code panicked: %v", r)
			}
		}()

		middleware.ServeHTTP(rr, req)
	})

	t.Run("DefaultErrorHandler", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()
		var err *HttpError = nil

		// This should not panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("The code panicked: %v", r)
			}
		}()

		DefaultErrorHandler(rr, req, err)
	})
}
