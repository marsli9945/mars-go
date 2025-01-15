package marsHttp

import (
	"encoding/json"
	"fmt"
	"github.com/marsli9945/mars-go/marsHttp"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 模拟 HTTP 服务器
func testServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/get":
			_, err := fmt.Fprintln(w, "GET response")
			if err != nil {
				return
			}
		case "/post":
			_, err := fmt.Fprintln(w, "POST response")
			if err != nil {
				return
			}
		case "/error":
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		case "/json":
			jsonData := map[string]string{"key": "value"}
			err := json.NewEncoder(w).Encode(jsonData)
			if err != nil {
				return
			}
		default:
			http.NotFound(w, r)
		}
	}))
}

func TestGet_Success(t *testing.T) {
	server := testServer()
	defer server.Close()

	url := server.URL + "/get"
	body, err := marsHttp.Get(url)
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if body != "GET response\n" {
		t.Errorf("Unexpected response: %s", body)
	}
}

func TestGet_Error(t *testing.T) {
	server := testServer()
	defer server.Close()

	url := server.URL + "/error"
	body, err := marsHttp.Get(url)
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	if body != "" {
		t.Errorf("Unexpected response: %s", body)
	}
}

func TestPost_Success(t *testing.T) {
	server := testServer()
	defer server.Close()

	url := server.URL + "/post"
	data := map[string]any{"key": "value"}
	body, err := marsHttp.Post(url, data)
	if err != nil {
		t.Errorf("Post failed: %v", err)
	}
	if body != "POST response\n" {
		t.Errorf("Unexpected response: %s", body)
	}
}

func TestPostAndHeaderForStruct_Success(t *testing.T) {
	server := testServer()
	defer server.Close()

	url := server.URL + "/json"
	headers := map[string]string{"Authorization": "Bearer token"}
	var result map[string]string
	err := marsHttp.PostAndHeaderForStruct(url, nil, headers, &result)
	if err != nil {
		t.Errorf("PostAndHeaderForStruct failed: %v", err)
	}
	if result["key"] != "value" {
		t.Errorf("Unexpected result: %v", result)
	}
}

func TestGetAndHeaderForStruct_Success(t *testing.T) {
	server := testServer()
	defer server.Close()

	url := server.URL + "/json"
	headers := map[string]string{"Authorization": "Bearer token"}
	var result map[string]string
	err := marsHttp.GetAndHeaderForStruct(url, headers, &result)
	if err != nil {
		t.Errorf("GetAndHeaderForStruct failed: %v", err)
	}
	if result["key"] != "value" {
		t.Errorf("Unexpected result: %v", result)
	}
}
