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

func createMockRequest3(imgName string) *http.Request {
	path, _ := os.Getwd()
	file, _ := os.Open(path + "/assets/" + imgName)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path+"/assets/"+imgName))
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

func TestInvalidRatio(t *testing.T) {
	r := createMockRequest3("sample2.png")

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"dimensions:ratio=16/9"`
	}

	eStack := chekku.Validate(Request{
		File: fh,
	})

	if eStack == nil {
		t.Errorf("dimensions ratio should be invalid")
		return
	}
}

func TestValidRatio(t *testing.T) {
	r := createMockRequest3("sample.png")

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"dimensions:ratio=1/1"`
	}

	eStack := chekku.Validate(Request{
		File: fh,
	})

	if eStack != nil {
		t.Errorf("dimenstions ratio should be valid")
		return
	}
}

func TestValidMinWidth(t *testing.T) {
	r := createMockRequest3("sample.png")

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"dimensions:min_width=150"`
	}

	eStack := chekku.Validate(Request{
		File: fh,
	})

	if eStack != nil {
		t.Errorf("dimensions min_width should be valid")
		return
	}
}

func TestInvalidMaxWidth(t *testing.T) {
	r := createMockRequest3("sample.png")

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"dimensions:max_width=150"`
	}

	eStack := chekku.Validate(Request{
		File: fh,
	})

	if eStack == nil {
		t.Errorf("dimensions max_width should be invalid")
		return
	}
}

func TestValidMaxWidth(t *testing.T) {
	r := createMockRequest3("sample.png")

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"dimensions:max_width=300"`
	}

	eStack := chekku.Validate(Request{
		File: fh,
	})

	if eStack != nil {
		t.Errorf("dimensions max_width should be valid")
		return
	}
}

func TestInvalidMinWidth(t *testing.T) {
	r := createMockRequest3("sample.png")

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"dimensions:min_width=300"`
	}

	eStack := chekku.Validate(Request{
		File: fh,
	})

	if eStack == nil {
		t.Errorf("dimensions min_width should be invalid")
		return
	}
}

func TestInvalidMinHeight(t *testing.T) {
	r := createMockRequest3("sample.png")

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"dimensions:min_height=300"`
	}

	eStack := chekku.Validate(Request{
		File: fh,
	})

	if eStack == nil {
		t.Errorf("dimensions min_height should be invalid")
		return
	}
}

func TestValidMinHeight(t *testing.T) {
	r := createMockRequest3("sample.png")

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"dimensions:min_height=150"`
	}

	eStack := chekku.Validate(Request{
		File: fh,
	})

	if eStack != nil {
		t.Errorf("dimensions min_height should be valid")
		return
	}
}

func TestInvalidMaxHeight(t *testing.T) {
	r := createMockRequest3("sample.png")

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"dimensions:max_height=150"`
	}

	eStack := chekku.Validate(Request{
		File: fh,
	})

	if eStack == nil {
		t.Errorf("dimensions max_height should be invalid")
		return
	}
}

func TestValidMaxHeight(t *testing.T) {
	r := createMockRequest3("sample.png")

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"dimensions:max_height=300"`
	}

	eStack := chekku.Validate(Request{
		File: fh,
	})

	if eStack != nil {
		t.Errorf("dimensions min_height should be valid")
		return
	}
}

func TestValidMultipleWidthHeight(t *testing.T) {
	r := createMockRequest3("sample.png")

	_, fh, e := r.FormFile("file")
	if e != nil {
		t.Errorf("%v", e)
		return
	}

	type Request struct {
		File *multipart.FileHeader `chekku:"dimensions:max_width=300,min_width=150,max_height=300,min_height=150,ratio=1/1"`
	}

	eStack := chekku.Validate(Request{
		File: fh,
	})

	if eStack != nil {
		t.Errorf("dimensions max_width, min_width, max_height, min_height and ratio should be valid")
		return
	}
}
