package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceReturnsOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.NotEmpty(t, responseRecorder.Body)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

}

func TestWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=samara", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := "wrong city value"
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, expected, responseRecorder.Body.String())
}

func TestWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := 4
	result := responseRecorder.Body.String()
	list := strings.Split(result, ",")
	assert.Equal(t, len(list), expected)

}
