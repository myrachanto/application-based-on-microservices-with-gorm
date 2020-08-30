package model

import(
	"time"
	"github.com/myrachanto/allmicro/gormmicro/customermicroservice/httperors"
)

type Customer struct {
	Name string `gorm:"not null" `
	Company string `gorm:"not null"`
	Phone string `gorm:"not null"`
	Address string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Invoice []Invoice `gorm:"foreignkey:UserRefer"`//has many invoices
	Base
}

type Invoice struct {
	CustomerID uint64 `gorm:"not null"`
	Customer Customer `gorm:"foreignKey:CustomerID; not null"`
	Title string `gorm:"not null"`
	Dated time.Time `gorm:"not null"`
	Due_date time.Time `gorm:"not null"`
	Discount float64 `gorm:"not null"`
	Sub_total float64 `gorm:"not null"`
	Total float64 `gorm:"not null"`
	InvoiceItem []InvoiceItem `gorm:"foreignkey:UserRefer"`//has many invoiceitems
	Base
}
type InvoiceItem struct {
	InvoiceID uint64 `gorm:"not null"`
	Invoice Invoice `gorm:"foreignKey:InvoiceID; not null"`
	Description string `gorm:"not null"`
	Qty uint64 `gorm:"not null"`
	Unit_price float64 `gorm:"not null"`
	Base
}
type Base struct{
	ID int `gorm:"id" json:"id"`
	Created_At time.Time `gorm:"created_at"`
	Updated_At time.Time `gorm:"updated_at"`
	Delete_At *time.Time `gorm:"deleted_at"`

}
func (customer Customer) Validate() *httperors.HttpError{
	if customer.Name == "" {
		return httperors.NewNotFoundError("Invalid Name")
	}
	if customer.Company == "" {
		return httperors.NewNotFoundError("Invalid Company")
	}
	if customer.Phone == "" {
		return httperors.NewNotFoundError("Invalid Phone")
	}
	if customer.Email == "" {
		return httperors.NewNotFoundError("Invalid Email")
	}
	if customer.Address == "" {
		return httperors.NewNotFoundError("Invalid Address")
	}
	return nil
}