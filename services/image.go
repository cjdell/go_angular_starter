package services

/*
	Attaching and retrieving image files
*/

import (
	"github.com/cjdell/go_angular_starter/config"
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
	FilePath string
}

func GetImages(imageable Imageable) ([]*Image, error) {
	var (
		typeName           = strings.ToLower(imageable.GetTypeName())
		idStr              = strconv.FormatInt(imageable.GetId(), 10)
		webRoot            = config.App.WebRoot()
		imageableFolder    = path.Join("uploads", typeName, idStr)
		imageableFolderAbs = path.Join(webRoot, imageableFolder)
	)

	files, _ := ioutil.ReadDir(imageableFolderAbs)

	var images []*Image

	for _, f := range files {
		// Exclude hidden files
		if f.Name()[0] != '.' {
			image := &Image{
				FileName: f.Name(),
				FilePath: "/" + path.Join(imageableFolder, f.Name())}

			images = append(images, image)
		}
	}

	return images, nil
}

func AssignImage(imageable Imageable, tempPath string, name string) error {
	var (
		typeName        = strings.ToLower(imageable.GetTypeName())
		idStr           = strconv.FormatInt(imageable.GetId(), 10)
		webRoot         = config.App.WebRoot()
		tempFilePath    = path.Join(webRoot, "uploads", "temp", tempPath) // Temporary file path
		imageableFolder = path.Join(webRoot, "uploads", typeName, idStr)
		destFilePath    = path.Join(imageableFolder, name+path.Ext(tempPath))
	)

	err := os.MkdirAll(imageableFolder, 0775)

	if err != nil {
		return err
	}

	err = os.Rename(tempFilePath, destFilePath)

	if err != nil {
		return err
	}

	return nil
}
