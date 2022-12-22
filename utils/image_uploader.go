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
}

type media struct {}

func NewMediaUpload() mediaUpload {
    return &media{}
}

func (*media) FileUpload(username string, reqBody entity.UserProfileUploadReqBody) (string, error) {
    err := validate.Struct(reqBody)
    if err != nil {
        return "", err
    }

	file, err := reqBody.Img.Open()
	if err != nil {
		return "", err
	 }
	defer file.Close()

	err = ValidateContentType(file)
	if err != nil {
        return "", err
    }

    uploadUrl, err := db.ImageUploadHelper(username, file)
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
