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

// BadRequest creates a 400 Bad Request error.
// 잘못된 요청: 서버가 요청의 구문을 인식하지 못했습니다.
func BadRequest(message ...string) *HttpError {
	msg := http.StatusText(http.StatusBadRequest)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusBadRequest, msg)
}

// Unauthorized creates a 401 Unauthorized error.
// 인증 실패: 요청된 리소스에 대한 유효한 인증 자격 증명이 부족합니다.
func Unauthorized(message ...string) *HttpError {
	msg := http.StatusText(http.StatusUnauthorized)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusUnauthorized, msg)
}

// PaymentRequired creates a 402 Payment Required error.
// 결제 필요: 요청을 완료하려면 결제가 필요합니다.
func PaymentRequired(message ...string) *HttpError {
	msg := http.StatusText(http.StatusPaymentRequired)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusPaymentRequired, msg)
}

// Forbidden creates a 403 Forbidden error.
// 접근 금지: 서버가 요청을 이해했지만 승인을 거부했습니다.
func Forbidden(message ...string) *HttpError {
	msg := http.StatusText(http.StatusForbidden)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusForbidden, msg)
}

// NotFound creates a 404 Not Found error.
// 찾을 수 없음: 서버가 요청한 리소스를 찾을 수 없습니다.
func NotFound(message ...string) *HttpError {
	msg := http.StatusText(http.StatusNotFound)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusNotFound, msg)
}

// MethodNotAllowed creates a 405 Method Not Allowed error.
// 허용되지 않은 메소드: 요청한 리소스에 대해 요청한 메소드가 허용되지 않습니다.
func MethodNotAllowed(message ...string) *HttpError {
	msg := http.StatusText(http.StatusMethodNotAllowed)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusMethodNotAllowed, msg)
}

// NotAcceptable creates a 406 Not Acceptable error.
// 수용할 수 없음: 서버가 요청의 Accept 헤더에 따라 수용할 수 없는 응답을 생성할 수 없습니다.
func NotAcceptable(message ...string) *HttpError {
	msg := http.StatusText(http.StatusNotAcceptable)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusNotAcceptable, msg)
}

// ProxyAuthRequired creates a 407 Proxy Authentication Required error.
// 프록시 인증 필요: 프록시를 통해 인증해야 합니다.
func ProxyAuthRequired(message ...string) *HttpError {
	msg := http.StatusText(http.StatusProxyAuthRequired)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusProxyAuthRequired, msg)
}

// RequestTimeout creates a 408 Request Timeout error.
// 요청 시간 초과: 서버가 요청을 기다리는 동안 시간이 초과되었습니다.
func RequestTimeout(message ...string) *HttpError {
	msg := http.StatusText(http.StatusRequestTimeout)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusRequestTimeout, msg)
}

// Conflict creates a 409 Conflict error.
// 충돌: 요청이 리소스의 현재 상태와 충돌하여 완료될 수 없습니다.
func Conflict(message ...string) *HttpError {
	msg := http.StatusText(http.StatusConflict)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusConflict, msg)
}

// Gone creates a 410 Gone error.
// 사라짐: 요청한 리소스가 영구적으로 삭제되었습니다.
func Gone(message ...string) *HttpError {
	msg := http.StatusText(http.StatusGone)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusGone, msg)
}

// LengthRequired creates a 411 Length Required error.
// 길이 필요: Content-Length 헤더 없이 요청이 거부되었습니다.
func LengthRequired(message ...string) *HttpError {
	msg := http.StatusText(http.StatusLengthRequired)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusLengthRequired, msg)
}

// PreconditionFailed creates a 412 Precondition Failed error.
// 사전 조건 실패: 서버가 요청자가 요청에 지정한 사전 조건 중 하나를 충족하지 못했습니다.
func PreconditionFailed(message ...string) *HttpError {
	msg := http.StatusText(http.StatusPreconditionFailed)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusPreconditionFailed, msg)
}

// PayloadTooLarge creates a 413 Payload Too Large error.
// 페이로드 너무 큼: 요청 페이로드가 서버가 처리할 수 있는 한도보다 큽니다.
func PayloadTooLarge(message ...string) *HttpError {
	msg := http.StatusText(http.StatusRequestEntityTooLarge)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusRequestEntityTooLarge, msg)
}

