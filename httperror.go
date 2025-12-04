package httperror

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// ErrorHandler defines the function signature for custom error handlers.
// ErrorHandler는 사용자 정의 오류 핸들러를 위한 함수 시그니처를 정의합니다.
type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

// currentErrorHandler stores the currently active error handler.
// Defaults to DefaultErrorHandler.
var currentErrorHandler ErrorHandler = DefaultErrorHandler

// SetErrorHandler sets the global error handler.
// If nil is provided, it sets the handler to DefaultErrorHandler.
// SetErrorHandler는 전역 오류 핸들러를 설정합니다. nil이 제공되면 기본 핸들러로 설정됩니다.
func SetErrorHandler(handler ErrorHandler) {
	if handler == nil {
		currentErrorHandler = DefaultErrorHandler
	} else {
		currentErrorHandler = handler
	}
}

// Respond calls the globally configured error handler to handle the error.
// Respond는 설정된 전역 오류 핸들러를 호출하여 오류를 처리합니다.
func Respond(w http.ResponseWriter, r *http.Request, err error) {
	currentErrorHandler(w, r, err)
}

// DefaultErrorHandler provides a default implementation for handling errors.
// It checks if the error is an HttpError and writes the appropriate JSON or HTML response
// based on the Request's Accept header.
// For any other error, it returns a 500 Internal Server Error.
// DefaultErrorHandler는 오류 처리를 위한 기본 구현을 제공합니다.
// 오류가 HttpError인지 확인하고 요청의 Accept 헤더에 따라 적절한 JSON 또는 HTML 응답을 작성합니다.
// 다른 모든 오류에 대해서는 500 내부 서버 오류를 반환합니다.
func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	// Simple Content Negotiation:
	accept := r.Header.Get("Accept")
	useHTML := false
	if accept != "" {
		if strings.Contains(accept, "text/html") || strings.Contains(accept, "application/xhtml+xml") {
			useHTML = true
		}
	}

	// Ensure we are dealing with an HttpError
	var httpErr *HttpError
	if e, ok := err.(*HttpError); ok && e != nil {
		httpErr = e
	} else {
		httpErr = InternalServerErrorError()
	}

	// Header MUST be set before WriteHeader
	if useHTML {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(httpErr.Status)
		io.WriteString(w, `<div class="http-error">`+httpErr.Message+`</div>`)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(httpErr.Status)
		json.NewEncoder(w).Encode(httpErr)
	}
}
