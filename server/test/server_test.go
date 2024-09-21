package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/codescalersinternships/Datetime-Server-Fatma-Ebrahim/pkg/ginserver"
	"github.com/codescalersinternships/Datetime-Server-Fatma-Ebrahim/pkg/httpserver"
)

func TestHTTPServer(t *testing.T) {
	t.Run("successfull request with plain text", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/datetime", nil)
		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()
		request.Header.Set("Content-Type", "plain text")

		httpserver.HTTPHandler(response, request)
		got := strings.Split(response.Body.String(), ":")[0]
		want := strings.Split(time.Now().Format("Monday 02-01-2006 15:04:05"), ":")[0]
		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, response.Result().StatusCode)
		}
		if got != want {
			t.Errorf("expected %q, want %q", got, want)
		}
	})

	t.Run("successfull request with json", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/datetime", nil)
		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()
		request.Header.Set("Content-Type", "application/json")

		httpserver.HTTPHandler(response, request)
		want := strings.Split(time.Now().Format("Monday 02-01-2006 15:04:05"), ":")[0]
		var result map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Fatal(err)
		}

		got := strings.Split(result["datetime"].(string), ":")[0]
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("unsuccessfull request", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()
		httpserver.HTTPHandler(response, request)
		if response.Result().StatusCode != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, response.Result().StatusCode)
		}
	})

}

func TestGinServer(t *testing.T) {
	t.Run("successfull request with plain text", func(t *testing.T) {
		router := ginserver.GinHandler()
		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/datetime", nil)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "plain text")
		router.ServeHTTP(response, request)
		if response.Code != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, response.Code)
		}
		got := strings.Split(response.Body.String(), ":")[0]
		want := strings.Split(time.Now().Format("Monday 02-01-2006 15:04:05"), ":")[0]
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("successfull request with application/json", func(t *testing.T) {
		router := ginserver.GinHandler()
		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/datetime", nil)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(response, request)
		if response.Code != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, response.Code)
		}
		want := strings.Split(time.Now().Format("Monday 02-01-2006 15:04:05"), ":")[0]
		var result map[string]interface{}
		err = json.Unmarshal(response.Body.Bytes(), &result)
		if err != nil {
			t.Fatal(err)
		}

		got := strings.Split(result["datetime"].(string), ":")[0]
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("unsuccessfull request", func(t *testing.T) {
		router := ginserver.GinHandler()
		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "", nil)
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(response, request)
		if response.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, response.Code)
		}
	})

}
