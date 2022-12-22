package utils

import (
	"final-project-backend/db"
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"mime/multipart"
	"net/http"

	"github.com/go-playground/validator/v10"
)


var (
    validate = validator.New()
)

type mediaUpload interface {
    FileUpload(username string, file entity.UserProfileUploadReqBody) (string, error)
    RemoteUpload(username string, url entity.UserProfileUploadResBody) (string, error)
}

type media struct {}

func NewMediaUpload() mediaUpload {
    return &media{}
}

func (*media) FileUpload(username string, file entity.UserProfileUploadReqBody) (string, error) {
    err := validate.Struct(file)
    if err != nil {
        return "", err
    }

	err = ValidateContentType(file.File)
	if err != nil {
        return "", err
    }

    uploadUrl, err := db.ImageUploadHelper(username, file.File)
    if err != nil {
        return "", err
    }
    return uploadUrl, nil
}

func ValidateContentType(file multipart.File) (error) {
	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
	   return err
	}
 
	contentType := http.DetectContentType(buff)
	if 	contentType != "image/jpg" && 
		contentType != "image/jpeg" && 
		contentType != "image/png" {
		return errorlist.BadRequestError(
			"Invalid input datatype. Only receiving jpg/ png/ jpeg file format.",
			"INVALID_INPUT",
		)
	}

	if _, err := file.Seek(0, 0); err != nil{
		return err
	}
	return nil
 }

func (*media) RemoteUpload(username string, url entity.UserProfileUploadResBody) (string, error) {
    err := validate.Struct(url)
    if err != nil {
        return "", err
    }

    uploadUrl, errUrl := db.ImageUploadHelper(username, url.Url)
    if errUrl != nil {
        return "", err
    }
    return uploadUrl, nil
}