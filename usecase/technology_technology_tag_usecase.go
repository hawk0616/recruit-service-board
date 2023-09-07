package usecase

import (
	"recruit-info-service/model"
	"recruit-info-service/repository"
)

type ITechnologyTechnologyTagUsecase interface {
	CreateTechnologyTechnologyTag(technologyTechnologyTag model.TechnologyTechnologyTag) (model.TechnologyTechnologyTagResponse, error)
	DeleteTechnologyTechnologyTag(technologyId uint, technologyTagId uint) error
}

type TechnologyTechnologyTagUsecase struct {
	tttr repository.ITechnologyTechnologyTagRepository
}

func NewTechnologyTechnologyTagUsecase(tttr repository.ITechnologyTechnologyTagRepository) ITechnologyTechnologyTagUsecase {
	return &TechnologyTechnologyTagUsecase{tttr}
}

func (tttu *TechnologyTechnologyTagUsecase) CreateTechnologyTechnologyTag(technologyTechnologyTag model.TechnologyTechnologyTag) (model.TechnologyTechnologyTagResponse, error) {
	if err := tttu.tttr.CreateTechnologyTechnologyTag(&technologyTechnologyTag); err != nil {
		return model.TechnologyTechnologyTagResponse{}, err
	}
	resTechnologyTechnologyTag := model.TechnologyTechnologyTagResponse{
		ID: technologyTechnologyTag.ID,
		TechnologyID: technologyTechnologyTag.TechnologyID,
		TechnologyTagID: technologyTechnologyTag.TechnologyTagID,
	}
	return resTechnologyTechnologyTag, nil
}

func (tttu *TechnologyTechnologyTagUsecase) DeleteTechnologyTechnologyTag(technologyId uint, technologyTagId uint) error {
	if err := tttu.tttr.DeleteTechnologyTechnologyTag(technologyId, technologyTagId); err != nil {
		return err
	}
	return nil
}