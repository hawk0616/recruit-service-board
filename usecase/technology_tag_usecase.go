package usecase

import (
	"recruit-info-service/model"
	"recruit-info-service/repository"
)

type ITechnologyTagUsecase interface {
	GetAllTechnologyTags(userId uint) ([]model.TechnologyTagResponse, error)
	GetTechnologyTagById(userId uint, technologyTagId uint) (model.TechnologyTagResponse, error)
	CreateTechnologyTag(technologyTag model.TechnologyTag) (model.TechnologyTagResponse, error)
	UpdateTechnologyTag(technologyTag model.TechnologyTag, userId uint, technologyTagId uint) (model.TechnologyTagResponse, error)
	DeleteTechnologyTag(userId uint, technologyTagId uint) error
}

type TechnologyTagUsecase struct {
	ttr repository.ITechnologyTagRepository
}

func NewTechnologyTagUsecase(ttr repository.ITechnologyTagRepository) ITechnologyTagUsecase {
	return &TechnologyTagUsecase{ttr}
}

func (ttu *TechnologyTagUsecase) GetAllTechnologyTags(userId uint) ([]model.TechnologyTagResponse, error) {
	technologyTags := []model.TechnologyTag{}
	if err := ttu.ttr.GetAllTechnologyTags(&technologyTags, userId); err != nil {
		return nil, err
	}
	resTechnologyTags := []model.TechnologyTagResponse{}
	for _, v := range resTechnologyTags {
		t := model.TechnologyTagResponse{
			ID: v.ID,
			Name: v.Name,
		}
		resTechnologyTags = append(resTechnologyTags, t)
	}
	return resTechnologyTags, nil
}

func (ttu *TechnologyTagUsecase) GetTechnologyTagById(userId uint, technologyTagId uint) (model.TechnologyTagResponse, error) {
	technologyTag := model.TechnologyTag{}
	if err := ttu.ttr.GetTechnologyTagById(&technologyTag, userId, technologyTagId); err != nil {
		return model.TechnologyTagResponse{}, err
	}
	resTechnologyTag := model.TechnologyTagResponse{
		ID: technologyTag.ID,
		Name: technologyTag.Name,
	}
	return resTechnologyTag, nil
}

func (ttu *TechnologyTagUsecase) CreateTechnologyTag(technologyTag model.TechnologyTag) (model.TechnologyTagResponse, error) {
	if err := ttu.ttr.CreateTechnologyTag(&technologyTag); err != nil {
		return model.TechnologyTagResponse{}, err
	}
	resTechnologyTag := model.TechnologyTagResponse{
		ID: technologyTag.ID,
		Name: technologyTag.Name,
	}
	return resTechnologyTag, nil
}

func (ttu *TechnologyTagUsecase) UpdateTechnologyTag(technologyTag model.TechnologyTag, userId uint, technologyTagId uint) (model.TechnologyTagResponse, error) {
	if err := ttu.ttr.UpdateTechnologyTag(&technologyTag, userId, technologyTagId); err != nil {
		return model.TechnologyTagResponse{}, err
	}
	resTechnologyTag := model.TechnologyTagResponse{
		ID: technologyTag.ID,
		Name: technologyTag.Name,
	}
	return resTechnologyTag, nil
}

func (ttu *TechnologyTagUsecase) DeleteTechnologyTag(userId uint, technologyTagId uint) error {
	if err := ttu.ttr.DeleteTechnologyTag(userId, technologyTagId); err != nil {
		return err
	}
	return nil
}