package cloudinary

import (
	"context"

	"github.com/satryanararya/go-chefbot/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService interface {
	UploadImage(ctx context.Context, input interface{}) (string, error)
}

type cloudinaryService struct {
	cloudinary *cloudinary.Cloudinary
}

func NewCloudinaryService(cloudinary *cloudinary.Cloudinary) CloudinaryService {
	return &cloudinaryService{
		cloudinary: cloudinary,
	}
}

func (c *cloudinaryService) UploadImage(ctx context.Context, input interface{}) (string, error) {
	uploadParams := uploader.UploadParams{
		Folder:     config.EnvCloudUploadFolder(),
	}

	result, err := c.cloudinary.Upload.Upload(ctx, input, uploadParams)
	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}
