package save

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/EmotivesProject/common/logger"
)

const randomLength = 10

var fileDir = os.Getenv("FILE_LOCATION")

func File(imageData *image.Image, filename, username string, randomise bool) (string, error) {
	path, err := setUpPathForUser(username)
	if err != nil {
		return "", err
	}

	logger.Infof("Created path %s", path)

	filename = createFinalFilename(filename, randomise)

	return saveFile(filename, path, username, imageData)
}

func setUpPathForUser(username string) (string, error) {
	path := fmt.Sprintf(".%s/%s", fileDir, username)
	err := os.MkdirAll(path, os.ModePerm)

	return path, err
}

func createFinalFilename(filename string, randomise bool) string {
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))

	if randomise {
		generatedFilename, err := generateRandomString(randomLength)
		if err != nil {
			generatedFilename = ""

			logger.Error(err)
		}

		filename = fmt.Sprintf("%s_%s", generatedFilename, filename)
	}

	filename = fmt.Sprintf("%s.png", filename)
	logger.Infof("Saving file with %s", filename)

	return filename
}

func saveFile(filename, path, username string, imageData *image.Image) (string, error) {
	filePath := fmt.Sprintf("%s/%s", path, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	err = png.Encode(file, *imageData)
	urlPath := fmt.Sprintf("%s/%s/%s", fileDir, username, filename)

	return urlPath, err
}
