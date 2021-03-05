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
	"os"
)

func uplaodImage(image *db.File) error {
	var logger = l.NewLogger()
	conf := config.GetConf()

	url := conf.HttpClient.GetBaseApiUrl() + conf.HttpClient.GetUploadUri()

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", image.GetPath())

	file, err := os.Open(image.GetPath())
	if err != nil {
		logger.Error(image.GetPath()+" has errors:", err)
		return err
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		logger.Error(image.GetPath()+" has errors:", err)
		return err
	}
	// Close multipart writer.
	writer.Close()
	file.Close()

	req, err := c.NewRequest("POST", url, bytes.NewReader(body.Bytes()))
	if err != nil {
		logger.Error(image.GetPath()+" NewRequest Errors :", err)
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Error(image.GetPath()+" POST Errors :", err)
		return err
	}
	// read the body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(" error on read body : ", err)
		return err
	}
	// init upload struct
	var u db.Upload
	// unmarshal body to upload struct
	err = json.Unmarshal(b, &u)
	if err != nil {
		logger.Error(" error on Unmarshal body : ", err)
		return err
	}

	// assign fle if from files table
	u.FileID = image.ID
	// TODO testing
	//u.DateTimeDigitized = time.Now()
	// creating new record
	err = u.Create()
	if err != nil {
		logger.Error(" error on u.Create : ", err)
		return err
	}
	return nil
}
