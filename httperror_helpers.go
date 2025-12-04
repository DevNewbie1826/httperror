package httperror

import "net/http"

// HttpError represents an error with an associated HTTP status code.
// HttpError는 HTTP 상태 코드와 관련된 오류를 나타냅니다.
type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Error returns the error message.
// Error는 오류 메시지를 반환합니다.
func (e *HttpError) Error() string {
	return e.Message
}

// New creates a new HttpError.
// New는 새로운 HttpError를 생성합니다.
func New(status int, message string) *HttpError {
	return &HttpError{
		Status:  status,
		Message: message,
	}
}

// joinMessages is a helper to handle the variadic message argument.
func joinMessages(defaultMsg string, message []string) string {
	if len(message) > 0 {
		return message[0]
	}
	return defaultMsg
}

// InternalServerErrorError creates the HttpError struct for 500.
// This is used internally by DefaultErrorHandler for unknown errors.
func InternalServerErrorError(message ...string) *HttpError {
	return New(http.StatusInternalServerError, joinMessages(http.StatusText(http.StatusInternalServerError), message))
}

// --- Helper Functions ---

// BadRequest responds with a 400 Bad Request error.
// 잘못된 요청: 서버가 요청의 구문을 인식하지 못했습니다.
func BadRequest(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusBadRequest, joinMessages(http.StatusText(http.StatusBadRequest), message))
	Respond(w, r, err)
}

// Unauthorized responds with a 401 Unauthorized error.
// 인증 실패: 요청된 리소스에 대한 유효한 인증 자격 증명이 부족합니다.
func Unauthorized(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusUnauthorized, joinMessages(http.StatusText(http.StatusUnauthorized), message))
	Respond(w, r, err)
}

// PaymentRequired responds with a 402 Payment Required error.
// 결제 필요: 요청을 완료하려면 결제가 필요합니다.
func PaymentRequired(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusPaymentRequired, joinMessages(http.StatusText(http.StatusPaymentRequired), message))
	Respond(w, r, err)
}

// Forbidden responds with a 403 Forbidden error.
// 접근 금지: 서버가 요청을 이해했지만 승인을 거부했습니다.
func Forbidden(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusForbidden, joinMessages(http.StatusText(http.StatusForbidden), message))
	Respond(w, r, err)
}

// NotFound responds with a 404 Not Found error.
// 찾을 수 없음: 서버가 요청한 리소스를 찾을 수 없습니다.
func NotFound(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusNotFound, joinMessages(http.StatusText(http.StatusNotFound), message))
	Respond(w, r, err)
}

// MethodNotAllowed responds with a 405 Method Not Allowed error.
// 허용되지 않은 메소드: 요청한 리소스에 대해 요청한 메소드가 허용되지 않습니다.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusMethodNotAllowed, joinMessages(http.StatusText(http.StatusMethodNotAllowed), message))
	Respond(w, r, err)
}

// NotAcceptable responds with a 406 Not Acceptable error.
// 수용할 수 없음: 서버가 요청의 Accept 헤더에 따라 수용할 수 없는 응답을 생성할 수 없습니다.
func NotAcceptable(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusNotAcceptable, joinMessages(http.StatusText(http.StatusNotAcceptable), message))
	Respond(w, r, err)
}

// ProxyAuthRequired responds with a 407 Proxy Authentication Required error.
// 프록시 인증 필요: 프록시를 통해 인증해야 합니다.
func ProxyAuthRequired(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusProxyAuthRequired, joinMessages(http.StatusText(http.StatusProxyAuthRequired), message))
	Respond(w, r, err)
}

// RequestTimeout responds with a 408 Request Timeout error.
// 요청 시간 초과: 서버가 요청을 기다리는 동안 시간이 초과되었습니다.
func RequestTimeout(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusRequestTimeout, joinMessages(http.StatusText(http.StatusRequestTimeout), message))
	Respond(w, r, err)
}

// Conflict responds with a 409 Conflict error.
// 충돌: 요청이 리소스의 현재 상태와 충돌하여 완료될 수 없습니다.
func Conflict(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusConflict, joinMessages(http.StatusText(http.StatusConflict), message))
	Respond(w, r, err)
}

// Gone responds with a 410 Gone error.
// 사라짐: 요청한 리소스가 영구적으로 삭제되었습니다.
func Gone(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusGone, joinMessages(http.StatusText(http.StatusGone), message))
	Respond(w, r, err)
}

// LengthRequired responds with a 411 Length Required error.
// 길이 필요: Content-Length 헤더 없이 요청이 거부되었습니다.
func LengthRequired(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusLengthRequired, joinMessages(http.StatusText(http.StatusLengthRequired), message))
	Respond(w, r, err)
}

