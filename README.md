# httperror

[![Go Report Card](https://goreportcard.com/badge/github.com/DevNewbie1826/httperror)](https://goreportcard.com/report/github.com/DevNewbie1826/httperror)
[![GoDoc](https://godoc.org/github.com/DevNewbie1826/httperror?status.svg)](https://godoc.org/github.com/DevNewbie1826/httperror)
[![codecov](https://codecov.io/gh/DevNewbie1826/httperror/graph/badge.svg)](https://codecov.io/gh/DevNewbie1826/httperror)

A simple Go middleware for centralized HTTP error handling, with convenient helper functions for generating standard HTTP error responses.

---

## English

### 📖 Overview

`httperror` provides a simple yet effective way to handle errors in your Go web application. It offers a middleware that catches errors reported from your HTTP handlers and responds with a standardized JSON error message. It also includes a rich set of helper functions to easily generate errors for most standard HTTP status codes.

### 🚀 Installation

```bash
go get github.com/DevNewbie1826/httperror
```

### ✨ Features

-   Centralized error handling via middleware.
-   Default error handler that responds with JSON.
-   Option to use a custom error handler.
-   A comprehensive set of helper functions for creating and reporting errors (e.g., `httperror.NotFound()`, `httperror.ReportBadRequest()`).
-   Context-based error reporting that avoids polluting handler function signatures.

### 💡 Usage

#### Basic Example

Wrap your router or `http.Handler` with `httperror.ErrorReporterMiddleware` and use the `Report...` functions within your handlers to propagate errors.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevNewbie1826/httperror"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// Some logic to find a user...
		userFound := false

		if !userFound {
			// If the user is not found, report a 404 Not Found error.
			// The middleware will handle the response.
			httperror.ReportNotFound(r, "User with the specified ID was not found.")
			return
		}

		fmt.Fprintln(w, "User data would be here.")
	})
    mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
        // Report a 403 Forbidden error.
        httperror.ReportForbidden(r)
    })


	// Wrap the mux with the error reporter middleware.
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", httperror.ErrorReporterMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}
```

When you run this server and access `http://localhost:8080/users`, you will get a JSON response like this:

```json
{
	"status": 404,
	"message": "User with the specified ID was not found."
}
```

#### Custom Error Handler

You can also provide your own custom error handling logic by using `NewErrorReporterMiddleware`.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevNewbie1826/httperror"
)

// A custom handler to log errors and write a custom response.
func myCustomErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("An error occurred: %v", err)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if httpErr, ok := err.(*httperror.HttpError); ok {
		w.WriteHeader(httpErr.Status)
		fmt.Fprintf(w, `{"errorCode": %d, "errorMessage": "%s"}`, httpErr.Status, httpErr.Message)
		return
	}

	// Fallback for non-HttpError types
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, `{"errorCode": 500, "errorMessage": "An unexpected error occurred."}`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
        // Report a 400 Bad Request error.
		httperror.ReportBadRequest(r, "Missing required query parameter 'id'.")
	})

	// Create a new middleware instance with the custom handler.
	errorMiddleware := httperror.NewErrorReporterMiddleware(myCustomErrorHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", errorMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}
```

#### Usage with `chi` router

When using a router like `chi`, you might want to handle routing errors (like 404 Not Found or 405 Method Not Allowed) in the same way as your application errors. `chi` allows you to set custom handlers for these cases.

You can integrate `httperror` by setting custom handlers that use the `Report...` functions.

**Important:** The `httperror.ErrorReporterMiddleware` must be registered before your routes so that the context is available to the `NotFound` and `MethodNotAllowed` handlers.

```go
package main

import (
	"log"
	"net/http"

	"github.com/DevNewbie1826/httperror"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// The ErrorReporterMiddleware must be applied before the router's handlers.
	// This ensures the context is available for the NotFound and MethodNotAllowed handlers.
	r.Use(httperror.ErrorReporterMiddleware)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Override chi's default NotFound handler
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		httperror.ReportNotFound(r, "The requested resource path does not exist.")
	})

	// Override chi's default MethodNotAllowed handler
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		httperror.ReportMethodNotAllowed(r)
	})

	log.Println("Server starting on :8080 with chi")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