// URITooLong creates a 414 URI Too Long error.
// URI 너무 긺: 클라이언트가 요청한 URI가 서버가 해석할 수 있는 것보다 깁니다.
func URITooLong(message ...string) *HttpError {
	msg := http.StatusText(http.StatusRequestURITooLong)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusRequestURITooLong, msg)
}

// UnsupportedMediaType creates a 415 Unsupported Media Type error.
// 지원되지 않는 미디어 유형: 서버가 요청 페이로드의 미디어 형식을 지원하지 않습니다.
func UnsupportedMediaType(message ...string) *HttpError {
	msg := http.StatusText(http.StatusUnsupportedMediaType)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusUnsupportedMediaType, msg)
}

// RangeNotSatisfiable creates a 416 Range Not Satisfiable error.
// 범위 만족할 수 없음: 요청의 Range 헤더 필드에 지정된 범위를 충족할 수 없습니다.
func RangeNotSatisfiable(message ...string) *HttpError {
	msg := http.StatusText(http.StatusRequestedRangeNotSatisfiable)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusRequestedRangeNotSatisfiable, msg)
}

// ExpectationFailed creates a 417 Expectation Failed error.
// 기대 실패: Expect 요청 헤더 필드에 지정된 기대를 충족할 수 없습니다.
func ExpectationFailed(message ...string) *HttpError {
	msg := http.StatusText(http.StatusExpectationFailed)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusExpectationFailed, msg)
}

// Teapot creates a 418 I'm a teapot error.
// 나는 찻주전자: 나는 찻주전자입니다.
func Teapot(message ...string) *HttpError {
	msg := http.StatusText(http.StatusTeapot)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusTeapot, msg)
}

// MisdirectedRequest creates a 421 Misdirected Request error.
// 잘못된 요청: 요청이 응답을 생성할 수 없는 서버로 전달되었습니다.
func MisdirectedRequest(message ...string) *HttpError {
	msg := http.StatusText(http.StatusMisdirectedRequest)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusMisdirectedRequest, msg)
}

// UnprocessableEntity creates a 422 Unprocessable Entity error.
// 처리할 수 없는 엔티티: 서버가 요청을 이해했지만, 의미론적 오류로 인해 처리할 수 없습니다.
func UnprocessableEntity(message ...string) *HttpError {
	msg := http.StatusText(http.StatusUnprocessableEntity)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusUnprocessableEntity, msg)
}

// Locked creates a 423 Locked error.
// 잠김: 접근하려는 리소스가 잠겨 있습니다.
func Locked(message ...string) *HttpError {
	msg := http.StatusText(http.StatusLocked)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusLocked, msg)
}

// FailedDependency creates a 424 Failed Dependency error.
// 실패한 종속성: 이전 요청이 실패했기 때문에 현재 요청이 실패했습니다.
func FailedDependency(message ...string) *HttpError {
	msg := http.StatusText(http.StatusFailedDependency)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusFailedDependency, msg)
}

// TooEarly creates a 425 Too Early error.
// 너무 이름: 서버가 아직 처리 준비가 되지 않은 요청을 처리하려고 시도했습니다.
func TooEarly(message ...string) *HttpError {
	msg := http.StatusText(http.StatusTooEarly)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusTooEarly, msg)
}

// UpgradeRequired creates a 426 Upgrade Required error.
// 업그레이드 필요: 클라이언트는 다른 프로토콜로 업그레이드해야 합니다.
func UpgradeRequired(message ...string) *HttpError {
	msg := http.StatusText(http.StatusUpgradeRequired)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusUpgradeRequired, msg)
}

// PreconditionRequired creates a 428 Precondition Required error.
// 사전 조건 필요: 원본 서버는 요청이 조건부여야 함을 요구합니다.
func PreconditionRequired(message ...string) *HttpError {
	msg := http.StatusText(http.StatusPreconditionRequired)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusPreconditionRequired, msg)
}

// TooManyRequests creates a 429 Too Many Requests error.
// 너무 많은 요청: 사용자가 지정된 시간 동안 너무 많은 요청을 보냈습니다.
func TooManyRequests(message ...string) *HttpError {
	msg := http.StatusText(http.StatusTooManyRequests)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusTooManyRequests, msg)
}

