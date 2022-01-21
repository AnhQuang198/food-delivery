package model

import (
	"errors"
	"food-delivery/commons"
)

const EntityName = "Upload"

type Upload struct {
	commons.SQLModel `json:",inline"`
	commons.Image    `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

//
//func (u *Upload) Mask(isAdmin bool) {
//	u.GenUID(common.DBTypeUpload, 1)
//}

var (
	ErrFileTooLarge = commons.NewCustomError(
		errors.New("file too large"),
		"file too large",
		"ErrFileTooLarge",
	)
)

func ErrFileIsNotImage(err error) *commons.AppError {
	return commons.NewCustomError(
		err,
		"file is not image",
		"ErrFileIsNotImage",
	)
}

func ErrCannotSaveFile(err error) *commons.AppError {
	return commons.NewCustomError(
		err,
		"cannot save uploaded file",
		"ErrCannotSaveFile",
	)
}
