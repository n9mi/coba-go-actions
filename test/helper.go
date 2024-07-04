package test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gofiber/fiber"
)

func newRequest(method string, url string, requestBody string) *http.Request {
	request := httptest.NewRequest(method, url, strings.NewReader(requestBody))
	request.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return request
}
