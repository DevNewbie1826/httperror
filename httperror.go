package httperror

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// --- Context-based Error Reporter Middleware ---

// contextKey is a private, zero-sized type used for context keys.
// Using an empty struct provides the strongest guarantee of key uniqueness.
// contextKey는 컨텍스트 키로 사용되는 비공개 제로 사이즈 타입입니다.
// 빈 구조체를 사용하면 키의 고유성을 가장 강력하게 보장할 수 있습니다.
type contextKey struct{}

// errorKey is the context key used to store the error reporting function.
// errorKey는 오류 보고 함수를 저장하는 데 사용되는 컨텍스트 키입니다.
var errorKey = contextKey{}

// ReportError retrieves the error reporting function from the context and calls it.
// If the reporter function is not found, it does nothing.
// ReportError는 컨텍스트에서 오류 보고 함수를 검색하여 호출합니다.
// 보고 함수를 찾지 못하면 아무 작업도 수행하지 않습니다.
func ReportError(r *http.Request, err error) {
	if reporter, ok := r.Context().Value(errorKey).(func(error)); ok {
		reporter(err)
	}
}

// ErrorHandler defines the function signature for custom error handlers.
// ErrorHandler는 사용자 정의 오류 핸들러를 위한 함수 시그니처를 정의합니다.
type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

// DefaultErrorHandler provides a default implementation for handling errors.
// It checks if the error is an HttpError and writes the appropriate JSON response.
// For any other error, it returns a 500 Internal Server Error.
// DefaultErrorHandler는 오류 처리를 위한 기본 구현을 제공합니다.
// 오류가 HttpError인지 확인하고 적절한 JSON 응답을 작성합니다.
// 다른 모든 오류에 대해서는 500 내부 서버 오류를 반환합니다.
func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Check if the error is a known HttpError.
	if httpErr, ok := err.(*HttpError); ok {
		w.WriteHeader(httpErr.Status)
		json.NewEncoder(w).Encode(httpErr)
		return
	}

	// For any other error, return a 500 Internal Server Error.
	// It's a good practice to log the actual error here.
	internalErr := InternalServerError() // Using the helper from httperror_helpers.go
	w.WriteHeader(internalErr.Status)
	json.NewEncoder(w).Encode(internalErr)
}

// responseWriter is a wrapper around http.ResponseWriter that tracks whether
// the response has been written to. It also implements io.ReaderFrom to support
// zero-copy file transfers (sendfile).
type responseWriter struct {
	http.ResponseWriter
	wroteHeader bool
	code        int
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}
	rw.code = code
	rw.wroteHeader = true
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.wroteHeader {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(b)
}

// ReadFrom implements io.ReaderFrom to allow optimized sendfile system calls.
func (rw *responseWriter) ReadFrom(r io.Reader) (n int64, err error) {
	if !rw.wroteHeader {
		rw.WriteHeader(http.StatusOK)
	}
	if rf, ok := rw.ResponseWriter.(io.ReaderFrom); ok {
		return rf.ReadFrom(r)
	}
	return io.Copy(rw.ResponseWriter, r)
}

// Unwrap returns the original http.ResponseWriter.
// This is useful for http.ResponseController and other tools that need to access the underlying writer.
func (rw *responseWriter) Unwrap() http.ResponseWriter {
	return rw.ResponseWriter
}

// Flush implements http.Flusher to allow flushing buffered data to the client.
func (rw *responseWriter) Flush() {
	if f, ok := rw.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

// NewErrorReporterMiddleware creates a new error reporting middleware.
// It takes an optional ErrorHandler. If no handler is provided, DefaultErrorHandler is used.
// NewErrorReporterMiddleware는 새로운 오류 보고 미들웨어를 생성합니다.
// 선택적으로 ErrorHandler를 인자로 받으며, 핸들러가 제공되지 않으면 DefaultErrorHandler가 사용됩니다.
func NewErrorReporterMiddleware(handler ErrorHandler) func(http.Handler) http.Handler {
	if handler == nil {
		handler = DefaultErrorHandler
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var handlerError error
			reportError := func(err error) {
				handlerError = err
			}
			ctx := context.WithValue(r.Context(), errorKey, reportError)

			// Wrap the ResponseWriter to track if the response has been committed.
			rw := &responseWriter{ResponseWriter: w}

			// Call the next handler with the wrapped writer and the new context.
			next.ServeHTTP(rw, r.WithContext(ctx))

			// After the handler has run, check if an error was reported.
			if handlerError != nil {
				// If the response header has already been written, we cannot change the status code or body.
				// In this case, we log the error (if a logger were available) and do nothing,
				// as the response is already on its way to the client.
				if rw.wroteHeader {
					return
				}

				// If the response hasn't been committed yet, we can use the error handler
				// to send a proper error response.
				handler(w, r, handlerError)
			}
		})
	}
}

// ErrorReporterMiddleware provides a centralized error handling mechanism using the default handler.
// For custom error handling, use NewErrorReporterMiddleware.
// ErrorReporterMiddleware는 기본 핸들러를 사용하여 중앙 집중식 오류 처리 메커니즘을 제공합니다.
// 사용자 정의 오류 처리를 위해서는 NewErrorReporterMiddleware를 사용하십시오.
func ErrorReporterMiddleware(next http.Handler) http.Handler {
	// Maintain backward compatibility by using the new constructor with the default handler.
	return NewErrorReporterMiddleware(nil)(next)
}
