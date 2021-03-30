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

// | | | relationships | | |
type PostSidecarRelationships struct {
	PostSidecarRelationshipData *PostSidecarRelationshipData `json:"metadata"`
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

// Post sidecar data wrapper
type PostSidecarMeta struct {
	PostSidecarData PostSidecarData `json:"data"`
}

// post Sidecar data as Managed Metadata
func postSideCarManagedMeta(file *db.File, meta *db.Meta) error {
	// custom json for all time formats
	var json = jsontime.ConfigWithCustomTimeFormat
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()
	// define url for managed meta
	url := conf.HttpClient.GetBaseApiUrl() + conf.HttpClient.GetManagedMetaUri()
	logger.Debug("post managed meta url : ", url)

	// populate struct from sidecar yum file
	scf, err := getSidecarYmlFile(file)
	if err != nil {
		logger.Error("unable to get Sidecar File for : ", file.GetPath())
	}

	// loop managed attributes as key value
	for key, value := range scf.ManagedAttributes {
		// populate post data struct
		postData := getPostData(key, value, meta)
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
	return err
}

// Helper function to build full struct for post data
func getPostData(key string, value string, meta *db.Meta) *PostSidecarMeta {
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

	return &PostSidecarMeta{
		*postSidecarData,
	}

}

// get content sidecar yml file
func getSidecarYmlFile(file *db.File) (*walker.SidecarFile, error) {
	path := walker.GetSideCarPathByFilePath(file.GetPath())
	return walker.GetSidecarFile(path)
}
