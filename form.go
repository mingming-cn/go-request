package request

import (
	"bytes"
	"io"
	"mime/multipart"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

// formReader 根据Args中的参数生成 Form 格式的 Reader
func (a *Args) formReader() (io.Reader, error) {
	if len(a.Files) > 0 {
		return a.multipartReader()
	}

	return strings.NewReader(a.FormData.Encode()), nil
}

// multipartReader 根据Args中的参数生成 Multipart Form 格式的 Reader
func (a *Args) multipartReader() (io.Reader, error) {
	bodyBuffer := new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(bodyBuffer)
	defer bodyWriter.Close()

	for key := range a.Files {
		fileWriter, err := bodyWriter.CreateFormFile(a.Files[key].FieldName, a.Files[key].FileName)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(fileWriter, a.Files[key].File)
		if err != nil {
			return nil, err
		}
	}

	for k, arr := range a.FormData {
		for n := range arr {
			bodyWriter.WriteField(k, arr[n])
		}
	}

	a.setContentType(NewContentType(bodyWriter.FormDataContentType()))
	return bodyBuffer, nil
}

// jsonReader 根据Args中的参数生成 json 格式的 Reader
func (a *Args) jsonReader() (io.Reader, error) {
	b, err := jsoniter.Marshal(a.JSON)
	if err != nil {
		return nil, err
	}

	a.setContentType(ApplicationJSON)
	return bytes.NewReader(b), err
}
