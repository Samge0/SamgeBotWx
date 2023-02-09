package u_http

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"sync"
	"time"
)

var client *http.Client
var once sync.Once

// GetClient 获取http的客户端
func GetClient() *http.Client {
	once.Do(func() {
		client = &http.Client{
			//Timeout: 4800 * time.Millisecond,
			Timeout: 120 * time.Second,
		}
	})
	return client
}

// Get 网络请求-Get方式
func Get(url string, headers map[string]string) (*[]byte, error) {
	return DoReq(url, headers, nil)
}

// Post 网络请求-Post方式
func Post(url string, headers map[string]string, params any) (*[]byte, error) {
	return DoReq(url, headers, params)
}

// DoReq 执行请求
func DoReq(url string, headers map[string]string, params any) (*[]byte, error) {
	var request *http.Request
	var err error

	// 判断是 Post 还是 Get
	if params != nil {
		dataByte, err := json.Marshal(params)
		if err != nil {
			errorMsg := fmt.Sprintf("[json.Marshal]失败, 错误内容: %s\n", err.Error())
			return nil, errors.New(errorMsg)
		}
		bodyReader := bytes.NewReader(dataByte)
		request, err = http.NewRequestWithContext(context.Background(), http.MethodPost, url, bodyReader)
	} else {
		request, err = http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	}

	if err != nil {
		errorMsg := fmt.Sprintf("[http.NewRequestWithContext]失败, 错误内容: %s\n", err.Error())
		return nil, errors.New(errorMsg)
	}

	// 设置请求头
	if headers != nil {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}

	resp, err := GetClient().Do(request)
	if err != nil {
		errorMsg := fmt.Sprintf("[client.Do]失败, 错误内容: %s\n", err.Error())
		return nil, errors.New(errorMsg)
	}
	fmt.Println("statusCode: ", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorMsg := fmt.Sprintf("[io.ReadAll]失败, 错误内容: %s\n", err.Error())
		return nil, errors.New(errorMsg)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	return &body, nil
}
