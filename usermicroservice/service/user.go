package service

import (
	"fmt"
	"github.com/myrachanto/allmicro/gormmicro/usermicroservice/httperors"
	"github.com/myrachanto/allmicro/gormmicro/usermicroservice/model"
	r "github.com/myrachanto/allmicro/gormmicro/usermicroservice/repository"
)

var (
	UserService userService = userService{}
	repo = r.ChooseRepo()

) 
type redirectUser interface{
	Create(customer *model.User) (*model.User, *httperors.HttpError)
	Login(auser *model.LoginUser) (*model.Auth, *httperors.HttpError)
	Logout(token string) (*httperors.HttpError)
	GetOne(id int) (*model.User, *httperors.HttpError)
	GetAll(users []model.User) ([]model.User, *httperors.HttpError)
	Update(id int, user *model.User) (*model.User, *httperors.HttpError)
	Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError)
}


type userService struct {
	respository r.Redirectrepository
}
func NewRedirectService(respository r.Redirectrepository) redirectUser{
	return &userService{
		respository,
	}
}

func (service userService) Create(user *model.User) (*model.User, *httperors.HttpError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}	
	user, err1 := repo.Create(user)
	if err1 != nil {
		return nil, err1
	}
	 return user, nil

}
func (service userService) Login(auser *model.LoginUser) (*model.Auth, *httperors.HttpError) {
	
	fmt.Println(auser)
	user, err1 := repo.Login(auser)
	if err1 != nil {
		return nil, err1
	}
	return user, nil
}
func (service userService) Logout(token string) (*httperors.HttpError) {
	err1 := repo.Logout(token)
	if err1 != nil {
		return err1
	}
	return nil
}
func (service userService) GetOne(id int) (*model.User, *httperors.HttpError) {
	user, err1 := repo.GetOne(id)
	if err1 != nil {
		return nil, err1
	}
	return user, nil
}

func (service userService) GetAll(users []model.User) ([]model.User, *httperors.HttpError) {
	users, err := repo.GetAll(users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service userService) Update(id int, user *model.User) (*model.User, *httperors.HttpError) {
	
	fmt.Println("update1-controller")
	fmt.Println(id)
	user, err1 := repo.Update(id, user)
	if err1 != nil {
		return nil, err1
	}
	
	return user, nil
}
func (service userService) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	
		success, failure := repo.Delete(id)
		return success, failure
}
