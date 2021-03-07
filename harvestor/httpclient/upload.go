package httpclient

import (
	"bytes"
	"encoding/json"
	c "github.com/hashicorp/go-retryablehttp"
	"harvestor/config"
	"harvestor/db"
	l "harvestor/logger"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func uplaodImage(image *db.File) (db.Upload, error) {
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()
	// init upload
	var uplaod db.Upload

	// define full resource URL
	url := conf.HttpClient.GetBaseApiUrl() + conf.HttpClient.GetUploadUri()
	logger.Debug("post upload url : ", url)

	// allocate a buffer
	body := &bytes.Buffer{}
	// New multipart writer.
	writer := multipart.NewWriter(body)
	// create multi part for upload
	fw, err := writer.CreateFormFile("file", image.GetPath())

	// validating on file open
	logger.Debug("opening file : ", image.GetPath())
	file, err := os.Open(image.GetPath())
	if err != nil {
		logger.Error("CAN NOT open "+image.GetPath()+" ||| details: ", err)
		return uplaod, err
	}
	// validating on file content copy
	_, err = io.Copy(fw, file)
	if err != nil {
		logger.Error("CAN NOT copy "+image.GetPath()+" ||| details: ", err)
		return uplaod, err
	}

	logger.Debug("upload  multi part has been created for :", image.GetPath())
	// Close multipart writer.
	writer.Close()
	logger.Debug("writer closed")
	// Close file
	file.Close()
	logger.Debug("file closed")
	// Close file

	// at this point we are sure that the local file can be uploaded
	logger.Debug("http request body size : ", body.Len())
	req, err := c.NewRequest("POST", url, bytes.NewReader(body.Bytes()))
	if err != nil {
		logger.Error(image.GetPath()+" Request Errors :", err)
		return uplaod, err
	}
	logger.Debug("request has been created for ", image.GetPath())

	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Error(image.GetPath()+" POST Errors :", err)
		return uplaod, err
	}
	logger.Debug("POST request has been made to : ", url)
	logger.Debug("Response Status Code : ", resp.StatusCode)
	// Check on response status 200
	if resp.StatusCode == http.StatusOK {
		// close the body when done
		defer resp.Body.Close()
		// read the body
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Error(" error on read body : ", err)
			return uplaod, err
		}

		logger.Debug("about to unmarshal body : ", string(b))
		// unmarshal body to upload struct
		err = json.Unmarshal(b, &uplaod)
		if err != nil {
			logger.Error(" error on unmarshal body : ", err)
			return uplaod, err
		}
		// assign a file id
		uplaod.FileID = image.ID
		// creating new record
		logger.Debug("about to create upload record : ", uplaod)
		err = db.CreateUpload(&uplaod)
		if err != nil {
			logger.Error(" error on upload Create : ", err)
			return uplaod, err
		}
	}
	logger.Debug("returning upload record : ", uplaod)
	return uplaod, err
}