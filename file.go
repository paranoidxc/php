package php

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

// FileGetContent 根据传入的参数类型，进行文件读取或者网络请求，并返回内容
func FileGetContent(source string) ([]byte, error) {
	// Check if source is a file
	if _, err := os.Stat(source); err == nil {
		// File exists, read contents from file
		return fileContent(source)
	} else if !os.IsNotExist(err) {
		// File error occurred
		return nil, err
	}

	if strings.HasPrefix(source, "http") {
		// Source is not a file, treat it as URL and perform network request
		return urlContent(source)
	}

	return nil, errors.New("file not found")
}

// urlContent 发送 HTTP GET 请求并返回响应体的内容
func urlContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// fileContent 读取文件内容并返回
func fileContent(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
