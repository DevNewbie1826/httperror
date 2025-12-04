package httperror

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
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

// TestHelperFunctions tests all the helper functions (BadRequest, NotFound, etc.).
func TestHelperFunctions(t *testing.T) {
	testCases := []struct {
		name           string
		function       func(http.ResponseWriter, *http.Request, ...string)
		expectedStatus int
		customMessage  string
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

	// Ensure we use the default handler for testing status codes and messages
	SetErrorHandler(nil)

	for _, tc := range testCases {
		// Test with custom message
		t.Run(fmt.Sprintf("%s with message", tc.name), func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)

			tc.function(rr, req, tc.customMessage)

			if rr.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, rr.Code)
			}
			// Verify JSON body contains the message
			if !strings.Contains(rr.Body.String(), tc.customMessage) {
				t.Errorf("expected body to contain '%s', got '%s'", tc.customMessage, rr.Body.String())
			}
		})

		// Test without message (should use default status text)
		t.Run(fmt.Sprintf("%s without message", tc.name), func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)

			tc.function(rr, req)

			if rr.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, rr.Code)
			}
			expectedMsg := http.StatusText(tc.expectedStatus)
			if !strings.Contains(rr.Body.String(), expectedMsg) {
				t.Errorf("expected body to contain '%s', got '%s'", expectedMsg, rr.Body.String())
			}
		})
	}
}
