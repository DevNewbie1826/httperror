# httperror

[![Go Report Card](https://goreportcard.com/badge/github.com/DevNewbie1826/httperror)](https://goreportcard.com/report/github.com/DevNewbie1826/httperror)
[![GoDoc](https://godoc.org/github.com/DevNewbie1826/httperror?status.svg)](https://godoc.org/github.com/DevNewbie1826/httperror)
[![codecov](https://codecov.io/gh/DevNewbie1826/httperror/graph/badge.svg)](https://codecov.io/gh/DevNewbie1826/httperror)

A simple, lightweight Go package for centralized HTTP error handling, using helper functions to generate standard HTTP error responses directly.

---

## English

### 📖 Overview

`httperror` provides a simple way to handle errors in your Go web application without complex middleware configurations. It offers a set of helper functions that directly write standardized JSON (or HTML) error responses to the `http.ResponseWriter`. It also supports a global error handler configuration for custom error rendering.

### 🚀 Installation

```bash
go get github.com/DevNewbie1826/httperror
```

### ✨ Features

-   **Direct Usage**: No middleware required. Call `httperror.NotFound(w, r)` directly in your handlers.
-   **Standardized Responses**: Default error handler responds with JSON (or HTML for browsers) automatically.
-   **Customizable**: You can replace the default error handler globally using `httperror.SetErrorHandler`.
-   **Comprehensive Helpers**: Covers almost all standard HTTP error status codes (e.g., `httperror.BadRequest`, `httperror.Forbidden`, etc.).

### 💡 Usage

#### Basic Example

Simply import the package and use the helper functions in your HTTP handlers.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevNewbie1826/httperror"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// Some logic to find a user...
		userFound := false

		if !userFound {
			// Directly respond with a 404 Not Found error.
			// This writes the status code and the JSON body to the ResponseWriter.
			httperror.NotFound(w, r, "User with the specified ID was not found.")
			return
		}

		fmt.Fprintln(w, "User data would be here.")
	})
    
    http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
        // Respond with a 403 Forbidden error using the default message.
        httperror.Forbidden(w, r)
    })

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
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

You can provide your own custom error handling logic globally using `SetErrorHandler`. This is useful if you want to render custom HTML error pages or change the JSON structure.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevNewbie1826/httperror"
)

func main() {
	// Set a custom global error handler.
	httperror.SetErrorHandler(func(w http.ResponseWriter, r *http.Request, err error) {
		// You can type-assert to access the status code if needed
		status := http.StatusInternalServerError
		message := "Internal Server Error"
		
		if httpErr, ok := err.(*httperror.HttpError); ok {
			status = httpErr.Status
			message = httpErr.Message
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		fmt.Fprintf(w, `{"error": true, "code": %d, "msg": "%s"}`, status, message)
	})

	http.HandleFunc("/oops", func(w http.ResponseWriter, r *http.Request) {
		httperror.BadRequest(w, r, "Something went wrong!")
	})

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
```

---

## 한국어

### 📖 개요

`httperror`는 복잡한 미들웨어 설정 없이 Go 웹 애플리케이션의 오류를 처리할 수 있는 간단하고 가벼운 패키지입니다. `http.ResponseWriter`에 표준화된 JSON(또는 브라우저의 경우 HTML) 오류 응답을 직접 작성하는 헬퍼 함수들을 제공합니다. 또한, 전역 오류 핸들러 설정을 통해 커스텀 렌더링 로직을 적용할 수 있습니다.

### 🚀 설치

```bash
go get github.com/DevNewbie1826/httperror
```

### ✨ 주요 기능

-   **직관적인 사용**: 미들웨어가 필요 없습니다. 핸들러에서 `httperror.NotFound(w, r)`와 같이 직접 호출하세요.
-   **표준화된 응답**: 기본 오류 핸들러가 자동으로 JSON(또는 브라우저 요청 시 HTML)으로 응답합니다.
-   **커스터마이징**: `httperror.SetErrorHandler`를 사용하여 기본 오류 핸들러를 전역적으로 교체할 수 있습니다.
-   **포괄적인 헬퍼**: 거의 모든 표준 HTTP 오류 상태 코드를 지원합니다 (예: `httperror.BadRequest`, `httperror.Forbidden` 등).

### 💡 사용법

#### 기본 예제

패키지를 import 하고 HTTP 핸들러 내에서 헬퍼 함수를 직접 사용하시면 됩니다.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevNewbie1826/httperror"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// 사용자를 찾는 로직이 있다고 가정...
		userFound := false

		if !userFound {
			// 404 Not Found 오류로 즉시 응답합니다.
			// 이 함수가 상태 코드와 JSON 본문을 ResponseWriter에 작성합니다.
			httperror.NotFound(w, r, "지정된 ID의 사용자를 찾을 수 없습니다.")
			return
		}

		fmt.Fprintln(w, "사용자 데이터가 여기에 출력됩니다.")
	})
    
    http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
        // 기본 메시지를 사용하여 403 Forbidden 오류를 응답합니다.
        httperror.Forbidden(w, r)
    })

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
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

`SetErrorHandler`를 사용하면 전역 오류 처리 로직을 직접 정의할 수 있습니다. 커스텀 HTML 오류 페이지를 렌더링하거나 JSON 구조를 변경하고 싶을 때 유용합니다.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevNewbie1826/httperror"
)

func main() {
	// 커스텀 전역 오류 핸들러 설정
	httperror.SetErrorHandler(func(w http.ResponseWriter, r *http.Request, err error) {
		// 필요하다면 타입 단언(type assertion)을 통해 상태 코드에 접근할 수 있습니다.
		status := http.StatusInternalServerError
		message := "Internal Server Error"
		
		if httpErr, ok := err.(*httperror.HttpError); ok {
			status = httpErr.Status
			message = httpErr.Message
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		fmt.Fprintf(w, `{"error": true, "code": %d, "msg": "%s"}`, status, message)
	})

	http.HandleFunc("/oops", func(w http.ResponseWriter, r *http.Request) {
		httperror.BadRequest(w, r, "무언가 잘못되었습니다!")
	})

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
```

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.