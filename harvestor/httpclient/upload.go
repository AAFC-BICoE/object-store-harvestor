package httpclient

import (
	"bytes"
	"fmt"
	"github.com/h2non/filetype"
	c "github.com/hashicorp/go-retryablehttp"
	"github.com/liamylian/jsontime"
	"harvestor/config"
	"harvestor/db"
	l "harvestor/logger"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"
)

// helper var
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

// helper function
func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// allow custom content types to be given for form files
func CreateFormFileWithContentType(w *multipart.Writer, fieldname string, image *db.File) (io.Writer, error) {
	// init logger
	var logger = l.NewLogger()
	// default value
	contentType := "application/octet-stream"
	// reading the file
	buf, err := ioutil.ReadFile(image.GetPath())
	// checkking for errors
	if err != nil {
		logger.Error("CAN NOT read "+image.GetPath()+" ||| details: ", err)
	}
	// getting the kind of the current file
	kind, _ := filetype.Match(buf)
	// if can't detect the kind
	// using default value : "application/octet-stream"
	if kind != filetype.Unknown {
		logger.Debug("kind for "+image.GetPath()+" has been detected as : ", kind.MIME.Value)
		contentType = kind.MIME.Value
	}
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(image.GetName())))
	h.Set("Content-Type", escapeQuotes(contentType))
	return w.CreatePart(h)
}

func uplaodImage(image *db.File) (db.Upload, error) {
	// custom json for all time formats
	var json = jsontime.ConfigWithCustomTimeFormat
	// init logger
	var logger = l.NewLogger()
	// init conf
	conf := config.GetConf()
	// init upload
	var uplaod db.Upload

	// define full resource URL
	url := conf.HttpClient.GetBaseApiUrl() +
		conf.HttpClient.GetUploadUri() +
		"/" +
		conf.HttpClient.GetUploadGroup()
	if image.GetUploadType() == "derivative" {
		url = url + "/derivative"
	}
	logger.Debug("post upload url : ", url)

	// allocate a buffer
	body := &bytes.Buffer{}
	// New multipart writer.
	writer := multipart.NewWriter(body)
	// create multi part for upload
	fw, err := CreateFormFileWithContentType(writer, "file", image)
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
	logger.Debug("request struct has been created for ", image.GetPath())

	req.Header.Set("Content-Type", writer.FormDataContentType())
	// check if we need Authorization
	if conf.Keycloak.IsEnabled() {
		var bearer = "Bearer " + GetAccessToken()
		req.Header.Set("Authorization", bearer)
	}
	// custom header for https://www.crnk.io/releases/stable/documentation/
	req.Header.Add("crnk-compact", "true")

	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Error(image.GetPath()+" POST Errors :", err)
		return uplaod, err
	}
	logger.Debug("POST request has been made to : ", url)
	logger.Debug("Response Status Code : ", resp.StatusCode)
	// TODO Maybe we need a common package for HTTP response status codes
	// Check on response status 401
	if resp.StatusCode == http.StatusUnauthorized {
		logger.Fatal("Error : You are Unauthorized for Uploading, Please check your config file")
	}
	// Check on response status 403
	if resp.StatusCode == http.StatusForbidden {
		logger.Fatal("Error : You are Forbidden from Uploading, Please check your config file")
	}
	// Check on response status 200
	if resp.StatusCode == http.StatusOK {
		// close the body when done
		defer resp.Body.Close()
		// read the body
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error(" error on read body : ", err)
			return uplaod, err
		}

		logger.Debug("about to unmarshal response body : ", string(b))
		// unmarshal body to upload struct
		err = json.Unmarshal(b, &uplaod)
		logger.Debug("DB upload record before create : ", logger.PrettyGoStruct(uplaod))
		if err != nil {
			logger.Error(" error on unmarshal body : ", err)
			return uplaod, err
		}
		// assign a file id
		uplaod.FileID = image.ID
		// creating new record
		err = db.CreateUpload(&uplaod)
		logger.Debug("About to create upload record from response body in DB ...")
		if err != nil {
			logger.Error(" error on upload Create : ", err)
			return uplaod, err
		}
		logger.Debug("DB upload record has been created : ", logger.PrettyGoStruct(uplaod))
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
	return uplaod, err
}
