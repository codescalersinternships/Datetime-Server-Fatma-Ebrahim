package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestDateTimeServer(t *testing.T) {
	t.Run("successfull request", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/datetime", nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()
		GetDateAndTime(response, request)
		got := strings.Split(response.Body.String(), ":")[0]
		want := strings.Split(time.Now().Format("Monday 02-01-2006 15:04:05"), ":")[0]
		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, response.Result().StatusCode)
		}
		if got != want {
			t.Errorf("expected %q, want %q", got, want)
		}
	})

	t.Run("unsuccessfull request", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()
		GetDateAndTime(response, request)
		if response.Result().StatusCode != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, response.Result().StatusCode)
		}
	})

}

func TestPingRoute(t *testing.T) {

	t.Run("successfull request", func(t *testing.T) {
		router := GinGetDateAndTime()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/datetime", nil)
		router.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, w.Code)
		}
		want := strings.Split(time.Now().Format("Monday 02-01-2006 15:04:05"), ":")[0]
		var result map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &result)

		got := strings.Split(result["datetime"].(string), ":")[0]
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("unsuccessfull request", func(t *testing.T) {
		router := GinGetDateAndTime()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "", nil)
		router.ServeHTTP(w, req)
		if w.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})

}