// PreconditionFailed responds with a 412 Precondition Failed error.
// 사전 조건 실패: 서버가 요청자가 요청에 지정한 사전 조건 중 하나를 충족하지 못했습니다.
func PreconditionFailed(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusPreconditionFailed, joinMessages(http.StatusText(http.StatusPreconditionFailed), message))
	Respond(w, r, err)
}

// PayloadTooLarge responds with a 413 Payload Too Large error.
// 페이로드 너무 큼: 요청 페이로드가 서버가 처리할 수 있는 한도보다 큽니다.
func PayloadTooLarge(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusRequestEntityTooLarge, joinMessages(http.StatusText(http.StatusRequestEntityTooLarge), message))
	Respond(w, r, err)
}

// URITooLong responds with a 414 URI Too Long error.
// URI 너무 긺: 클라이언트가 요청한 URI가 서버가 해석할 수 있는 것보다 깁니다.
func URITooLong(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusRequestURITooLong, joinMessages(http.StatusText(http.StatusRequestURITooLong), message))
	Respond(w, r, err)
}

// UnsupportedMediaType responds with a 415 Unsupported Media Type error.
// 지원되지 않는 미디어 유형: 서버가 요청 페이로드의 미디어 형식을 지원하지 않습니다.
func UnsupportedMediaType(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusUnsupportedMediaType, joinMessages(http.StatusText(http.StatusUnsupportedMediaType), message))
	Respond(w, r, err)
}

// RangeNotSatisfiable responds with a 416 Range Not Satisfiable error.
// 범위 만족할 수 없음: 요청의 Range 헤더 필드에 지정된 범위를 충족할 수 없습니다.
func RangeNotSatisfiable(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusRequestedRangeNotSatisfiable, joinMessages(http.StatusText(http.StatusRequestedRangeNotSatisfiable), message))
	Respond(w, r, err)
}

// ExpectationFailed responds with a 417 Expectation Failed error.
// 기대 실패: Expect 요청 헤더 필드에 지정된 기대를 충족할 수 없습니다.
func ExpectationFailed(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusExpectationFailed, joinMessages(http.StatusText(http.StatusExpectationFailed), message))
	Respond(w, r, err)
}

// Teapot responds with a 418 I'm a teapot error.
// 나는 찻주전자: 나는 찻주전자입니다.
func Teapot(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusTeapot, joinMessages(http.StatusText(http.StatusTeapot), message))
	Respond(w, r, err)
}

// MisdirectedRequest responds with a 421 Misdirected Request error.
// 잘못된 요청: 요청이 응답을 생성할 수 없는 서버로 전달되었습니다.
func MisdirectedRequest(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusMisdirectedRequest, joinMessages(http.StatusText(http.StatusMisdirectedRequest), message))
	Respond(w, r, err)
}

// UnprocessableEntity responds with a 422 Unprocessable Entity error.
// 처리할 수 없는 엔티티: 서버가 요청을 이해했지만, 의미론적 오류로 인해 처리할 수 없습니다.
func UnprocessableEntity(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusUnprocessableEntity, joinMessages(http.StatusText(http.StatusUnprocessableEntity), message))
	Respond(w, r, err)
}

// Locked responds with a 423 Locked error.
// 잠김: 접근하려는 리소스가 잠겨 있습니다.
func Locked(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusLocked, joinMessages(http.StatusText(http.StatusLocked), message))
	Respond(w, r, err)
}

// FailedDependency responds with a 424 Failed Dependency error.
// 실패한 종속성: 이전 요청이 실패했기 때문에 현재 요청이 실패했습니다.
func FailedDependency(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusFailedDependency, joinMessages(http.StatusText(http.StatusFailedDependency), message))
	Respond(w, r, err)
}

// TooEarly responds with a 425 Too Early error.
// 너무 이름: 서버가 아직 처리 준비가 되지 않은 요청을 처리하려고 시도했습니다.
func TooEarly(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusTooEarly, joinMessages(http.StatusText(http.StatusTooEarly), message))
	Respond(w, r, err)
}

// UpgradeRequired responds with a 426 Upgrade Required error.
// 업그레이드 필요: 클라이언트는 다른 프로토콜로 업그레이드해야 합니다.
func UpgradeRequired(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusUpgradeRequired, joinMessages(http.StatusText(http.StatusUpgradeRequired), message))
	Respond(w, r, err)
}

// PreconditionRequired responds with a 428 Precondition Required error.
// 사전 조건 필요: 원본 서버는 요청이 조건부여야 함을 요구합니다.
func PreconditionRequired(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusPreconditionRequired, joinMessages(http.StatusText(http.StatusPreconditionRequired), message))
	Respond(w, r, err)
}

// TooManyRequests responds with a 429 Too Many Requests error.
// 너무 많은 요청: 사용자가 지정된 시간 동안 너무 많은 요청을 보냈습니다.
func TooManyRequests(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusTooManyRequests, joinMessages(http.StatusText(http.StatusTooManyRequests), message))
	Respond(w, r, err)
}

