package validator

import (
	"recruit-info-service/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ICommentValidator interface {
	CommentValidate(comment model.Comment) error
}

type commentValidator struct{}

func NewCommentValidator() ICommentValidator {
	return &commentValidator{}
}

func (cv *commentValidator) CommentValidate(comment model.Comment) error {
	return validation.ValidateStruct(&comment,
		validation.Field(
			&comment.Content,
			validation.Required.Error("コメントは必須です"),
			validation.RuneLength(1, 140).Error("企業名は最大140文字までです"),
		),
	)
}