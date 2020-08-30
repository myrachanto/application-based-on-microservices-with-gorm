package repository

import(
    "fmt"
	"github.com/myrachanto/allmicro/gormmicro/customermicroservice/httperors"
	"github.com/myrachanto/allmicro/gormmicro/customermicroservice/model"
)
var (
	MockRepository mockrepository = mockrepository{}
	customers     = []*model.Customer{}
	currentId int = 1
)

///curtesy to gorm
type mockrepository struct{}
func (repository mockrepository) Create(customer *model.Customer) (*model.Customer, *httperors.HttpError) {
	if err := customer.Validate(); err != nil {
		return nil, err
	}
	customer.ID = currentId
	currentId++
	customers[customer.ID] = customer
	return customer, nil
}

func (repository mockrepository) GetOne(id int) (*model.Customer, *httperors.HttpError) {
	
	if customer := customers[id]; customer != nil {
		return customer, nil
	}
	return nil, httperors.NewNotFoundError(fmt.Sprintf("Customer with Id %d not found", id))
}

func (repository mockrepository) GetAll(customers []model.Customer) ([]model.Customer, *httperors.HttpError) {
	///////////////figure how to convert pointer
	// customers = cust
	return customers, nil
}

func (repository mockrepository) Update(id int, customer *model.Customer) (*model.Customer, *httperors.HttpError) {
	
	cust := customers[id]
	if cust == nil {
		return nil, httperors.NewNotFoundError("No results found")
	}
	if customer.Name == "" {
		customer.Name = cust.Name
	}
	if customer.Company == "" {
		customer.Company = cust.Company
	}
	if customer.Email == "" {
		customer.Email = cust.Email
	}
	if customer.Phone == "" {
		customer.Phone = cust.Phone
	}
	if customer.Address == "" {
		customer.Address = cust.Address
	}

	customers[id].Name = customer.Name
	customers[id].Company = customer.Company
	customers[id].Email = customer.Email
	customers[id].Phone = customer.Phone
	customers[id].Address = customer.Address
	return customer, nil
}
func (repository mockrepository) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	
	if customer := customers[id]; customer == nil {
		return nil, httperors.NewNotFoundError("No results found")
	}
	//figure how to delete from a slice
	//delete(customers, id)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}