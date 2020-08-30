package model

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/myrachanto/allmicro/gormmicro/invoicemicroservice/httperors"
)

type Product struct {
	Name string `gorm:"not null"`
	Title string `gorm:"not null"`
	Description string `gorm:"not null"`
	//SubCategory SubCategory `gorm:"foreignKey:UserID; not null"`
	SubCategoryID uint `gorm:"not null"`
	Picture string 
	gorm.Model
}
type Customer struct {
	Name string `gorm:"not null"`
	Company string `gorm:"not null"`
	Phone string `gorm:"not null"`
	Address string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Invoice []Invoice `gorm:"foreignkey:UserRefer"`//has many invoices
	gorm.Model
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
	PaidStatus bool 
	InvoiceItem []InvoiceItem `gorm:"foreignkey:UserRefer"`//has many invoiceitems
	gorm.Model
}
type InvoiceItem struct {
	InvoiceID uint64 `gorm:"not null"`
	Invoice Invoice `gorm:"foreignKey:InvoiceID; not null"`
	Name string `gorm:"not null"`
	Description string `gorm:"not null"`
	Qty uint64 `gorm:"not null"`
	Unit_price float64 `gorm:"not null"`
	gorm.Model
}

func (invoice Invoice) Validate() *httperors.HttpError{
	if invoice.Title == "" && len(invoice.Title) < 5 {
		return httperors.NewNotFoundError("Invalid title")
	}
	if invoice.Total < 0{
		return httperors.NewNotFoundError("Invalid total")
	}
	if invoice.CustomerID < 0{
		return httperors.NewNotFoundError("Invalid customer")
	}
	return nil
}