// RequestHeaderFieldsTooLarge creates a 431 Request Header Fields Too Large error.
// 요청 헤더 필드 너무 큼: 요청 헤더 필드가 너무 커서 서버가 처리할 수 없습니다.
func RequestHeaderFieldsTooLarge(message ...string) *HttpError {
	msg := http.StatusText(http.StatusRequestHeaderFieldsTooLarge)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusRequestHeaderFieldsTooLarge, msg)
}

// UnavailableForLegalReasons creates a 451 Unavailable For Legal Reasons error.
// 법적 이유로 사용할 수 없음: 법적인 이유로 요청한 리소스에 접근할 수 없습니다.
func UnavailableForLegalReasons(message ...string) *HttpError {
	msg := http.StatusText(http.StatusUnavailableForLegalReasons)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusUnavailableForLegalReasons, msg)
}

// InternalServerError creates a 500 Internal Server Error.
// 내부 서버 오류: 서버에 예기치 않은 오류가 발생했습니다.
func InternalServerError(message ...string) *HttpError {
	msg := http.StatusText(http.StatusInternalServerError)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusInternalServerError, msg)
}

// NotImplemented creates a 501 Not Implemented error.
// 구현되지 않음: 서버가 요청을 수행하는 데 필요한 기능을 지원하지 않습니다.
func NotImplemented(message ...string) *HttpError {
	msg := http.StatusText(http.StatusNotImplemented)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusNotImplemented, msg)
}

// BadGateway creates a 502 Bad Gateway error.
// 잘못된 게이트웨이: 서버가 게이트웨이 또는 프록시 역할을 하는 동안 업스트림 서버로부터 잘못된 응답을 받았습니다.
func BadGateway(message ...string) *HttpError {
	msg := http.StatusText(http.StatusBadGateway)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusBadGateway, msg)
}

// ServiceUnavailable creates a 503 Service Unavailable error.
// 서비스 사용 불가: 서버가 일시적으로 요청을 처리할 수 없습니다.
func ServiceUnavailable(message ...string) *HttpError {
	msg := http.StatusText(http.StatusServiceUnavailable)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusServiceUnavailable, msg)
}

// GatewayTimeout creates a 504 Gateway Timeout error.
// 게이트웨이 시간 초과: 서버가 게이트웨이 또는 프록시 역할을 하는 동안 업스트림 서버로부터 응답을 받지 못했습니다.
func GatewayTimeout(message ...string) *HttpError {
	msg := http.StatusText(http.StatusGatewayTimeout)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusGatewayTimeout, msg)
}

// HTTPVersionNotSupported creates a 505 HTTP Version Not Supported error.
// 지원되지 않는 HTTP 버전: 서버가 요청에 사용된 HTTP 버전을 지원하지 않습니다.
func HTTPVersionNotSupported(message ...string) *HttpError {
	msg := http.StatusText(http.StatusHTTPVersionNotSupported)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusHTTPVersionNotSupported, msg)
}

// VariantAlsoNegotiates creates a 506 Variant Also Negotiates error.
// 변형도 협상함: 서버에 내부 구성 오류가 있습니다.
func VariantAlsoNegotiates(message ...string) *HttpError {
	msg := http.StatusText(http.StatusVariantAlsoNegotiates)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusVariantAlsoNegotiates, msg)
}

// InsufficientStorage creates a 507 Insufficient Storage error.
// 저장 공간 부족: 서버에 요청을 완료하는 데 필요한 저장 공간이 부족합니다.
func InsufficientStorage(message ...string) *HttpError {
	msg := http.StatusText(http.StatusInsufficientStorage)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusInsufficientStorage, msg)
}

// LoopDetected creates a 508 Loop Detected error.
// 루프 감지됨: 서버가 요청을 처리하는 동안 무한 루프를 감지했습니다.
func LoopDetected(message ...string) *HttpError {
	msg := http.StatusText(http.StatusLoopDetected)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusLoopDetected, msg)
}

// NotExtended creates a 510 Not Extended error.
// 확장되지 않음: 요청을 이행하기 위해 추가 확장이 필요합니다.
func NotExtended(message ...string) *HttpError {
	msg := http.StatusText(http.StatusNotExtended)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusNotExtended, msg)
}

