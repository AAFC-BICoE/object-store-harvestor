package httpclient

import (
	"bytes"
	"encoding/json"
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

	resp, err := httpClient.Post(url, "application/vnd.api+json", bytes.NewBuffer(payload))
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
		logger.Debug("about to create meta record : ", meta)
		err = db.CreateMeta(&meta)
		if err != nil {
			logger.Error(" error on Upload Create : ", err)
			return meta, err
		}
	}
	logger.Debug("returning meta record : ", meta)
	return meta, err
}
