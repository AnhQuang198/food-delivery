package uploadprovider

import (
	"context"
	"food-delivery/commons"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*commons.Image, error)
}
