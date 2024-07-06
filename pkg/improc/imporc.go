package improc

import (
	"errors"
	"gomodule/pkg/pass"
	"image"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

type ImageConf struct {
	Dir        string
	CustomName string
	MaxWidth   int
}

func (i *ImageConf) UploadImage(r *http.Request) (path, name string, size int64, err error) {

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		return "", "", 0, err
	}
	defer file.Close()

	allowedExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}
	fileExtension := filepath.Ext(fileHeader.Filename)
	if !allowedExtensions[fileExtension] {
		return "", "", 0, errors.New("invalid format")
	}
	os.MkdirAll(i.Dir, 0755)
	imgName := pass.RandUlid() + "_.webp"
	filePath := filepath.Join(i.Dir, i.CustomName+"_"+imgName)

	newCreate, err := os.Create(filePath)
	if err != nil {
		return "", "", 0, err
	}
	defer newCreate.Close()
	_, err = io.Copy(newCreate, file)
	if err != nil {
		return "", "", 0, err
	}
	return filePath, imgName, fileHeader.Size, nil
}

func (i *ImageConf) Webp(src string) error {
	img, err := imgio.Open(src)
	if err != nil {
		return err
	}

	output, err := os.Create(src)
	if err != nil {
		return err
	}
	defer output.Close()
	w, h := i.Rescale(img)
	img = transform.Resize(img, w, h, transform.Linear)
	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 100)
	if err != nil {
		return err
	}

	if err := webp.Encode(output, img, options); err != nil {
		return err
	}
	return nil
}

func (i *ImageConf) Rescale(input image.Image) (newWidth, newHeight int) {
	originalWidth := input.Bounds().Dx()
	originalHeight := input.Bounds().Dy()
	scale := float64(i.MaxWidth) / float64(originalWidth)
	newWidth = i.MaxWidth
	newHeight = int(float64(originalHeight) * scale)
	return newWidth, newHeight
}