// NetworkAuthenticationRequired creates a 511 Network Authentication Required error.
// 네트워크 인증 필요: 클라이언트는 네트워크 접근 권한을 얻기 위해 인증해야 합니다.
func NetworkAuthenticationRequired(message ...string) *HttpError {
	msg := http.StatusText(http.StatusNetworkAuthenticationRequired)
	if len(message) > 0 {
		msg = message[0]
	}
	return New(http.StatusNetworkAuthenticationRequired, msg)
}

// --- Reporter Functions ---

// ReportBadRequest reports a 400 Bad Request error to the middleware.
// ReportBadRequest는 400 잘못된 요청 오류를 미들웨어에 보고합니다.
func ReportBadRequest(r *http.Request, message ...string) {
	ReportError(r, BadRequest(message...))
}

// ReportUnauthorized reports a 401 Unauthorized error to the middleware.
// ReportUnauthorized는 401 인증 실패 오류를 미들웨어에 보고합니다.
func ReportUnauthorized(r *http.Request, message ...string) {
	ReportError(r, Unauthorized(message...))
}

// ReportPaymentRequired reports a 402 Payment Required error to the middleware.
// ReportPaymentRequired는 402 결제 필요 오류를 미들웨어에 보고합니다.
func ReportPaymentRequired(r *http.Request, message ...string) {
	ReportError(r, PaymentRequired(message...))
}

// ReportForbidden reports a 403 Forbidden error to the middleware.
// ReportForbidden는 403 접근 금지 오류를 미들웨어에 보고합니다.
func ReportForbidden(r *http.Request, message ...string) {
	ReportError(r, Forbidden(message...))
}

// ReportNotFound reports a 404 Not Found error to the middleware.
// ReportNotFound는 404 찾을 수 없음 오류를 미들웨어에 보고합니다.
func ReportNotFound(r *http.Request, message ...string) {
	ReportError(r, NotFound(message...))
}

// ReportMethodNotAllowed reports a 405 Method Not Allowed error to the middleware.
// ReportMethodNotAllowed는 405 허용되지 않은 메소드 오류를 미들웨어에 보고합니다.
func ReportMethodNotAllowed(r *http.Request, message ...string) {
	ReportError(r, MethodNotAllowed(message...))
}

// ReportNotAcceptable reports a 406 Not Acceptable error to the middleware.
// ReportNotAcceptable는 406 수용할 수 없음 오류를 미들웨어에 보고합니다.
func ReportNotAcceptable(r *http.Request, message ...string) {
	ReportError(r, NotAcceptable(message...))
}

// ReportProxyAuthRequired reports a 407 Proxy Authentication Required error to the middleware.
// ReportProxyAuthRequired는 407 프록시 인증 필요 오류를 미들웨어에 보고합니다.
func ReportProxyAuthRequired(r *http.Request, message ...string) {
	ReportError(r, ProxyAuthRequired(message...))
}

// ReportRequestTimeout reports a 408 Request Timeout error to the middleware.
// ReportRequestTimeout는 408 요청 시간 초과 오류를 미들웨어에 보고합니다.
func ReportRequestTimeout(r *http.Request, message ...string) {
	ReportError(r, RequestTimeout(message...))
}

// ReportConflict reports a 409 Conflict error to the middleware.
// ReportConflict는 409 충돌 오류를 미들웨어에 보고합니다.
func ReportConflict(r *http.Request, message ...string) {
	ReportError(r, Conflict(message...))
}

// ReportGone reports a 410 Gone error to the middleware.
// ReportGone는 410 사라짐 오류를 미들웨어에 보고합니다.
func ReportGone(r *http.Request, message ...string) {
	ReportError(r, Gone(message...))
}

// ReportLengthRequired reports a 411 Length Required error to the middleware.
// ReportLengthRequired는 411 길이 필요 오류를 미들웨어에 보고합니다.
func ReportLengthRequired(r *http.Request, message ...string) {
	ReportError(r, LengthRequired(message...))
}

// ReportPreconditionFailed reports a 412 Precondition Failed error to the middleware.
// ReportPreconditionFailed는 412 사전 조건 실패 오류를 미들웨어에 보고합니다.
func ReportPreconditionFailed(r *http.Request, message ...string) {
	ReportError(r, PreconditionFailed(message...))
}

