package http

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	_ "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"id": 1, "name": "John", "info": "Developer"}`))
}

func TestMakeHTTP(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	fmt.Println(server.URL)

	// time.Sleep(50 * time.Second)

	want := &Response{
		ID:   1,
		Name: "John",
		Info: "Developer",
	}

	t.Run("Happy server response", func(t *testing.T) {
		resp, err := MakeHTTPCall(server.URL)

		if !reflect.DeepEqual(resp, want) {
			t.Errorf("expected (%v), got (%v)", want, resp)
		}

		if !errors.Is(err, nil) {
			t.Errorf("expected (%v), got (%v)", nil, err)
		}
	})
}
