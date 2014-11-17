package services

/*
	Attaching and retrieving image files
*/

import (
	"github.com/nfnt/resize"
	"image"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type Imageable interface {
	GetTypeName() string
	GetId() int64
}

type Image struct {
	FileName string
	Src      string
	Handle   string
	Name     string
	Desc     string
	Thumbs   struct {
		TinySrc   string
		SmallSrc  string
		MediumSrc string
		LargeSrc  string
	}
}

func generateThumbnail(origFilePath string, imageFolder string, width uint) (string, error) {
	thumbFileName := strconv.FormatUint(uint64(width), 10) + path.Ext(origFilePath)

	resizedFilePath := path.Join(imageFolder, thumbFileName)

	if _, err := os.Stat(resizedFilePath); err == nil {
		return thumbFileName, nil
	}

	file, err := os.Open(origFilePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		return "", err
	}

	m := resize.Resize(width, 0, img, resize.Lanczos3)

	if err = os.MkdirAll(imageFolder, 0775); err != nil {
		return "", err
	}

	out, err := os.Create(resizedFilePath)

	if err != nil {
		return "", err
	}

	defer out.Close()

	return thumbFileName, png.Encode(out, m)
}

func getImageDescription(imageFolder string) (string, error) {
	descriptionFilePath := path.Join(imageFolder, "desc.txt")

	data, err := ioutil.ReadFile(descriptionFilePath)

	if err != nil {
		return "", nil // Assume no description
	}

	return string(data), nil
}

func setImageDescription(imageFolder string, desc string) error {
	descriptionFilePath := path.Join(imageFolder, "desc.txt")

	if err := os.MkdirAll(imageFolder, 0775); err != nil {
		return err
	}

	if err := ioutil.WriteFile(descriptionFilePath, []byte(desc), 0666); err != nil {
		return err
	}

	return nil
}

func GetImages(imageable Imageable) ([]*Image, error) {
	var (
		typeName        = strings.ToLower(imageable.GetTypeName())
		idStr           = strconv.FormatInt(imageable.GetId(), 10)
		imageableFolder = AssetFilePath(path.Join("uploads", typeName, idStr))
	)

	files, _ := ioutil.ReadDir(imageableFolder.Abs())

	var images []*Image

	for _, f := range files {
		if f.Name()[0] == '.' || f.IsDir() {
			continue
		}

		fileName := f.Name()
		fileNameNoExt := strings.Replace(fileName, path.Ext(fileName), "", 1)

		imageFilePathRel := imageableFolder.Append(fileName)
		imageFolderRel := imageableFolder.Append(fileNameNoExt)

		image := &Image{
			FileName: fileName,
			Src:      imageFilePathRel.WebPath(),
			Handle:   fileNameNoExt,
			Name:     fileNameNoExt} // TODO: Humanise

		desc, err := getImageDescription(imageFolderRel.Abs())

		if err != nil {
			return nil, err
		}

		image.Desc = desc

		thumbFileName, err := generateThumbnail(imageFilePathRel.Abs(), imageFolderRel.Abs(), 80)
		image.Thumbs.TinySrc = imageFolderRel.Append(thumbFileName).WebPath()

		thumbFileName, err = generateThumbnail(imageFilePathRel.Abs(), imageFolderRel.Abs(), 160)
		image.Thumbs.SmallSrc = imageFolderRel.Append(thumbFileName).WebPath()

		thumbFileName, err = generateThumbnail(imageFilePathRel.Abs(), imageFolderRel.Abs(), 320)
		image.Thumbs.MediumSrc = imageFolderRel.Append(thumbFileName).WebPath()

		thumbFileName, err = generateThumbnail(imageFilePathRel.Abs(), imageFolderRel.Abs(), 640)
		image.Thumbs.LargeSrc = imageFolderRel.Append(thumbFileName).WebPath()

		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}

func SetImageDescription(imageable Imageable, handle string, desc string) error {
	var (
		typeName        = strings.ToLower(imageable.GetTypeName())
		idStr           = strconv.FormatInt(imageable.GetId(), 10)
		imageableFolder = AssetFilePath(path.Join("uploads", typeName, idStr))
		fileNameNoExt   = handle
		imageFolder     = imageableFolder.Append(fileNameNoExt)
	)

	return setImageDescription(imageFolder.Abs(), desc)
}

func AssignImage(imageable Imageable, tempPath string, handle string, desc string) error {
	var (
		typeName        = strings.ToLower(imageable.GetTypeName())
		idStr           = strconv.FormatInt(imageable.GetId(), 10)
		tempFilePath    = AssetFilePath(path.Join("uploads", "temp", tempPath)) // Temporary file path
		imageableFolder = AssetFilePath(path.Join("uploads", typeName, idStr))
		fileNameNoExt   = handle
		imageFolder     = imageableFolder.Append(fileNameNoExt)
		destFilePath    = imageableFolder.Append(fileNameNoExt + ".png")
	)

	var (
		err      error
		tempFile *os.File
		img      image.Image
		destFile *os.File
	)

	if err = setImageDescription(imageFolder.Abs(), desc); err != nil {
		return err
	}

	if tempFile, err = os.Open(tempFilePath.Abs()); err != nil {
		return err
	}

	defer tempFile.Close()

	if img, _, err = image.Decode(tempFile); err != nil {
		return err
	}

	if err = os.MkdirAll(imageableFolder.Abs(), 0775); err != nil {
		return err
	}

	if destFile, err = os.OpenFile(destFilePath.Abs(), os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		return err
	}

	defer destFile.Close()

	if err = png.Encode(destFile, img); err != nil {
		return err
	}

	err = os.Remove(tempFilePath.Abs())

	return err
}
