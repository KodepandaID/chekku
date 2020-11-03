package rules

import (
	"fmt"
	"image"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	// image/jpeg, image/png, image/gif to decode image
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"golang.org/x/text/message"
)

// Mimetype validation file. You see the full listing of MIME types at this link
// https://svn.apache.org/repos/asf/httpd/httpd/trunk/docs/conf/mime.types
func (r Rules) Mimetype(fieldName string, v reflect.Value, m string) error {
	validMime := strings.Split(m, ",")
	if v.IsNil() {
		return fmt.Errorf("Invalid file not found")
	}

	switch v.Interface().(type) {
	case *multipart.FileHeader:
		f, _ := v.Interface().(*multipart.FileHeader).Open()
		size := v.Interface().(*multipart.FileHeader).Size
		defer f.Close()

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

// Filesize validation file size. The file size cannot be exceeded the max.
func (r Rules) Filesize(fieldName string, v reflect.Value, m string) error {
	max, _ := strconv.ParseInt(m, 10, 64)
	if v.IsNil() {
		return fmt.Errorf("Invalid file not found")
	}

	switch v.Interface().(type) {
	case *multipart.FileHeader:
		size := v.Interface().(*multipart.FileHeader).Size / 1000
		if size > max {
			p := message.NewPrinter(message.MatchLanguage("id"))
			return fmt.Errorf("The file size cannot be exceeded %vKB", p.Sprintf("%d", max))
		}

		return nil
	default:
		return fmt.Errorf("Invalid type only supports *multipart.FileHeader")
	}
}

// Dimensions validation image dimension with minimal width and height or
// maximal width and height. You can also set the ratio to validate.
func (r Rules) Dimensions(fieldName string, v reflect.Value, m string) error {
	c := strings.Split(m, ",")
	if v.IsNil() {
		return fmt.Errorf("Invalid file not found")
	}

	switch v.Interface().(type) {
	case *multipart.FileHeader:
		f, _ := v.Interface().(*multipart.FileHeader).Open()
		defer f.Close()

		e := dimensionOperation(c, f)
		if e != nil {
			return fmt.Errorf("%v", e)
		}

		return nil
	default:
		return fmt.Errorf("Invalid type only supports *multipart.FileHeader")
	}
}

func dimensionOperation(c []string, f multipart.File) error {
	w, h, e := getImageDimension(f)
	if e != nil {
		return e
	}

	for _, v := range c {
		arg := strings.Split(v, "=")
		num, _ := strconv.Atoi(arg[1])

		switch {
		case arg[0] == "width":
			if w != num {
				return fmt.Errorf("Image width should be %vpx", num)
			}

			return nil
		case arg[0] == "min_width":
			if w < num {
				return fmt.Errorf("Image width should not be less than %vpx", num)
			}

			return nil
		case arg[0] == "max_width":
			if w > num {
				return fmt.Errorf("Image width should not be more than %vpx", num)
			}

			return nil
		case arg[0] == "height":
			if h != num {
				return fmt.Errorf("Image height should be %vpx", num)
			}

			return nil
		case arg[0] == "min_height":
			if h < num {
				return fmt.Errorf("Image height should not be less than %vpx", num)
			}

			return nil
		case arg[0] == "max_height":
			if h > num {
				return fmt.Errorf("Image height should not be more than %vpx", num)
			}

			return nil
		case arg[0] == "ratio":
			e := aspectRatio(w, h, arg[1])
			if e != nil {
				return fmt.Errorf("%v", e)
			}

			return nil
		default:
			return fmt.Errorf("Dimensions cannot process your statements please use min_width, max_width, min_height, max_height or ratio")
		}
	}

	return nil
}

func getImageDimension(f multipart.File) (int, int, error) {
	img, _, e := image.DecodeConfig(f)
	if e != nil {
		return 0, 0, e
	}

	return img.Width, img.Height, nil
}

func aspectRatio(w, h int, ratio string) error {
	var dividend int
	var divisor int

	if h == w && ratio == "1/1" {
		return nil
	}

	if h > w {
		dividend = h
		divisor = w
	} else if w > h {
		dividend = w
		divisor = h
	}

	gcd := -1
	for gcd == -1 {
		r := dividend % divisor
		if r == 0 {
			gcd = divisor
		} else {
			dividend = divisor
			divisor = r
		}
	}

	hr := w / gcd
	vr := h / gcd
	nowRatio := fmt.Sprintf("%v/%v", hr, vr)

	if nowRatio != ratio {
		return fmt.Errorf("Image ratio is %v, the image ratio should be %v", nowRatio, ratio)
	}

	return nil
}