```

---

## 한국어

### 📖 개요

`httperror`는 Go 웹 애플리케이션의 오류를 처리하는 간단하고 효과적인 방법을 제공합니다. 이 라이브러리는 HTTP 핸들러에서 보고된 오류를 감지하여 표준화된 JSON 오류 메시지로 응답하는 미들웨어를 제공합니다. 또한, 대부분의 표준 HTTP 상태 코드에 대한 오류를 쉽게 생성할 수 있는 풍부한 헬퍼 함수들을 포함하고 있습니다.

### 🚀 설치

```bash
go get github.com/DevNewbie1826/httperror
```

### ✨ 주요 기능

-   미들웨어를 통한 중앙 집중식 오류 처리
-   JSON 형식으로 응답하는 기본 오류 핸들러
-   사용자 정의 오류 핸들러를 지정할 수 있는 옵션
-   오류 생성 및 보고를 위한 포괄적인 헬퍼 함수 세트 (예: `httperror.NotFound()`, `httperror.ReportBadRequest()`)
-   핸들러 함수의 시그니처를 오염시키지 않는 컨텍스트 기반 오류 보고

### 💡 사용법

#### 기본 예제

사용하시는 라우터나 `http.Handler`를 `httperror.ErrorReporterMiddleware`로 감싸고, 핸들러 내에서 `Report...` 함수들을 사용하여 오류를 미들웨어로 보고하세요.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevNewbie1826/httperror"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// 사용자를 찾는 로직이 있다고 가정...
		userFound := false

		if !userFound {
			// 사용자를 찾지 못했다면, 404 Not Found 오류를 보고합니다.
			// 응답 처리는 미들웨어가 담당합니다.
			httperror.ReportNotFound(r, "지정된 ID의 사용자를 찾을 수 없습니다.")
			return
		}

		fmt.Fprintln(w, "사용자 데이터가 여기에 출력됩니다.")
	})
    mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
        // 403 Forbidden 오류를 보고합니다.
        httperror.ReportForbidden(r)
    })


	// Mux를 오류 보고 미들웨어로 감쌉니다.
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", httperror.ErrorReporterMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}
```

이 서버를 실행하고 `http://localhost:8080/users` 주소로 접속하면, 다음과 같은 JSON 응답을 받게 됩니다.

```json
{
	"status": 404,
	"message": "지정된 ID의 사용자를 찾을 수 없습니다."
}
```

#### 사용자 정의 오류 핸들러

`NewErrorReporterMiddleware`를 사용하면 자신만의 오류 처리 로직을 구현할 수도 있습니다.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevNewbie1826/httperror"
)

// 오류를 로깅하고 커스텀 응답을 작성하는 핸들러
func myCustomErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("오류 발생: %v", err)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if httpErr, ok := err.(*httperror.HttpError); ok {
		w.WriteHeader(httpErr.Status)
		fmt.Fprintf(w, `{"errorCode": %d, "errorMessage": "%s"}`, httpErr.Status, httpErr.Message)
		return
	}

	// HttpError가 아닌 다른 타입의 오류를 위한 처리
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, `{"errorCode": 500, "errorMessage": "예상치 못한 오류가 발생했습니다."}`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
        // 400 Bad Request 오류를 보고합니다.
		httperror.ReportBadRequest(r, "'id' 쿼리 파라미터가 필요합니다.")
	})

	// 커스텀 핸들러를 사용하여 새로운 미들웨어 인스턴스를 생성합니다.
	errorMiddleware := httperror.NewErrorReporterMiddleware(myCustomErrorHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", errorMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}
```

#### `chi` 라우터와 함께 사용하기

`chi`와 같은 라우터를 사용할 때, 라우팅 오류(404 Not Found, 405 Method Not Allowed 등)를 애플리케이션 오류와 동일한 방식으로 처리하고 싶을 수 있습니다. `chi`는 이러한 경우를 위한 사용자 정의 핸들러 설정을 지원합니다.

`Report...` 함수를 사용하는 커스텀 핸들러를 설정하여 `httperror`를 통합할 수 있습니다.

**중요:** `httperror.ErrorReporterMiddleware`는 `NotFound` 및 `MethodNotAllowed` 핸들러가 컨텍스트에 접근할 수 있도록 라우트보다 먼저 등록되어야 합니다.

```go
package main

import (
	"log"
	"net/http"

	"github.com/DevNewbie1826/httperror"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// ErrorReporterMiddleware는 라우터의 핸들러들보다 먼저 적용되어야 합니다.
	// 이렇게 해야 NotFound, MethodNotAllowed 핸들러에서 컨텍스트를 사용할 수 있습니다.
	r.Use(httperror.ErrorReporterMiddleware)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// chi의 기본 NotFound 핸들러를 오버라이드합니다.
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		httperror.ReportNotFound(r, "요청하신 리소스 경로는 존재하지 않습니다.")
	})

	// chi의 기본 MethodNotAllowed 핸들러를 오버라이드합니다.
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		httperror.ReportMethodNotAllowed(r)
	})

	log.Println("Server starting on :8080 with chi")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
```

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
