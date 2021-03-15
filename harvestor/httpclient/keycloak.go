package httpclient

import (
	"context"
	"crypto/tls"
	"github.com/Nerzal/gocloak/v8"
	"harvestor/config"
	l "harvestor/logger"
	"time"
)

// watching ExpiresIn on the token
var start time.Time
var keycloak gocloak.GoCloak
var token *gocloak.JWT

func getNewKeycloak() {
	// init start time when we get first token
	start = time.Now()
	// Getting config
	conf := config.GetConf()
	// Getting logger
	logger := l.NewLogger()

	// !!! Important !!!
	// host name for example `keycloak` should be the same as an object-store-api uses
	// please check in object-store-api Env var : KEYCLOAK_HOSTNAME
	// tokens are issued based on the domain of the Keycloak when Keycloak is called
	// all consumers of Keycloak have to use the same Keycloak domain to get tokens
	// for example:
	// for local development add `127.0.0.1 keycloak` to your hosts
	// now you can reference `keycloak` as a domain in your calls
	logger.Info("About to init new Keycloak ...")
	keycloak = gocloak.NewClient(conf.Keycloak.GetHost())
	restyClient := keycloak.RestyClient()
	restyClient.SetDebug(conf.Keycloak.IsDebug())
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	logger.Info("New Keycloak has been init !!!")

	// Lets login on behalf of an admin client
	var err error
	ctx := context.Background()
	logger.Info("About to login on behalf of an admin client ...")
	token, err = DinaLogin(ctx, "user", "user", "dina")
	if err != nil {
		logger.Fatal("Something wrong with the credentials or url", err)
	}

	logger.Info("login on behalf of an admin client success !!!")

	//logger.Info("token.AccessToken is : ", token.AccessToken)
	//logger.Info("token.ExpiresIn is : ", token.ExpiresIn)
	//logger.Info("token.RefreshToken is : ", token.RefreshToken)
	//logger.Info("token.RefreshExpiresIn is : ", token.RefreshExpiresIn)
}

// GetAccessToken for the request header
func GetAccessToken() string {
	// checking if keycloak has never been initialized
	if keycloak == nil {
		getNewKeycloak()
		return token.AccessToken
	}
	checkToken()
	return token.AccessToken
}

// Checking if we need a new token due to a close expiry date of Access Token
func checkToken() {
	// Getting config
	conf := config.GetConf()
	// Getting logger
	logger := l.NewLogger()
	// getting time NOW
	t := time.Now()
	// compare with when we started
	elapsed := t.Sub(start)
	// compute time left before our Access token expires
	timeLeft := token.ExpiresIn - int(elapsed/time.Second)
	logger.Debug("seconds left for Access Token : ", timeLeft)
	// check if need to get new token
	if timeLeft < conf.Keycloak.GetNewTokenBefore() {
		getNewKeycloak()
	}
}

// DinaLogin performs a login on behalf of an admin client
func DinaLogin(ctx context.Context, username, password, realm string) (*gocloak.JWT, error) {
	conf := config.GetConf()
	return keycloak.GetToken(ctx, realm, gocloak.TokenOptions{
		ClientID:  StringP(conf.Keycloak.GetAdminClientID()),
		GrantType: StringP(conf.Keycloak.GetGrantType()),
		Username:  &username,
		Password:  &password,
	})
}

// StringP returns a pointer of a string variable
func StringP(value string) *string {
	return &value
}
