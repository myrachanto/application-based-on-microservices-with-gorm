package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/myrachanto/allmicro/gormmicro/categorymicroservice/httperors"
	"github.com/myrachanto/allmicro/gormmicro/categorymicroservice/model"
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

	GormDB.AutoMigrate(&model.Category{})
	GormDB.AutoMigrate(&model.MajorCategory{})
	GormDB.AutoMigrate(&model.Category{})
	GormDB.AutoMigrate(&model.SubCategory{})
	return GormDB, nil
}
func DbClose(GormDB *gorm.DB) {
	defer GormDB.Close()
}
func (repository sqlrepository) Create(category *model.Category) (*model.Category, *httperors.HttpError) {
	if err := category.Validate(); err != nil {
		return nil, err
	}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Create(&category)
	DbClose(GormDB)
	return category, nil
}
func (repository sqlrepository) GetOne(id int) (*model.Category, *httperors.HttpError) {
	ok := repository.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("category with that id does not exists!")
	}
	category := model.Category{}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	
	GormDB.Model(&category).Where("id = ?", id).First(&category)
	DbClose(GormDB)
	
	return &category, nil
}

func (repository sqlrepository) GetAll(categorys []model.Category) ([]model.Category, *httperors.HttpError) {
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	Category := model.Category{}
	GormDB.Model(&Category).Find(&categorys)
	
	DbClose(GormDB)
	if len(categorys) == 0 {
		return nil, httperors.NewNotFoundError("No results found!")
	}
	return categorys, nil
}

func (repository sqlrepository) Update(id int, category *model.Category) (*model.Category, *httperors.HttpError) {
	ok := repository.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("category with that id does not exists!")
	}
	
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	Category := model.Category{}
	acategory := model.Category{}
	
	GormDB.Model(&Category).Where("id = ?", id).First(&acategory)
	if category.Name  == "" {
		category.Name = acategory.Name
	}
	if category.Title  == "" {
		category.Title = acategory.Title
	}
	if category.Description  == "" {
		category.Description = acategory.Description
	}
	GormDB.Model(&Category).Where("id = ?", id).First(&Category).Update(&acategory)
	
	DbClose(GormDB)

	return category, nil
}
func (repository sqlrepository) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := repository.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("Product with that id does not exists!")
	}
	category := model.Category{}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&category).Where("id = ?", id).First(&category)
	GormDB.Delete(category)
	DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (repository sqlrepository)ProductUserExistByid(id int) bool {
	category := model.Category{}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return false
	}
	if GormDB.First(&category, "id =?", id).RecordNotFound(){
	   return false
	}
	DbClose(GormDB)
	return true
	
}