package db

import (
	"context"
	"final-project-backend/config"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)


func ImageUploadHelper(username string, input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conf := config.Config.CloudinaryConfig
	cld, err := cloudinary.NewFromParams(conf.CloudName, conf.APIKey, conf.APISecret)
	if err != nil {
		return "", err
	}

	uploadResult, err := cld.Upload.Upload(
		ctx, 
		input, 
		uploader.UploadParams{
			PublicID:       username,
			UniqueFilename: api.Bool(false),
			Overwrite:      api.Bool(true),
			Folder: 		conf.Folder,
			AllowedFormats: api.CldAPIArray{"jpg","jpeg","png"},
		},
	)

	if err != nil {
		return "", err
	}
	return uploadResult.SecureURL, nil
}

func GetTransformedImageHelper(username string) (string, error){
	conf := config.Config.CloudinaryConfig
	cld, err := cloudinary.NewFromParams(conf.CloudName, conf.APIKey, conf.APISecret)
	if err != nil {
		return "", err
	}

    qs_img, err := cld.Image(username)
    if err != nil {
		return "", err
    }

    qs_img.Transformation = "w_400,h_400,c_fill,r_max"

    new_url, err := qs_img.String()
    if err != nil {
        return "", err
    } 
    
	return new_url, nil
}