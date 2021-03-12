package httpclient

import (
	"bytes"
	c "github.com/hashicorp/go-retryablehttp"
	"github.com/liamylian/jsontime"
	"harvestor/config"
	"harvestor/db"
	l "harvestor/logger"
	"io/ioutil"
	"net/http"
)

// structs for http POST
type PostAttributes struct {
	FileIdentifier string `json:"fileIdentifier"`
	Bucket         string `json:"bucket"`
	//DateTimeDigitized time.Time `json:"acDigitizationDate"`
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
	var json = jsontime.ConfigWithCustomTimeFormat
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()
	// init meta struct
	var meta db.Meta

	// define full resource URL
	url := conf.HttpClient.GetBaseApiUrl() + conf.HttpClient.GetMetaUri()
	logger.Debug("post meta url : ", url)
	postAttributes := &PostAttributes{
		upload.GetFileIdentifier(),
		upload.GetBucket(),
		// TODO may be WIP
		//upload.GetDateTimeDigitized(),
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
	// custom header for https://www.crnk.io/releases/stable/documentation/
	req.Header.Set("crnk-compact", "true")
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Error(" post meta fail details :", err)
		return meta, err
	}

	// Check on response status 201
	if resp.StatusCode == http.StatusCreated {
		// close the body when done
		defer resp.Body.Close()
		// read the body
		b, err := ioutil.ReadAll(resp.Body)
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
		meta.UploadID = upload.ID
		// creating new record
		logger.Debug("About to create meta record from response body in DB ...")
		err = db.CreateMeta(&meta)
		if err != nil {
			logger.Error(" error on Upload Create : ", err)
			return meta, err
		}
		logger.Debug("DB meta record has been created : ", logger.PrettyGoStruct(meta))
	}
	return meta, err
}
