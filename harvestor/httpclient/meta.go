package httpclient

import (
	"bytes"
	"encoding/json"
	c "github.com/hashicorp/go-retryablehttp"
	"harvestor/config"
	"harvestor/db"
	l "harvestor/logger"
	"harvestor/walker"
	"io"
	"net/http"
	"time"
)

// structs for http POST
type ManagedAttributes struct {
	ManagedAttributeKey string `json:"managed_attribute_key"`
}
type PostAttributes struct {
	FileIdentifier         string            `json:"fileIdentifier"`
	Bucket                 string            `json:"bucket"`
	DateTimeDigitized      *string           `json:"acDigitizationDate"` // this is a pointer, since we need to support Null value in json
	Orientation            int               `json:"orientation"`
	DcRights               string            `json:"dcRights"`
	ManagedAttributeValues map[string]string `json:"managedAttributes"`
}
type PostData struct {
	Type           string         `json:"type"`
	PostAttributes PostAttributes `json:"attributes"`
}
type PostMeta struct {
	PostData PostData `json:"data"`
}

// Structs for http response
type ResponseData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}
type ResponseMeta struct {
	ResponseData ResponseData `json:"data"`
}

func postMeta(upload *db.Upload) (db.Meta, error) {
	// custom json for all time formats
	//var json = jsontime.ConfigWithCustomTimeFormat
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()
	// init meta struct
	var meta db.Meta

	// define full resource URL
	url := conf.HttpClient.GetBaseApiUrl() + conf.HttpClient.GetMetaUri()
	logger.Debug("post meta url : ", url)
	orientation := walker.GetUploadMediaOrientation(upload)
	logger.Debug("sidecar orientation : ", orientation)

	scf, err := walker.GetSidecariFileByUpload(upload)
	if err != nil {
		logger.Error("GetSidecariFileByUpload :", err)
	}

	// prepare ManagedAttributes for meta
	// init map
	var managedAttributes map[string]string
	// make map
	managedAttributes = make(map[string]string)
	// populate map
	if scf != nil {
		for key, value := range scf.ManagedAttributes {
			managedAttributes[key] = value
		}
	}

	// checking if the GetDateTimeDigitized is actually zero
	postAttributes := &PostAttributes{
		FileIdentifier:         upload.GetFileIdentifier(),
		Bucket:                 upload.GetBucket(),
		DateTimeDigitized:      upload.GetDateTimeDigitized(),
		Orientation:            orientation,
		DcRights:               conf.App.GetDcRights(),
		ManagedAttributeValues: managedAttributes,
	}
	if upload.GetDateTimeDigitized() != nil {
		layout := "2006-01-02T15:04:05"
		dtd := upload.GetDateTimeDigitized()
		loc, err := time.LoadLocation(conf.App.GetObjectTimezone())
		if err != nil {
			logger.Fatal("time.LoadLocation :", err)
		}
		t, err := time.ParseInLocation(layout, *dtd, loc)
		if err != nil {
			logger.Fatal(" time.Parse :", err)
		}
		fixedTime := t.Format(time.RFC3339)
		logger.Debug(" Fixed DateTimeDigitized :", fixedTime)
		postAttributes = &PostAttributes{
			FileIdentifier:         upload.GetFileIdentifier(),
			Bucket:                 upload.GetBucket(),
			DateTimeDigitized:      &fixedTime,
			Orientation:            orientation,
			ManagedAttributeValues: managedAttributes,
		}
	}
	// Building payload
	postData := &PostData{"metadata", *postAttributes}
	postMeta := &PostMeta{*postData}
	payload, err := json.Marshal(*postMeta)
	logger.Debug(" post meta payload : ", string(payload))
	if err != nil {
		logger.Error(" json.Marshal fail details :", err)
		return meta, err
	}

	req, err := c.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		logger.Error(" New request Errors :", err)
		return meta, err
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
		logger.Error(" post meta fail details :", err)
		return meta, err
	}
	// TODO Maybe we need a common package for HTTP response status codes
	// Check on response status 401
	if resp.StatusCode == http.StatusUnauthorized {
		logger.Fatal("Error : You are Unauthorized for MetaData, Please check your config file")
	}
	// Check on response status 403
	if resp.StatusCode == http.StatusForbidden {
		logger.Fatal("Error : You are Forbidden from MetaData, Please check your config file")
	}
	// Check on response status 201
	if resp.StatusCode == http.StatusCreated {
		// close the body when done
		defer resp.Body.Close()
		// read the body
		b, err := io.ReadAll(resp.Body)
		logger.Debug(" post meta response body : ", string(b))
		if err != nil {
			logger.Error(" error on read body : ", err)
			return meta, err
		}

		var rm ResponseMeta
		// unmarshal body to response meta struct
		err = json.Unmarshal(b, &rm)
		logger.Debug(" response meta struct : ", rm)
		if err != nil {
			logger.Error(" error on Unmarshal body : ", err)
			return meta, err
		}
		// assign a meta id
		meta.MetaID = rm.ResponseData.Id
		// assign a file id
		meta.UploadID = upload.GetID()
		meta.FileID = upload.GetFileID()
		// creating new record
		logger.Debug("About to create meta record from response body in DB ...")
		err = db.CreateMeta(&meta)
		if err != nil {
			logger.Error(" error on Upload Create : ", err)
			return meta, err
		}
		logger.Debug("DB meta record has been created : ", logger.PrettyGoStruct(meta))
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
	return meta, err
}
