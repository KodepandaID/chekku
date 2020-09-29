package test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/KodepandaID/chekku"
)

func createMockRequest2() *http.Request {
	path, _ := os.Getwd()
	file, _ := os.Open(path + "/assets/sample2.png")
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path+"/assets/sample2.png"))
	if err != nil {
		return nil
	}
	_, _ = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil
	}

	req, _ := http.NewRequest("POST", "", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req
}

func TestInvalidFilesize(t *testing.T) {
	r := createMockRequest2()

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"filesize:100"`
	}

	e = chekku.Validate(Request{
		File: fh,
	})

	if e == nil {
		t.Error("filesize should be invalid")
		return
	}
}

func TestValidFilesize(t *testing.T) {
	r := createMockRequest2()

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"filesize:512"`
	}

	e = chekku.Validate(Request{
		File: fh,
	})

	if e != nil {
		t.Error("filesize should be valid")
		return
	}
}
