package uploadmodel

import (
	"errors"
	"food-delivery/common"
)

const EntityName = "Upload"

type Upload struct {
	common.SQLModel
	common.Image
}

func (Upload) TableName() string {
	return "uploads"
}

var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New("File too large"),
		"File too large",
		"ErrFileTooLarge",
	)
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err, "file is not image", "ErrFileIsNotImage",
	)
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err, "cannot save file", "ErrCannotSaveFile",
	)
}
