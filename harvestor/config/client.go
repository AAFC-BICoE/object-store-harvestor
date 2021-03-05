package config

// place holder for now
type HttpClientConfiguration struct {
	TimeOut    int
	BaseApiUrl string
	Uri        string
}

// Define all interfaces for this struct
type IHttpClientConfiguration interface {
	GetTimeOut() int
	GetBaseApiUrl() string
	GetUri() string
}

// Implementation
func (h HttpClientConfiguration) GetTimeOut() int {
	return h.TimeOut
}

func (h HttpClientConfiguration) GetBaseApiUrl() string {
	return h.BaseApiUrl
}

func (h HttpClientConfiguration) GetUri() string {
	return h.Uri
}
