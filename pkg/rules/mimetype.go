package rules

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
)

// Mimetype validation file. You see the full listing of MIME types at this link
// https://svn.apache.org/repos/asf/httpd/httpd/trunk/docs/conf/mime.types
func (r Rules) Mimetype(fieldName string, v reflect.Value, m string) error {
	validMime := strings.Split(m, ",")

	switch v.Interface().(type) {
	case *multipart.FileHeader:
		f, _ := v.Interface().(*multipart.FileHeader).Open()
		size := v.Interface().(*multipart.FileHeader).Size

		fileHeader := make([]byte, size)
		if _, e := f.Read(fileHeader); e != nil {
			return fmt.Errorf("%v", e)
		}

		mime := http.DetectContentType(fileHeader)
		for _, v := range validMime {
			if mime == v {
				return nil
			}
		}

		return fmt.Errorf("Invalid mimetype, only supports mimetypes %v", strings.Join(validMime, ", "))
	default:
		return fmt.Errorf("Invalid type only supports *multipart.FileHeader")
	}
}