// ReportPayloadTooLarge reports a 413 Payload Too Large error to the middleware.
// ReportPayloadTooLarge는 413 페이로드 너무 큼 오류를 미들웨어에 보고합니다.
func ReportPayloadTooLarge(r *http.Request, message ...string) {
	ReportError(r, PayloadTooLarge(message...))
}

// ReportURITooLong reports a 414 URI Too Long error to the middleware.
// ReportURITooLong는 414 URI 너무 긺 오류를 미들웨어에 보고합니다.
func ReportURITooLong(r *http.Request, message ...string) {
	ReportError(r, URITooLong(message...))
}

// ReportUnsupportedMediaType reports a 415 Unsupported Media Type error to the middleware.
// ReportUnsupportedMediaType는 415 지원되지 않는 미디어 유형 오류를 미들웨어에 보고합니다.
func ReportUnsupportedMediaType(r *http.Request, message ...string) {
	ReportError(r, UnsupportedMediaType(message...))
}

// ReportRangeNotSatisfiable reports a 416 Range Not Satisfiable error to the middleware.
// ReportRangeNotSatisfiable는 416 범위 만족할 수 없음 오류를 미들웨어에 보고합니다.
func ReportRangeNotSatisfiable(r *http.Request, message ...string) {
	ReportError(r, RangeNotSatisfiable(message...))
}

// ReportExpectationFailed reports a 417 Expectation Failed error to the middleware.
// ReportExpectationFailed는 417 기대 실패 오류를 미들웨어에 보고합니다.
func ReportExpectationFailed(r *http.Request, message ...string) {
	ReportError(r, ExpectationFailed(message...))
}

// ReportTeapot reports a 418 I'm a teapot error to the middleware.
// ReportTeapot는 418 나는 찻주전자 오류를 미들웨어에 보고합니다.
func ReportTeapot(r *http.Request, message ...string) {
	ReportError(r, Teapot(message...))
}

// ReportMisdirectedRequest reports a 421 Misdirected Request error to the middleware.
// ReportMisdirectedRequest는 421 잘못된 요청 오류를 미들웨어에 보고합니다.
func ReportMisdirectedRequest(r *http.Request, message ...string) {
	ReportError(r, MisdirectedRequest(message...))
}

// ReportUnprocessableEntity reports a 422 Unprocessable Entity error to the middleware.
// ReportUnprocessableEntity는 422 처리할 수 없는 엔티티 오류를 미들웨어에 보고합니다.
func ReportUnprocessableEntity(r *http.Request, message ...string) {
	ReportError(r, UnprocessableEntity(message...))
}

// ReportLocked reports a 423 Locked error to the middleware.
// ReportLocked는 423 잠김 오류를 미들웨어에 보고합니다.
func ReportLocked(r *http.Request, message ...string) {
	ReportError(r, Locked(message...))
}

// ReportFailedDependency reports a 424 Failed Dependency error to the middleware.
// ReportFailedDependency는 424 실패한 종속성 오류를 미들웨어에 보고합니다.
func ReportFailedDependency(r *http.Request, message ...string) {
	ReportError(r, FailedDependency(message...))
}

// ReportTooEarly reports a 425 Too Early error to the middleware.
// ReportTooEarly는 425 너무 이름 오류를 미들웨어에 보고합니다.
func ReportTooEarly(r *http.Request, message ...string) {
	ReportError(r, TooEarly(message...))
}

// ReportUpgradeRequired reports a 426 Upgrade Required error to the middleware.
// ReportUpgradeRequired는 426 업그레이드 필요 오류를 미들웨어에 보고합니다.
func ReportUpgradeRequired(r *http.Request, message ...string) {
	ReportError(r, UpgradeRequired(message...))
}

// ReportPreconditionRequired reports a 428 Precondition Required error to the middleware.
// ReportPreconditionRequired는 428 사전 조건 필요 오류를 미들웨어에 보고합니다.
func ReportPreconditionRequired(r *http.Request, message ...string) {
	ReportError(r, PreconditionRequired(message...))
}

// ReportTooManyRequests reports a 429 Too Many Requests error to the middleware.
// ReportTooManyRequests는 429 너무 많은 요청 오류를 미들웨어에 보고합니다.
func ReportTooManyRequests(r *http.Request, message ...string) {
	ReportError(r, TooManyRequests(message...))
}

