package config

import (
	"time"
)

// place holder for now
type HttpClientConfiguration struct {
	TimeOut            int
	RetryMax           int
	RetryWaitMin       int
	MaxIdleConnections int
	BaseApiUrl         string
	Upload             string
	Meta               string
}

// Define all interfaces for this struct
type IHttpClientConfiguration interface {
	GetTimeOut() time.Duration
	GetRetryMax() int
	GetRetryWaitMin() int
	GetMaxIdleConnections() int
	GetBaseApiUrl() string
	GetUploadUri() string
	GetMetaUri() string
}

/*
   httpClient.RetryMax = conf.HttpClient.GetRetryMax()
   httpClient.RetryWaitMin = conf.HttpClient.GetRetryWaitMin()
   httpClient.HTTPClient.Timeout = time.Second * 10
*/

// Implementation
func (h HttpClientConfiguration) GetTimeOut() int {
	return h.TimeOut
}

func (h HttpClientConfiguration) GetRetryMax() int {
	return h.RetryMax
}

func (h HttpClientConfiguration) GetRetryWaitMin() int {
	return h.RetryWaitMin
}

func (h HttpClientConfiguration) GetMaxIdleConnections() int {
	return h.MaxIdleConnections
}

func (h HttpClientConfiguration) GetBaseApiUrl() string {
	return h.BaseApiUrl
}

func (h HttpClientConfiguration) GetUploadUri() string {
	return h.Upload
}

func (h HttpClientConfiguration) GetMetaUri() string {
	return h.Meta
}
