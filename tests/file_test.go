package php

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/paranoidxc/php"
)

func TestFileGetContentByUrl(t *testing.T) {
	// Mock a simple HTTP server for testing
	server := mockHTTPServer()
	defer server.Close()

	url := server.URL + "/test"
	expected := "Hello, World!"
	body, err := php.FileGetContent(url)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	if string(body) != expected {
		t.Errorf("Expected: %s, Got: %s", expected, string(body))
	}
}

func TestFileGetContentByFile(t *testing.T) {
	// Test case: Reading contents of a file
	currentDir, _ := os.Getwd()
	filename := currentDir + "/../testdata/test.txt"
	expected := "Hello, World!\n"
	content, err := php.FileGetContent(filename)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	if string(content) != expected {
		t.Errorf("Expected: %s, Got: %s", expected, string(content))
	}

	// Test case: Trying to read a non-existent file
	nonExistentFile := currentDir + "/../testdata/nonexistent.txt"
	_, err = php.FileGetContent(nonExistentFile)
	AssertEqual(t, err.Error(), "file not found")
	if err == nil {
		t.Errorf("Expected error for non-existent file, got none")
	}
}

func mockHTTPServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	server := httptest.NewServer(mux)
	return server
}
