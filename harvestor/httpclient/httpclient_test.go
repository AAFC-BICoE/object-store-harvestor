package httpclient

import (
	"github.com/stretchr/testify/assert"
	"harvestor/config"
	"testing"
	"time"
)

func TestInitHttpClient(t *testing.T) {
	// config file
	file := "../harvestor_config.yml"
	// load current config
	config.Load(file)
	// get conf
	conf := config.GetConf()
	// init http client
	InitHttpClient()
	// asserting
	assert.NotNil(t, httpClient)
	assert.Equal(t, conf.HttpClient.GetRetryMax(), httpClient.RetryMax)
	assert.Equal(t, time.Duration(conf.HttpClient.GetRetryWaitMin())*time.Second, httpClient.RetryWaitMin)
	assert.Equal(t, time.Duration(conf.HttpClient.GetTimeOut())*time.Second, httpClient.HTTPClient.Timeout)
}
