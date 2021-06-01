// Package config provides a functionality to read from yml config file
// and provides values for each key in the file
// the package is splitted into multiple for simplicity and readability
// Http Client config
package config

// Http client config struct
type HttpClientConfiguration struct {
	TimeOut            int    // how many seconds for Http client to timeout
	RetryMax           int    // how many times to retry
	RetryWaitMin       int    // how many seconds to wait for retry
	MaxIdleConnections int    // how many max idle connection to keep on init
	BaseApiUrl         string // Base API URL for object-store-api
	Upload             string // Upload API resource
	UploadGroup        string // Upload Group
	Meta               string // Meta API resource
	Derivative         string // Derivative API resource
}

// Define all interfaces for this struct
type IHttpClientConfiguration interface {
	GetTimeOut() int
	GetRetryMax() int
	GetRetryWaitMin() int
	GetMaxIdleConnections() int
	GetBaseApiUrl() string
	GetUploadUri() string
	GetUploadGroup() string
	GetMetaUri() string
	GetDerivativeUri() string
}

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

func (h HttpClientConfiguration) GetUploadGroup() string {
	return h.UploadGroup
}

func (h HttpClientConfiguration) GetMetaUri() string {
	return h.Meta
}

func (h HttpClientConfiguration) GetDerivativeUri() string {
	return h.Derivative
}
