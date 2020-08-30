package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/myrachanto/allmicro/gormmicro/customermicroservice/httperors"
	"github.com/myrachanto/allmicro/gormmicro/customermicroservice/model"
)

var (
	Sqlrepository sqlrepository = sqlrepository{}
)

///curtesy to gorm
type sqlrepository struct{}
func init(){
	Getconnected()
}

func Getconnected() (GormDB *gorm.DB, err *httperors.HttpError) {
	dbURI := "root@/micro?charset=utf8&parseTime=True&loc=Local"
	GormDB, err1 := gorm.Open("mysql", dbURI)
	if err1 != nil {
		return nil, httperors.NewNotFoundError("No Mysql db connection")
	}

	GormDB.AutoMigrate(&model.Customer{})
	GormDB.AutoMigrate(&model.Invoice{})
	GormDB.AutoMigrate(&model.InvoiceItem{})
	return GormDB, nil
}
func DbClose(GormDB *gorm.DB) {
	defer GormDB.Close()
}
func (repository sqlrepository) Create(customer *model.Customer) (*model.Customer, *httperors.HttpError) {
	if err := customer.Validate(); err != nil {
		return nil, err
	}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Create(&customer)
	fmt.Println("------------gorm callled--------------")
	DbClose(GormDB)
	return customer, nil
}

func (repository sqlrepository) GetOne(id int) (*model.Customer, *httperors.HttpError) {
	
	customer := model.Customer{}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	fmt.Println(id)
	GormDB.Model(&customer).Where("id = ?", id).First(&customer)
	DbClose(GormDB)
	if customer.ID != 0 {
		return &customer, nil
	}
	return nil, httperors.NewNotFoundError("No customer with that id")
}

func (repository sqlrepository) GetAll(customers []model.Customer) ([]model.Customer, *httperors.HttpError) {
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	Customer := model.Customer{}
	GormDB.Model(&Customer).Find(&customers).Association("Invoices")
	
	DbClose(GormDB)
	if len(customers) == 0 {
		return nil, httperors.NewNotFoundError("No results found!")
	}
	return customers, nil
}

func (repository sqlrepository) Update(id int, customer *model.Customer) (*model.Customer, *httperors.HttpError) {
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	Customer := model.Customer{}
	ucustomer := model.Customer{}
	GormDB.Model(&customer).Where("id = ?", id).First(&ucustomer)
	if ucustomer.ID == 0 {
		return nil, httperors.NewNotFoundError("No customer with that id")
	}
	if customer.Name  == "" {
		customer.Name = ucustomer.Name
	}
	if customer.Company  == "" {
		customer.Company = ucustomer.Company
	}
	if customer.Phone  == "" {
		customer.Phone = ucustomer.Phone
	}
	if customer.Email  == "" {
		customer.Email = ucustomer.Email
	} 
	if customer.Address  == "" {
		customer.Address = ucustomer.Address
	}
	GormDB.Model(&Customer).Where("id = ?", id).First(&Customer).Update(&customer)
	
	DbClose(GormDB)

	return customer, nil
}
func (repository sqlrepository) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	
	customer := model.Customer{}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&customer).Where("id = ?", id).First(&customer)
	/////////////////////////////////////////////////////////////////////////////////////
	////////////////todo check if customer has something//////////////////////////////////
	if customer.ID == 0  {
		return nil, httperors.NewNotFoundError("No results found")
	}
	GormDB.Delete(customer)
	DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}

