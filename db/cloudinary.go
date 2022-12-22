package db

import (
	"context"
	"final-project-backend/config"
	"fmt"
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

func TransformImage(cld *cloudinary.Cloudinary, ctx context.Context) {
    // Instantiate an object for the asset with public ID "my_image"
    qs_img, err := cld.Image("quickstart_butterfly")
    if err != nil {
        fmt.Println("error")
    }

    // Add the transformation
    qs_img.Transformation = "r_max/e_sepia"

    // Generate and log the delivery URL
    new_url, err := qs_img.String()
    if err != nil {
        fmt.Println("error")
    } else {
        print("****4. Transform the image****\nTransfrmation URL: ", new_url)
    }
}