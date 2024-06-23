package uploadbusiness

import (
	"bytes"
	"context"
	"fmt"
	"food-delivery/common"
	uploadprovider "food-delivery/component/provider"
	"food-delivery/module/upload/uploadmodel"
	"image"
	"io"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(ctx context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *uploadBiz {
	return &uploadBiz{
		provider: provider,
		imgStore: imgStore,
	}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimession(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil

}

func getImageDimession(data io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(data)
	if err != nil {
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
