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

func createMockRequest() *http.Request {
	path, _ := os.Getwd()
	file, _ := os.Open(path + "/assets/sample.png")
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path+"/assets/sample.png"))
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

func TestInvalidMimetype(t *testing.T) {
	r := createMockRequest()

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"mimetype:image/jpeg"`
	}

	e = chekku.Validate(Request{
		File: fh,
	})

	if e == nil {
		t.Error("mimetype should be invalid")
		return
	}
}

func TestInvalidMimetypeNotUsedMultipartFileHeader(t *testing.T) {
	r := createMockRequest()

	f, _, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.File `chekku:"mimetype:image/jpeg"`
	}

	e = chekku.Validate(Request{
		File: &f,
	})

	if e == nil {
		t.Error("mimetype should be invalid")
		return
	}
}

func TestValidMimetype(t *testing.T) {
	r := createMockRequest()

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"mimetype:image/png"`
	}

	e = chekku.Validate(Request{
		File: fh,
	})

	if e != nil {
		t.Error("mimetype should be valid")
		return
	}
}
