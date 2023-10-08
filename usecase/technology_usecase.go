package usecase

import (
	"recruit-info-service/model"
	"recruit-info-service/repository"
)

type ITechnologyUsecase interface {
	GetAllTechnologies() ([]model.TechnologyResponse, error)
	GetTechnologyById(technologyId uint) (model.TechnologyResponse, error)
	CreateTechnology(technology model.Technology) (model.TechnologyResponse, error)
	UpdateTechnology(technology model.Technology, technologyId uint) (model.TechnologyResponse, error)
	DeleteTechnology(technologyId uint) error
}

type TechnologyUsecase struct {
	tr repository.ITechnologyRepository
}

func NewTechnologyUsecase(tr repository.ITechnologyRepository) ITechnologyUsecase {
	return &TechnologyUsecase{tr}
}

func (tu *TechnologyUsecase) GetAllTechnologies() ([]model.TechnologyResponse, error) {
	technologies := []model.Technology{}
	if err := tu.tr.GetAllTechnologies(&technologies); err != nil {
		return nil, err
	}
	resTechnologies := []model.TechnologyResponse{}
	for _, v := range technologies {
		t := model.TechnologyResponse{
			ID: v.ID,
			Name: v.Name,
			Description: v.Description,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTechnologies = append(resTechnologies, t)
	}
	return resTechnologies, nil
}

func (tu *TechnologyUsecase) GetTechnologyById(technologyId uint) (model.TechnologyResponse, error) {
	technology := model.Technology{}
	if err := tu.tr.GetTechnologyById(&technology, technologyId); err != nil {
		return model.TechnologyResponse{}, err
	}
	resTechnology := model.TechnologyResponse{
		ID: technology.ID,
		Name: technology.Name,
		Description: technology.Description,
		CreatedAt: technology.CreatedAt,
		UpdatedAt: technology.UpdatedAt,
	}
	return resTechnology, nil
}

func (tu *TechnologyUsecase) CreateTechnology(technology model.Technology) (model.TechnologyResponse, error) {
	if err := tu.tr.CreateTechnology(&technology); err != nil {
		return model.TechnologyResponse{}, err
	}
	resTechnology := model.TechnologyResponse{
		ID: technology.ID,
		Name: technology.Name,
		Description: technology.Description,
		CreatedAt: technology.CreatedAt,
		UpdatedAt: technology.UpdatedAt,
	}
	return resTechnology, nil
}

func (tu *TechnologyUsecase) UpdateTechnology(technology model.Technology, technologyId uint) (model.TechnologyResponse, error) {
	if err := tu.tr.UpdateTechnology(&technology, technologyId); err != nil {
		return model.TechnologyResponse{}, err
	}
	resTechnology := model.TechnologyResponse{
		ID: technology.ID,
		Name: technology.Name,
		Description: technology.Description,
		CreatedAt: technology.CreatedAt,
		UpdatedAt: technology.UpdatedAt,
	}
	return resTechnology, nil
}

func (tu *TechnologyUsecase) DeleteTechnology(technologyId uint) error {
	if err := tu.tr.DeleteTechnology(technologyId); err != nil {
		return err
	}
	return nil
}