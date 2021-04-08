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
	DcType         string `json:"dcType"`
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

// post Sidecar data as Managed Metadata
func postSideCarManagedMeta(sidecar *db.Sidecar) error {
	// custom json for all time formats
	var json = jsontime.ConfigWithCustomTimeFormat
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()
	logger.Debug(" sidecar : ", logger.PrettyGoStruct(sidecar))
	// TODO Unknown
	// TODO Do we need to upload managed meta against Original or Derivative?
	file := sidecar.GetOriginalFile()
	meta, err := db.GetMetaByFile(&file)
	if err != nil {
		logger.Fatal("postSideCarManagedMeta on db.GetMetaByFile err : ", err)
	}

	// define url for managed meta
	url := conf.HttpClient.GetBaseApiUrl() + conf.HttpClient.GetManagedMetaUri()
	logger.Debug("post managed meta url : ", url)

	// Read the content of yml sidecar file
	scf, err := walker.GetSidecarYmlFile(sidecar)
	if err != nil {
		return err
	}

	logger.Debug("postSideCarManagedMeta : file : ", logger.PrettyGoStruct(file))
	logger.Debug("postSideCarManagedMeta : meta : ", logger.PrettyGoStruct(meta))
	logger.Debug("postSideCarManagedMeta : sidecar : ", logger.PrettyGoStruct(scf))
	// loop managed attributes as key value
	// POST will be done for eaxh key value pair
	for key, value := range scf.ManagedAttributes {
		// populate post data struct
		postData := getMetaPostData(key, value, meta)
		payload, err := json.Marshal(*postData)
		logger.Debug(" post managed meta payload : ", string(payload))
		if err != nil {
			logger.Error(" json.Marshal fail details :", err)
			return err
		}

		req, err := c.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			logger.Error(" New request Errors :", err)
			return err
		}
		req.Header.Set("Content-Type", "application/vnd.api+json")
		// check if we need Authorization
		if conf.Keycloak.IsEnabled() {
			var bearer = "Bearer " + GetAccessToken()
			req.Header.Set("Authorization", bearer)
		}
		// custom header for https://www.crnk.io/releases/stable/documentation/
		req.Header.Add("crnk-compact", "true")

		resp, err := httpClient.Do(req)
		if err != nil {
			logger.Error(" post managed meta fail details :", err)
			return err
		}
		// TODO Maybe we need a common package for HTTP response status codes
		// Check on response status 401
		if resp.StatusCode == http.StatusUnauthorized {
			logger.Fatal("Error : You are Unauthorized for Managed MetaData, Please check your config file")
		}
		// Check on response status 403
		if resp.StatusCode == http.StatusForbidden {
			logger.Fatal("Error : You are Forbidden from Managed MetaData, Please check your config file")
		}
		// Check on response status 201
		if resp.StatusCode == http.StatusCreated {
			// close the body when done
			defer resp.Body.Close()
			// read the body
			b, err := io.ReadAll(resp.Body)
			logger.Debug(" post managed meta response body : ", string(b))
			if err != nil {
				logger.Error(" error on read body : ", err)
				return err
			}
			logger.Debug(" Managed meta record has been posted : ", logger.PrettyGoStruct(postData))

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

	}
	return nil
}

func postSideCarDerivative(sidecar *db.Sidecar) error {
	// custom json for all time formats
	var json = jsontime.ConfigWithCustomTimeFormat
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()
	// init logger
	logger.Debug(" sidecar : ", logger.PrettyGoStruct(sidecar))

	// TODO Unknown
	// TODO Do we need to upload managed meta against Original or Derivative?
	originalFile := sidecar.GetOriginalFile()
	meta, err := db.GetMetaByFile(&originalFile)
	if err != nil {
		logger.Fatal("postSideCarDerivative on db.GetMetaByFile err : ", err)
	}
	logger.Debug(" postSideCarManagedMeta : meta : ", logger.PrettyGoStruct(meta))

	derivativeFile := sidecar.GetDerivativeFile()
	upload, err := db.GetUploadByFile(&derivativeFile)
	if err != nil {
		logger.Fatal("postSideCarDerivative on db.GetUploadByFile err : ", err)
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
	logger.Debug("postSideCarDerivative : originalFile : ", logger.PrettyGoStruct(originalFile))
	logger.Debug("postSideCarDerivative : derivativeFile : ", logger.PrettyGoStruct(derivativeFile))
	logger.Debug("postSideCarDerivative : meta : ", logger.PrettyGoStruct(meta))
	logger.Debug("postSideCarDerivative : upload : ", logger.PrettyGoStruct(upload))
	logger.Debug("postSideCarDerivative : sidecar : ", logger.PrettyGoStruct(scf))

	// populate post data struct
	postData := getDerivativePostData(upload, meta)
	logger.Debug("postSideCarDerivative : postData : ", logger.PrettyGoStruct(postData))
	payload, err := json.Marshal(*postData)
	logger.Debug(" post derivative payload : ", string(payload))
	if err != nil {
		logger.Error(" json.Marshal fail details :", err)
		return err
	}

	req, err := c.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		logger.Error(" New request Errors :", err)
		return err
	}
	req.Header.Set("Content-Type", "application/vnd.api+json")
	// check if we need Authorization
	if conf.Keycloak.IsEnabled() {
		var bearer = "Bearer " + GetAccessToken()
		req.Header.Set("Authorization", bearer)
	}
	// custom header for https://www.crnk.io/releases/stable/documentation/
	req.Header.Add("crnk-compact", "true")

	resp, err := httpClient.Do(req)
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

func getDerivativePostData(upload *db.Upload, meta *db.Meta) *PostSidecarDerivative {
	// uuid of the managed attribute with value
	postAttributes := &PostSidecarDerivativeAttributes{
		"Image",
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
