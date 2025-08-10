package httperror

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// TestHttpError_Error tests the Error method of the HttpError struct.
func TestHttpError_Error(t *testing.T) {
	err := HttpError{Message: "error message"}
	if err.Error() != "error message" {
		t.Errorf("expected 'error message', got '%s'", err.Error())
	}
}

// TestNew tests the New function.
func TestNew(t *testing.T) {
	err := New(http.StatusTeapot, "I'm a teapot")
	expected := &HttpError{Status: http.StatusTeapot, Message: "I'm a teapot"}
	if !reflect.DeepEqual(err, expected) {
		t.Errorf("expected %+v, got %+v", expected, err)
	}
}

// TestHelperFunctionsWithMessage tests all the helper functions that create HttpErrors.
func TestHelperFunctionsWithMessage(t *testing.T) {
	testCases := []struct {
		name     string
		function func(...string) *HttpError
		status   int
		message  string
	}{
		{"BadRequest", BadRequest, http.StatusBadRequest, "custom bad request"},
		{"Unauthorized", Unauthorized, http.StatusUnauthorized, "custom unauthorized"},
		{"PaymentRequired", PaymentRequired, http.StatusPaymentRequired, "custom payment required"},
		{"Forbidden", Forbidden, http.StatusForbidden, "custom forbidden"},
		{"NotFound", NotFound, http.StatusNotFound, "custom not found"},
		{"MethodNotAllowed", MethodNotAllowed, http.StatusMethodNotAllowed, "custom method not allowed"},
		{"NotAcceptable", NotAcceptable, http.StatusNotAcceptable, "custom not acceptable"},
		{"ProxyAuthRequired", ProxyAuthRequired, http.StatusProxyAuthRequired, "custom proxy auth required"},
		{"RequestTimeout", RequestTimeout, http.StatusRequestTimeout, "custom request timeout"},
		{"Conflict", Conflict, http.StatusConflict, "custom conflict"},
		{"Gone", Gone, http.StatusGone, "custom gone"},
		{"LengthRequired", LengthRequired, http.StatusLengthRequired, "custom length required"},
		{"PreconditionFailed", PreconditionFailed, http.StatusPreconditionFailed, "custom precondition failed"},
		{"PayloadTooLarge", PayloadTooLarge, http.StatusRequestEntityTooLarge, "custom payload too large"},
		{"URITooLong", URITooLong, http.StatusRequestURITooLong, "custom uri too long"},
		{"UnsupportedMediaType", UnsupportedMediaType, http.StatusUnsupportedMediaType, "custom unsupported media type"},
		{"RangeNotSatisfiable", RangeNotSatisfiable, http.StatusRequestedRangeNotSatisfiable, "custom range not satisfiable"},
		{"ExpectationFailed", ExpectationFailed, http.StatusExpectationFailed, "custom expectation failed"},
		{"Teapot", Teapot, http.StatusTeapot, "custom teapot"},
		{"MisdirectedRequest", MisdirectedRequest, http.StatusMisdirectedRequest, "custom misdirected request"},
		{"UnprocessableEntity", UnprocessableEntity, http.StatusUnprocessableEntity, "custom unprocessable entity"},
		{"Locked", Locked, http.StatusLocked, "custom locked"},
		{"FailedDependency", FailedDependency, http.StatusFailedDependency, "custom failed dependency"},
		{"TooEarly", TooEarly, http.StatusTooEarly, "custom too early"},
		{"UpgradeRequired", UpgradeRequired, http.StatusUpgradeRequired, "custom upgrade required"},
		{"PreconditionRequired", PreconditionRequired, http.StatusPreconditionRequired, "custom precondition required"},
		{"TooManyRequests", TooManyRequests, http.StatusTooManyRequests, "custom too many requests"},
		{"RequestHeaderFieldsTooLarge", RequestHeaderFieldsTooLarge, http.StatusRequestHeaderFieldsTooLarge, "custom request header fields too large"},
		{"UnavailableForLegalReasons", UnavailableForLegalReasons, http.StatusUnavailableForLegalReasons, "custom unavailable for legal reasons"},
		{"InternalServerError", InternalServerError, http.StatusInternalServerError, "custom internal server error"},
		{"NotImplemented", NotImplemented, http.StatusNotImplemented, "custom not implemented"},
		{"BadGateway", BadGateway, http.StatusBadGateway, "custom bad gateway"},
		{"ServiceUnavailable", ServiceUnavailable, http.StatusServiceUnavailable, "custom service unavailable"},
		{"GatewayTimeout", GatewayTimeout, http.StatusGatewayTimeout, "custom gateway timeout"},
		{"HTTPVersionNotSupported", HTTPVersionNotSupported, http.StatusHTTPVersionNotSupported, "custom http version not supported"},
		{"VariantAlsoNegotiates", VariantAlsoNegotiates, http.StatusVariantAlsoNegotiates, "custom variant also negotiates"},
		{"InsufficientStorage", InsufficientStorage, http.StatusInsufficientStorage, "custom insufficient storage"},
		{"LoopDetected", LoopDetected, http.StatusLoopDetected, "custom loop detected"},
		{"NotExtended", NotExtended, http.StatusNotExtended, "custom not extended"},
		{"NetworkAuthenticationRequired", NetworkAuthenticationRequired, http.StatusNetworkAuthenticationRequired, "custom network authentication required"},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" with message", func(t *testing.T) {
			err := tc.function(tc.message)
			if err.Status != tc.status {
				t.Errorf("expected status %d, got %d", tc.status, err.Status)
			}
			if err.Message != tc.message {
				t.Errorf("expected message '%s', got '%s'", tc.message, err.Message)
			}
		})

		t.Run(tc.name+" without message", func(t *testing.T) {
			err := tc.function()
			expectedMessage := http.StatusText(tc.status)
			if err.Status != tc.status {
				t.Errorf("expected status %d, got %d", tc.status, err.Status)
			}
			if err.Message != expectedMessage {
				t.Errorf("expected message '%s', got '%s'", expectedMessage, err.Message)
			}
		})
	}
}

