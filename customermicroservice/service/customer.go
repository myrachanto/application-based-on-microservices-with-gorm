package service

import (
	"fmt"
	"github.com/myrachanto/allmicro/gormmicro/customermicroservice/httperors"
	"github.com/myrachanto/allmicro/gormmicro/customermicroservice/model"
	r "github.com/myrachanto/allmicro/gormmicro/customermicroservice/repository"
)

var (
	CustomerService customerService = customerService{}
	repo = r.ChooseRepo()

) 
type RedirectCustomer interface{
	Create(customer *model.Customer) (*model.Customer, *httperors.HttpError)
	GetOne(id int) (*model.Customer, *httperors.HttpError)
	GetAll(customers []model.Customer) ([]model.Customer, *httperors.HttpError)
	Update(id int, customer *model.Customer) (*model.Customer, *httperors.HttpError)
	Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError)
}


type customerService struct {
	respository r.Redirectrepository
}
func NewRedirectService(respository r.Redirectrepository) RedirectCustomer{
	return &customerService{
		respository,
	}
}

func (service customerService) Create(customer *model.Customer) (*model.Customer, *httperors.HttpError) {
	if err := customer.Validate(); err != nil {
		return nil, err
	}	
	customer, err1 := repo.Create(customer)
	if err1 != nil {
		return nil, err1
	}
	 return customer, nil

}

func (service customerService) GetOne(id int) (*model.Customer, *httperors.HttpError) {
	customer, err1 := repo.GetOne(id)
	if err1 != nil {
		return nil, err1
	}
	return customer, nil
}

func (service customerService) GetAll(customers []model.Customer) ([]model.Customer, *httperors.HttpError) {
	customers, err := repo.GetAll(customers)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (service customerService) Update(id int, customer *model.Customer) (*model.Customer, *httperors.HttpError) {
	
	fmt.Println("update1-controller")
	fmt.Println(id)
	customer, err1 := repo.Update(id, customer)
	if err1 != nil {
		return nil, err1
	}
	
	return customer, nil
}
func (service customerService) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	
		success, failure := repo.Delete(id)
		return success, failure
}
