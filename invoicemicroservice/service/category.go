package service

import (
	// "fmt"
	"github.com/myrachanto/allmicro/gormmicro/categorymicroservice/httperors"
	"github.com/myrachanto/allmicro/gormmicro/categorymicroservice/model"
	r "github.com/myrachanto/allmicro/gormmicro/categorymicroservice/repository"
)

var (
	CategoryService categoryService = categoryService{}
	repo = r.ChooseRepo()

) 
type redirectCategroy interface{
	Create(category *model.Category) (*model.Category, *httperors.HttpError)
	GetOne(id int) (*model.Category, *httperors.HttpError)
	GetAll(users []model.Category) ([]model.Category, *httperors.HttpError)
	Update(id int, category *model.Category) (*model.Category, *httperors.HttpError)
	Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError)
}


type categoryService struct {
	respository r.Redirectrepository
}
func NewRedirectService(respository r.Redirectrepository) redirectCategroy{
	return &categoryService{
		respository,
	}
}

func (service categoryService) Create(category *model.Category) (*model.Category, *httperors.HttpError) {
	if err := category.Validate(); err != nil {
		return nil, err
	}	
	category, err1 := repo.Create(category)
	if err1 != nil {
		return nil, err1
	}
	 return category, nil

}
func (service categoryService) GetOne(id int) (*model.Category, *httperors.HttpError) {
	category, err1 := repo.GetOne(id)
	if err1 != nil {
		return nil, err1
	}
	return category, nil
}

func (service categoryService) GetAll(categorys []model.Category) ([]model.Category, *httperors.HttpError) {
	categorys, err := repo.GetAll(categorys)
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func (service categoryService) Update(id int, category *model.Category) (*model.Category, *httperors.HttpError) {
	category, err1 := repo.Update(id, category)
	if err1 != nil {
		return nil, err1
	}
	
	return category, nil
}
func (service categoryService) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	
		success, failure := repo.Delete(id)
		return success, failure
}