// TestReportHelperFunctions tests all the reporter helper functions.
func TestReportHelperFunctions(t *testing.T) {
	testCases := []struct {
		name           string
		reporterFunc   func(*http.Request, ...string)
		expectedStatus int
		customMessage  string
	}{
		{"ReportBadRequest", ReportBadRequest, http.StatusBadRequest, "custom bad request"},
		{"ReportUnauthorized", ReportUnauthorized, http.StatusUnauthorized, "custom unauthorized"},
		{"ReportPaymentRequired", ReportPaymentRequired, http.StatusPaymentRequired, "custom payment required"},
		{"ReportForbidden", ReportForbidden, http.StatusForbidden, "custom forbidden"},
		{"ReportNotFound", ReportNotFound, http.StatusNotFound, "custom not found"},
		{"ReportMethodNotAllowed", ReportMethodNotAllowed, http.StatusMethodNotAllowed, "custom method not allowed"},
		{"ReportNotAcceptable", ReportNotAcceptable, http.StatusNotAcceptable, "custom not acceptable"},
		{"ReportProxyAuthRequired", ReportProxyAuthRequired, http.StatusProxyAuthRequired, "custom proxy auth required"},
		{"ReportRequestTimeout", ReportRequestTimeout, http.StatusRequestTimeout, "custom request timeout"},
		{"ReportConflict", ReportConflict, http.StatusConflict, "custom conflict"},
		{"ReportGone", ReportGone, http.StatusGone, "custom gone"},
		{"ReportLengthRequired", ReportLengthRequired, http.StatusLengthRequired, "custom length required"},
		{"ReportPreconditionFailed", ReportPreconditionFailed, http.StatusPreconditionFailed, "custom precondition failed"},
		{"ReportPayloadTooLarge", ReportPayloadTooLarge, http.StatusRequestEntityTooLarge, "custom payload too large"},
		{"ReportURITooLong", ReportURITooLong, http.StatusRequestURITooLong, "custom uri too long"},
		{"ReportUnsupportedMediaType", ReportUnsupportedMediaType, http.StatusUnsupportedMediaType, "custom unsupported media type"},
		{"ReportRangeNotSatisfiable", ReportRangeNotSatisfiable, http.StatusRequestedRangeNotSatisfiable, "custom range not satisfiable"},
		{"ReportExpectationFailed", ReportExpectationFailed, http.StatusExpectationFailed, "custom expectation failed"},
		{"ReportTeapot", ReportTeapot, http.StatusTeapot, "custom teapot"},
		{"ReportMisdirectedRequest", ReportMisdirectedRequest, http.StatusMisdirectedRequest, "custom misdirected request"},
		{"ReportUnprocessableEntity", ReportUnprocessableEntity, http.StatusUnprocessableEntity, "custom unprocessable entity"},
		{"ReportLocked", ReportLocked, http.StatusLocked, "custom locked"},
		{"ReportFailedDependency", ReportFailedDependency, http.StatusFailedDependency, "custom failed dependency"},
		{"ReportTooEarly", ReportTooEarly, http.StatusTooEarly, "custom too early"},
		{"ReportUpgradeRequired", ReportUpgradeRequired, http.StatusUpgradeRequired, "custom upgrade required"},
		{"ReportPreconditionRequired", ReportPreconditionRequired, http.StatusPreconditionRequired, "custom precondition required"},
		{"ReportTooManyRequests", ReportTooManyRequests, http.StatusTooManyRequests, "custom too many requests"},
		{"ReportRequestHeaderFieldsTooLarge", ReportRequestHeaderFieldsTooLarge, http.StatusRequestHeaderFieldsTooLarge, "custom request header fields too large"},
		{"ReportUnavailableForLegalReasons", ReportUnavailableForLegalReasons, http.StatusUnavailableForLegalReasons, "custom unavailable for legal reasons"},
		{"ReportInternalServerError", ReportInternalServerError, http.StatusInternalServerError, "custom internal server error"},
		{"ReportNotImplemented", ReportNotImplemented, http.StatusNotImplemented, "custom not implemented"},
		{"ReportBadGateway", ReportBadGateway, http.StatusBadGateway, "custom bad gateway"},
		{"ReportServiceUnavailable", ReportServiceUnavailable, http.StatusServiceUnavailable, "custom service unavailable"},
		{"ReportGatewayTimeout", ReportGatewayTimeout, http.StatusGatewayTimeout, "custom gateway timeout"},
		{"ReportHTTPVersionNotSupported", ReportHTTPVersionNotSupported, http.StatusHTTPVersionNotSupported, "custom http version not supported"},
		{"ReportVariantAlsoNegotiates", ReportVariantAlsoNegotiates, http.StatusVariantAlsoNegotiates, "custom variant also negotiates"},
		{"ReportInsufficientStorage", ReportInsufficientStorage, http.StatusInsufficientStorage, "custom insufficient storage"},
		{"ReportLoopDetected", ReportLoopDetected, http.StatusLoopDetected, "custom loop detected"},
		{"ReportNotExtended", ReportNotExtended, http.StatusNotExtended, "custom not extended"},
		{"ReportNetworkAuthenticationRequired", ReportNetworkAuthenticationRequired, http.StatusNetworkAuthenticationRequired, "custom network authentication required"},
	}

	for _, tc := range testCases {
		// With custom message
		t.Run(fmt.Sprintf("%s with message", tc.name), func(t *testing.T) {
			var reportedError error
			reporter := func(err error) {
				reportedError = err
			}
			req := httptest.NewRequest("GET", "/", nil)
			ctx := context.WithValue(req.Context(), errorKey, reporter)
			req = req.WithContext(ctx)

			tc.reporterFunc(req, tc.customMessage)

			if reportedError == nil {
				t.Fatal("error was not reported")
			}
			httpErr, ok := reportedError.(*HttpError)
			if !ok {
				t.Fatalf("reported error is not of type HttpError: %T", reportedError)
			}
			if httpErr.Status != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, httpErr.Status)
			}
			if httpErr.Message != tc.customMessage {
				t.Errorf("expected message '%s', got '%s'", tc.customMessage, httpErr.Message)
			}
		})

		// Without custom message
		t.Run(fmt.Sprintf("%s without message", tc.name), func(t *testing.T) {
			var reportedError error
			reporter := func(err error) {
				reportedError = err
			}
			req := httptest.NewRequest("GET", "/", nil)
			ctx := context.WithValue(req.Context(), errorKey, reporter)
			req = req.WithContext(ctx)

			tc.reporterFunc(req)

			if reportedError == nil {
				t.Fatal("error was not reported")
			}
			httpErr, ok := reportedError.(*HttpError)
			if !ok {
				t.Fatalf("reported error is not of type HttpError: %T", reportedError)
			}
			expectedMessage := http.StatusText(tc.expectedStatus)
			if httpErr.Status != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, httpErr.Status)
			}
			if httpErr.Message != expectedMessage {
				t.Errorf("expected message '%s', got '%s'", expectedMessage, httpErr.Message)
			}
		})
	}
}
