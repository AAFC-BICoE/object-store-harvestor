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
	UploadGroup        string
	Meta               string
	ManagedMeta        string
	Derivative         string
}

// Define all interfaces for this struct
type IHttpClientConfiguration interface {
	GetTimeOut() time.Duration
	GetRetryMax() int
	GetRetryWaitMin() int
	GetMaxIdleConnections() int
	GetBaseApiUrl() string
	GetUploadUri() string
	GetUploadGroup() string
	GetMetaUri() string
	GetManagedMetaUri() string
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

func (h HttpClientConfiguration) GetManagedMetaUri() string {
	return h.ManagedMeta
}

func (h HttpClientConfiguration) GetDerivativeUri() string {
	return h.Derivative
}
