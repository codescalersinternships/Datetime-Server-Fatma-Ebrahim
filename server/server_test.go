package main

import (
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
		want := strings.Split(time.Now().Format("Monday 02-01-2006 15:04:05 \n"), ":")[0]
		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, response.Result().StatusCode)
		}
		if got != want {
			t.Errorf("expected %q, want %q", got, want)
		}
	})

	t.Run("bad request", func(t *testing.T) {
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