// ReportRequestHeaderFieldsTooLarge reports a 431 Request Header Fields Too Large error to the middleware.
// ReportRequestHeaderFieldsTooLarge는 431 요청 헤더 필드 너무 큼 오류를 미들웨어에 보고합니다.
func ReportRequestHeaderFieldsTooLarge(r *http.Request, message ...string) {
	ReportError(r, RequestHeaderFieldsTooLarge(message...))
}

// ReportUnavailableForLegalReasons reports a 451 Unavailable For Legal Reasons error to the middleware.
// ReportUnavailableForLegalReasons는 451 법적 이유로 사용할 수 없음 오류를 미들웨어에 보고합니다.
func ReportUnavailableForLegalReasons(r *http.Request, message ...string) {
	ReportError(r, UnavailableForLegalReasons(message...))
}

// ReportInternalServerError reports a 500 Internal Server Error to the middleware.
// ReportInternalServerError는 500 내부 서버 오류를 미들웨어에 보고합니다.
func ReportInternalServerError(r *http.Request, message ...string) {
	ReportError(r, InternalServerError(message...))
}

// ReportNotImplemented reports a 501 Not Implemented error to the middleware.
// ReportNotImplemented는 501 구현되지 않음 오류를 미들웨어에 보고합니다.
func ReportNotImplemented(r *http.Request, message ...string) {
	ReportError(r, NotImplemented(message...))
}

// ReportBadGateway reports a 502 Bad Gateway error to the middleware.
// ReportBadGateway는 502 잘못된 게이트웨이 오류를 미들웨어에 보고합니다.
func ReportBadGateway(r *http.Request, message ...string) {
	ReportError(r, BadGateway(message...))
}

// ReportServiceUnavailable reports a 503 Service Unavailable error to the middleware.
// ReportServiceUnavailable는 503 서비스 사용 불가 오류를 미들웨어에 보고합니다.
func ReportServiceUnavailable(r *http.Request, message ...string) {
	ReportError(r, ServiceUnavailable(message...))
}

// ReportGatewayTimeout reports a 504 Gateway Timeout error to the middleware.
// ReportGatewayTimeout는 504 게이트웨이 시간 초과 오류를 미들웨어에 보고합니다.
func ReportGatewayTimeout(r *http.Request, message ...string) {
	ReportError(r, GatewayTimeout(message...))
}

// ReportHTTPVersionNotSupported reports a 505 HTTP Version Not Supported error to the middleware.
// ReportHTTPVersionNotSupported는 505 지원되지 않는 HTTP 버전 오류를 미들웨어에 보고합니다.
func ReportHTTPVersionNotSupported(r *http.Request, message ...string) {
	ReportError(r, HTTPVersionNotSupported(message...))
}

// ReportVariantAlsoNegotiates reports a 506 Variant Also Negotiates error to the middleware.
// ReportVariantAlsoNegotiates는 506 변형도 협상함 오류를 미들웨어에 보고합니다.
func ReportVariantAlsoNegotiates(r *http.Request, message ...string) {
	ReportError(r, VariantAlsoNegotiates(message...))
}

// ReportInsufficientStorage reports a 507 Insufficient Storage error to the middleware.
// ReportInsufficientStorage는 507 저장 공간 부족 오류를 미들웨어에 보고합니다.
func ReportInsufficientStorage(r *http.Request, message ...string) {
	ReportError(r, InsufficientStorage(message...))
}

// ReportLoopDetected reports a 508 Loop Detected error to the middleware.
// ReportLoopDetected는 508 루프 감지됨 오류를 미들웨어에 보고합니다.
func ReportLoopDetected(r *http.Request, message ...string) {
	ReportError(r, LoopDetected(message...))
}

// ReportNotExtended reports a 510 Not Extended error to the middleware.
// ReportNotExtended는 510 확장되지 않음 오류를 미들웨어에 보고합니다.
func ReportNotExtended(r *http.Request, message ...string) {
	ReportError(r, NotExtended(message...))
}

// ReportNetworkAuthenticationRequired reports a 511 Network Authentication Required error to the middleware.
// ReportNetworkAuthenticationRequired는 511 네트워크 인증 필요 오류를 미들웨어에 보고합니다.
func ReportNetworkAuthenticationRequired(r *http.Request, message ...string) {
	ReportError(r, NetworkAuthenticationRequired(message...))
}
