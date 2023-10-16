package validator

import (
	"recruit-info-service/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITechnologyValidator interface {
	TechnologyValidate(Technology model.Technology) error
}

type TechnologyValidator struct{}

func NewTechnologyValidator() ITechnologyValidator {
	return &TechnologyValidator{}
}

func (cv *TechnologyValidator) TechnologyValidate(technology model.Technology) error {
	return validation.ValidateStruct(&technology,
		validation.Field(
			&technology.Name,
			validation.Required.Error("技術名は必須です"),
			validation.RuneLength(1, 140).Error("企業名は最大140文字までです"),
		),
		validation.Field(
			&technology.Description,
			validation.Required.Error("技術概要は必須です"),
			validation.RuneLength(1, 140).Error("技術概要は最大140文字までです"),
		),
	)
}