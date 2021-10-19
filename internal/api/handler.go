package api

import (
	"image/jpeg"
	"image/png"
	"media/internal/save"
	"media/messages"
	"media/model"
	"mime/multipart"
	"net/http"

	"github.com/TomBowyerResearchProject/common/logger"
	"github.com/TomBowyerResearchProject/common/response"
	"github.com/TomBowyerResearchProject/common/verification"
)

const (
	imageFormLocation = "image"
	imageFormName     = "name"
	limit             = 2
	MB                = 20
)

func uploadImage(w http.ResponseWriter, r *http.Request) {
	username, file, handle, err := extractDetailsFromRequest(r)
	if err != nil {
		logger.Error(err)
		response.ResultResponseJSON(w, false, http.StatusBadRequest, response.Message{Message: err.Error()})

		return
	}

	fileName, randomise := extractNameAndRandomise(r, handle)

	filePath, err := decodeAndSave(file, handle, fileName, username, randomise)
	if err != nil {
		logger.Error(err)
		response.ResultResponseJSON(w, false, http.StatusInternalServerError, response.Message{Message: err.Error()})

		return
	}

	uploadedFile := model.File{
		Owner:    username,
		Location: filePath,
	}

	logger.Infof("Created image at %s", filePath)

	response.ResultResponseJSON(w, false, http.StatusCreated, uploadedFile)
}

func uploadUserImage(w http.ResponseWriter, r *http.Request) {
	username, file, handle, err := extractDetailsFromRequest(r)
	if err != nil {
		logger.Error(err)
		response.ResultResponseJSON(w, false, http.StatusBadRequest, response.Message{Message: err.Error()})

		return
	}

	filePath, err := decodeAndSave(file, handle, username, "user", false)
	if err != nil {
		logger.Error(err)
		response.ResultResponseJSON(w, false, http.StatusInternalServerError, response.Message{Message: err.Error()})

		return
	}

	uploadedFile := model.File{
		Owner:    username,
		Location: filePath,
	}

	logger.Infof("Created image at %s", filePath)

	response.ResultResponseJSON(w, false, http.StatusCreated, uploadedFile)
}

func extractDetailsFromRequest(r *http.Request) (
	string,
	multipart.File,
	*multipart.FileHeader,
	error,
) {
	username, ok := r.Context().Value(verification.UserID).(string)
	if !ok {
		return "", nil, nil, messages.ErrFailedType
	}

	err := r.ParseMultipartForm(limit << MB) // 2MB limit
	if err != nil {
		return "", nil, nil, err
	}

	file, handle, err := r.FormFile(imageFormLocation)
	if err != nil {
		return "", nil, nil, err
	}

	return username, file, handle, nil
}

func extractNameAndRandomise(r *http.Request, handle *multipart.FileHeader) (string, bool) {
	randomise := true

	fileName := handle.Filename

	name := r.FormValue(imageFormName)
	if name != "" {
		fileName = name
		randomise = false
	}

	return fileName, randomise
}

func decodeAndSave(
	file multipart.File,
	handle *multipart.FileHeader,
	fileName,
	username string,
	randomise bool,
) (string, error) {
	switch handle.Header.Get("Content-Type") {
	case "image/jpeg":
		fallthrough
	case "image/jpg":
		image, err := jpeg.Decode(file)
		if err != nil {
			return "", err
		}

		return save.File(&image, fileName, username, randomise)
	case "image/png":
		image, err := png.Decode(file)
		if err != nil {
			return "", err
		}

		return save.File(&image, fileName, username, randomise)
	default:
		return "", messages.ErrIncorrectImage
	}
}
