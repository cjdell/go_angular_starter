package services

import (
	"github.com/cjdell/go_angular_starter/config"
	"path"
)

// A path to a file/folder that is always relative to the "assets" folder
type AssetFilePath string

func (filePath AssetFilePath) Append(appendPath string) AssetFilePath {
	return AssetFilePath(path.Join(string(filePath), appendPath))
}

func (filePath AssetFilePath) Abs() string {
	return path.Join(config.App.AssetRoot, string(filePath))
}

func (filePath AssetFilePath) WebPath() string {
	return "/" + path.Join(config.App.AssetRoot, string(filePath))
}

func (filePath AssetFilePath) String() string {
	return string(filePath)
}
