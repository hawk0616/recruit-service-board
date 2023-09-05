package validator

import (
	"recruit-info-service/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ICompanyValidator interface {
	CompanyValidate(company model.Company) error
}

type companyValidator struct{}

// companyValidator型がICompanyValidatorインターフェースを満たす
func NewCompanyValidator() ICompanyValidator {
	return &companyValidator{}
}

func (cv *companyValidator) CompanyValidate(company model.Company) error {
	return validation.ValidateStruct(&company,
		validation.Field(
			&company.Name,
			validation.Required.Error("企業名は必須です"),
			validation.RuneLength(1, 20).Error("企業名は最大20文字までです"),
		),
		validation.Field(
			&company.Description,
			validation.Required.Error("企業説明は必須です"),
			validation.RuneLength(1, 1000).Error("企業説明は最大1000文字までです"),
		),
	)
}