// sidecar for managed attributes
// based on existing metadata
// for more details please refer to https://gist.github.com/Jkeyuk/6cfc728665fac2dced079c770cf74fca

package httpclient

import (
	"bytes"
	c "github.com/hashicorp/go-retryablehttp"
	"github.com/liamylian/jsontime"
	"harvestor/config"
	"harvestor/db"
	l "harvestor/logger"
	"harvestor/walker"
	"io"
	"net/http"
)

// structs for http POST
// | | | attributes | | |
type PostSidecarValue struct {
	Value string `json:"value"`
}
type PostSidecarAttributes struct {
	Values map[string]*PostSidecarValue `json:"values"`
}
type PostSidecarDerivativeAttributes struct {
	DerivativeType string `json:"derivativeType"`
	DcType         string `json:"dcType"`
	Bucket         string `json:"bucket"`
	FileIdentifier string `json:"fileIdentifier"`
}

// | | | relationships | | |
type PostSidecarRelationships struct {
	PostSidecarRelationshipData *PostSidecarRelationshipData `json:"metadata"`
}
type PostSidecarDerivativeRelationships struct {
	PostSidecarRelationshipData *PostSidecarRelationshipData `json:"acDerivedFrom"`
}
type PostSidecarRelationshipData struct {
	PostSidecarMetaData PostSidecarMetaData `json:"data"`
}
type PostSidecarMetaData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Post sidecar wrapper
type PostSidecarData struct {
	Type                     string                   `json:"type"`
	PostSidecarAttributes    PostSidecarAttributes    `json:"attributes"`
	PostSidecarRelationships PostSidecarRelationships `json:"relationships"`
}
type PostSidecarDerivativeData struct {
	Type                     string                             `json:"type"`
	PostSidecarAttributes    PostSidecarDerivativeAttributes    `json:"attributes"`
	PostSidecarRelationships PostSidecarDerivativeRelationships `json:"relationships"`
}

// Post sidecar data wrapper
type PostSidecar struct {
	PostSidecarData PostSidecarData `json:"data"`
}
type PostSidecarDerivative struct {
	PostSidecarData PostSidecarDerivativeData `json:"data"`
}

func postSideCarDerivative(sidecar *db.Sidecar) error {
	// custom json for all time formats
	var json = jsontime.ConfigWithCustomTimeFormat
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()
	// Need to upload managed meta against Original
	// Get Original file from current sidecar
	originalFile := sidecar.GetOriginalFile()
	// Get Original meta from DB
	meta, err := db.GetMetaByFile(&originalFile)
	// Checking on errors
	if err != nil {
		logger.Fatal("Fatal Error : postSideCarDerivative on db.GetMetaByFile err : ", err)
	}
	// get derivative file
	derivativeFile := sidecar.GetDerivativeFile()
	// get derivative upload
	upload, err := db.GetUploadByFile(&derivativeFile)
	// checking on errors
	if err != nil {
		logger.Fatal("Fatal Error : postSideCarDerivative on db.GetUploadByFile err : ", err)
	}
	// define url for derivative
	url := conf.HttpClient.GetBaseApiUrl() + conf.HttpClient.GetDerivativeUri()
	logger.Debug("post derivative url : ", url)
	// Read the content of yml sidecar file
	scf, err := walker.GetSidecarYmlFile(sidecar)
	// Loading sidecar and checking if derivative is set
	if err != nil {
		return err
	}
	logger.Debug("post derivative : original file : ", logger.PrettyGoStruct(originalFile))
	logger.Debug("post derivative : derivative file : ", logger.PrettyGoStruct(derivativeFile))
	logger.Debug("post derivative : original meta : ", logger.PrettyGoStruct(meta))
	logger.Debug("post derivative : derivative upload : ", logger.PrettyGoStruct(upload))
	logger.Debug("post derivative : sidecar : ", logger.PrettyGoStruct(scf))

	// populate post data struct
	postData := getDerivativePostData(upload, meta)
	logger.Debug("post derivative : postData : ", logger.PrettyGoStruct(postData))
	payload, err := json.Marshal(*postData)
	logger.Debug(" post derivative payload : ", string(payload))
	if err != nil {
		logger.Error("json.Marshal fail details :", err)
		return err
	}
	// make new request
	req, err := c.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		logger.Error("New request Errors :", err)
		return err
	}
	// set headers
	req.Header.Set("Content-Type", "application/vnd.api+json")
	// check if we need Authorization
	if conf.Keycloak.IsEnabled() {
		var bearer = "Bearer " + GetAccessToken()
		req.Header.Set("Authorization", bearer)
	}
	// custom header for https://www.crnk.io/releases/stable/documentation/
	req.Header.Add("crnk-compact", "true")
	// make reqoest
	resp, err := httpClient.Do(req)
	// checking on errors from request
	if err != nil {
		logger.Error(" post derivative fail details :", err)
		return err
	}
	// TODO Maybe we need a common package for HTTP response status codes
	// Check on response status 401
	if resp.StatusCode == http.StatusUnauthorized {
		logger.Fatal("Error : You are Unauthorized for derivative, Please check your config file")
	}
	// Check on response status 403
	if resp.StatusCode == http.StatusForbidden {
		logger.Fatal("Error : You are Forbidden from derivative, Please check your config file")
	}
	// Check on response status 201
	if resp.StatusCode == http.StatusCreated {
		// close the body when done
		defer resp.Body.Close()
		// read the body
		b, err := io.ReadAll(resp.Body)
		logger.Debug(" post derivative response body : ", string(b))
		if err != nil {
			logger.Error(" error on read body : ", err)
			return err
		}
		logger.Debug(" derivative record has been posted : ", logger.PrettyGoStruct(postData))

	} else {
		// all other use cases are not allowed
		// app has to stop
		// something is really not right here
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Fatal(" error on read body : ", err)
		}
		logger.Fatal("Error : Status code : (", resp.StatusCode, ") details : ", string(b))
	}
	return err
}

// Helper functionis to build full struct for post data
func getMetaPostData(key string, value string, meta *db.Meta) *PostSidecar {
	// value to assign
	postValue := &PostSidecarValue{
		value,
	}
	// uuid of the managed attribute with value
	postAttributes := &PostSidecarAttributes{
		map[string]*PostSidecarValue{
			key: postValue,
		},
	}

	postSidecarMetaData := &PostSidecarMetaData{
		meta.GetMetaID(),
		"metadata",
	}
	postRelationships := &PostSidecarRelationships{
		&PostSidecarRelationshipData{
			*postSidecarMetaData,
		},
	}

	postSidecarData := &PostSidecarData{
		"managed-attribute-map",
		*postAttributes,
		*postRelationships,
	}

	return &PostSidecar{
		*postSidecarData,
	}
}

// Helper functionis to build full struct for post data
func getDerivativePostData(upload *db.Upload, meta *db.Meta) *PostSidecarDerivative {
	// uuid of the managed attribute with value
	postAttributes := &PostSidecarDerivativeAttributes{
		"LARGE_IMAGE",
		"Image",
		upload.GetBucket(),
		upload.GetFileIdentifier(),
	}

	postSidecarMetaData := &PostSidecarMetaData{
		meta.GetMetaID(),
		"metadata",
	}
	postRelationships := &PostSidecarDerivativeRelationships{
		&PostSidecarRelationshipData{
			*postSidecarMetaData,
		},
	}

	postSidecarData := &PostSidecarDerivativeData{
		"derivative",
		*postAttributes,
		*postRelationships,
	}

	return &PostSidecarDerivative{
		*postSidecarData,
	}
}