// RequestHeaderFieldsTooLarge responds with a 431 Request Header Fields Too Large error.
// 요청 헤더 필드 너무 큼: 요청 헤더 필드가 너무 커서 서버가 처리할 수 없습니다.
func RequestHeaderFieldsTooLarge(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusRequestHeaderFieldsTooLarge, joinMessages(http.StatusText(http.StatusRequestHeaderFieldsTooLarge), message))
	Respond(w, r, err)
}

// UnavailableForLegalReasons responds with a 451 Unavailable For Legal Reasons error.
// 법적 이유로 사용할 수 없음: 법적인 이유로 요청한 리소스에 접근할 수 없습니다.
func UnavailableForLegalReasons(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusUnavailableForLegalReasons, joinMessages(http.StatusText(http.StatusUnavailableForLegalReasons), message))
	Respond(w, r, err)
}

// InternalServerError responds with a 500 Internal Server Error.
// 내부 서버 오류: 서버에 예기치 않은 오류가 발생했습니다.
func InternalServerError(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusInternalServerError, joinMessages(http.StatusText(http.StatusInternalServerError), message))
	Respond(w, r, err)
}

// NotImplemented responds with a 501 Not Implemented error.
// 구현되지 않음: 서버가 요청을 수행하는 데 필요한 기능을 지원하지 않습니다.
func NotImplemented(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusNotImplemented, joinMessages(http.StatusText(http.StatusNotImplemented), message))
	Respond(w, r, err)
}

// BadGateway responds with a 502 Bad Gateway error.
// 잘못된 게이트웨이: 서버가 게이트웨이 또는 프록시 역할을 하는 동안 업스트림 서버로부터 잘못된 응답을 받았습니다.
func BadGateway(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusBadGateway, joinMessages(http.StatusText(http.StatusBadGateway), message))
	Respond(w, r, err)
}

// ServiceUnavailable responds with a 503 Service Unavailable error.
// 서비스 사용 불가: 서버가 일시적으로 요청을 처리할 수 없습니다.
func ServiceUnavailable(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusServiceUnavailable, joinMessages(http.StatusText(http.StatusServiceUnavailable), message))
	Respond(w, r, err)
}

// GatewayTimeout responds with a 504 Gateway Timeout error.
// 게이트웨이 시간 초과: 서버가 게이트웨이 또는 프록시 역할을 하는 동안 업스트림 서버로부터 응답을 받지 못했습니다.
func GatewayTimeout(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusGatewayTimeout, joinMessages(http.StatusText(http.StatusGatewayTimeout), message))
	Respond(w, r, err)
}

// HTTPVersionNotSupported responds with a 505 HTTP Version Not Supported error.
// 지원되지 않는 HTTP 버전: 서버가 요청에 사용된 HTTP 버전을 지원하지 않습니다.
func HTTPVersionNotSupported(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusHTTPVersionNotSupported, joinMessages(http.StatusText(http.StatusHTTPVersionNotSupported), message))
	Respond(w, r, err)
}

// VariantAlsoNegotiates responds with a 506 Variant Also Negotiates error.
// 변형도 협상함: 서버에 내부 구성 오류가 있습니다.
func VariantAlsoNegotiates(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusVariantAlsoNegotiates, joinMessages(http.StatusText(http.StatusVariantAlsoNegotiates), message))
	Respond(w, r, err)
}

// InsufficientStorage responds with a 507 Insufficient Storage error.
// 저장 공간 부족: 서버에 요청을 완료하는 데 필요한 저장 공간이 부족합니다.
func InsufficientStorage(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusInsufficientStorage, joinMessages(http.StatusText(http.StatusInsufficientStorage), message))
	Respond(w, r, err)
}

// LoopDetected responds with a 508 Loop Detected error.
// 루프 감지됨: 서버가 요청을 처리하는 동안 무한 루프를 감지했습니다.
func LoopDetected(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusLoopDetected, joinMessages(http.StatusText(http.StatusLoopDetected), message))
	Respond(w, r, err)
}

// NotExtended responds with a 510 Not Extended error.
// 확장되지 않음: 요청을 이행하기 위해 추가 확장이 필요합니다.
func NotExtended(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusNotExtended, joinMessages(http.StatusText(http.StatusNotExtended), message))
	Respond(w, r, err)
}

// NetworkAuthenticationRequired responds with a 511 Network Authentication Required error.
// 네트워크 인증 필요: 클라이언트는 네트워크 접근 권한을 얻기 위해 인증해야 합니다.
func NetworkAuthenticationRequired(w http.ResponseWriter, r *http.Request, message ...string) {
	err := New(http.StatusNetworkAuthenticationRequired, joinMessages(http.StatusText(http.StatusNetworkAuthenticationRequired), message))
	Respond(w, r, err)
}