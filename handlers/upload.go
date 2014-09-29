package handlers

import (
	"fmt"
	"github.com/cjdell/go_angular_starter/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	var (
		fileId   = r.Header.Get("X-Upload-File-ID")
		fileName = r.Header.Get("X-Upload-File-Name")
	)

	var (
		saveFileName = fileId + "_" + fileName
		savePath     = path.Join(config.App.WebRoot(), "uploads", "temp")
		saveFilePath = path.Join(savePath, saveFileName)
	)

	err = os.MkdirAll(savePath, 0775)

	fo, err := os.Create(saveFilePath)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = fo.Write(body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprint(w, saveFileName)
}
