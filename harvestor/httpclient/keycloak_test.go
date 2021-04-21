package httpclient

import (
	"github.com/liamylian/jsontime"
	"github.com/stretchr/testify/assert"
	"harvestor/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakeToken struct {
	AccseesToken     string `json:"access_token"`
	IdToken          string `json:"id_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

func TestKeyCloak(t *testing.T) {
	// custom json for all time formats
	var json = jsontime.ConfigWithCustomTimeFormat
	// config file
	file := "../harvestor_config.yml"
	// load config
	config.Load(file)
	// get conf
	conf := config.GetConf()
	// getting fake token struct to return from mock server
	ft := FakeToken{
		AccseesToken:     "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJNRllObnhMM1N0ZkU5SWhkYWJVd29ELW0yR3ZFUXhxTldDZk56ZDVMeXZZIn0.eyJleHAiOjE2MTg5NjQ2NzYsImlhdCI6MTYxODk2NDM3NiwianRpIjoiNjI5NmQ5ZGUtZjJkOS00YTg3LTllNzQtZjMzMDJlYzdlOWY1IiwiaXNzIjoiaHR0cDovL2tleWNsb2FrOjgwODAvYXV0aC9yZWFsbXMvZGluYSIsInN1YiI6IjAzZjM2NzQ5LThmNzQtNDY5ZC04YTc5LThhYzZjOTY4ODhkNiIsInR5cCI6IkJlYXJlciIsImF6cCI6Im9iamVjdHN0b3JlIiwic2Vzc2lvbl9zdGF0ZSI6ImFiY2E4NDE1LTU4ZTUtNDY5Mi05NTVkLTY5YjNjZTFmZGIzYyIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiKiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiY29sbGVjdGlvbi1tYW5hZ2VyIiwidXNlciJdfSwic2NvcGUiOiJlbWFpbCBkaW5hLWFnZW50IHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImFnZW50LWlkZW50aWZpZXIiOiJiYTA1ZDQ2Yi02Yzc1LTRjOTItODJkNi1lNzIyMzdkOTk4MmYiLCJncm91cHMiOlsiL2NuYy9jb2xsZWN0aW9uLW1hbmFnZXIiXSwicHJlZmVycmVkX3VzZXJuYW1lIjoiY25jLWNtIn0.S4ct4hSH5yE07-vGhrXiwPfThRo2fngW35PMhzyZgsUDl3UvJ5yMFIjRAeWQbtuvmc6VWL0Gee-R7th22gWdlEPZMcMBungDOifPVpWL2aMe8048rfHqy43D_mh6JC9MK797Iwxnjs71eBBFFmyl0WjUSSrDu7oRHSStjv9467ao-Vra-JpTVF7JPwtWnoRGjiPgdeuRr8RG56f7_u2m0ja7uPvlatXH3chy6qESIJ50uLmaYI4X5HJpJ7HxImI3wBHefXuoOuRyzeI5sktr_DLh32Y1aLB6sLKusKGVvAQt6TXrkq6fNlP8ZfeFIQb-KlVrpMGMEuesfonY6BlgVQ",
		IdToken:          "",
		ExpiresIn:        300,
		RefreshExpiresIn: 1800,
		RefreshToken:     "eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJkOGNkYzc0MS00OTgxLTQ3OWItYmZmMS1lYjg2NDdiZjI1MTcifQ.eyJleHAiOjE2MTg5NjYxNzYsImlhdCI6MTYxODk2NDM3NiwianRpIjoiMmRmY2JiNWEtOGNjYi00NDY1LTk5ZmUtYWQyNmNhMzc4OTdjIiwiaXNzIjoiaHR0cDovL2tleWNsb2FrOjgwODAvYXV0aC9yZWFsbXMvZGluYSIsImF1ZCI6Imh0dHA6Ly9rZXljbG9hazo4MDgwL2F1dGgvcmVhbG1zL2RpbmEiLCJzdWIiOiIwM2YzNjc0OS04Zjc0LTQ2OWQtOGE3OS04YWM2Yzk2ODg4ZDYiLCJ0eXAiOiJSZWZyZXNoIiwiYXpwIjoib2JqZWN0c3RvcmUiLCJzZXNzaW9uX3N0YXRlIjoiYWJjYTg0MTUtNThlNS00NjkyLTk1NWQtNjliM2NlMWZkYjNjIiwic2NvcGUiOiJlbWFpbCBkaW5hLWFnZW50IHByb2ZpbGUifQ.1xmBIM5n9f3td30WOEBZNixZYNU_rlBijI236Rt1ZiE",
		TokenType:        "Bearer",
		NotBeforePolicy:  0,
		SessionState:     "abca8415-58e5-4692-955d-69b3ce1fdb3c",
		Scope:            "email dina-agent profile",
	}
	// struct to json
	jData, err := json.Marshal(ft)
	// asserting
	assert.Nil(t, err)
	// create mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	}))
	// defer closing the server until the end of the test
	defer ts.Close()

	// define mock serve base pai url
	conf.Keycloak.Host = ts.URL
	// disable keycloak for now
	conf.Keycloak.Enabled = true

	// init http client
	InitHttpClient()
	// testing upload against mock server
	at := GetAccessToken()
	checkToken()
	// asserting
	assert.NotNil(t, at)
}
