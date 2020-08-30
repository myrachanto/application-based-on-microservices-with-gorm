package repository

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/myrachanto/allmicro/gormmicro/usermicroservice/httperors"
	"github.com/myrachanto/allmicro/gormmicro/usermicroservice/model"
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

	GormDB.AutoMigrate(&model.User{})
	GormDB.AutoMigrate(&model.Auth{})
	return GormDB, nil
}
func DbClose(GormDB *gorm.DB) {
	defer GormDB.Close()
}
func (repository sqlrepository) Create(user *model.User) (*model.User, *httperors.HttpError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	ok, err1 := user.ValidatePassword(user.Password)
	if !ok {
		return nil, err1
	}
	ok = user.ValidateEmail(user.Email)
	if !ok {
		return nil, httperors.NewNotFoundError("Your email format is wrong!")
	}
	ok = repository.UserExist(user.Email)
	if ok {
		return nil, httperors.NewNotFoundError("Your email already exists!")
	}
	hashpassword, err2 := user.HashPassword(user.Password)
	if err2 != nil {
		return nil, err2
	}
	user.Password = hashpassword
	//  encyKey := Enkey()
	// // p := support.Encrypt([]byte(user.Password), encyKey)
	// // user.Password = string(p)
	// user.Password = support.Hash(encyKey,user.Password)
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	
	fmt.Println(user)
	GormDB.Create(&user)
	DbClose(GormDB)
	return user, nil
}
func (repository sqlrepository) Login(auser *model.LoginUser) (*model.Auth, *httperors.HttpError) {
	if err := auser.Validate(); err != nil {
		return nil, err
	}
	ok := repository.UserExist(auser.Email)
	if !ok {
		return nil, httperors.NewNotFoundError("Your email does not exists!")
	}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	user := model.User{}
	GormDB.Model(&user).Where("email = ?", auser.Email).First(&user)
	ok = user.Compare(auser.Password, user.Password)
	if !ok {
		return nil, httperors.NewNotFoundError("wrong email password combo!")
	}
	tk := &model.Token{
		UserID: user.ID,
		UserName:   user.UName,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: model.ExpiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	encyKey := Enkey()
	tokenString, error := token.SignedString([]byte(encyKey))
	if error != nil {
		fmt.Println(error)
	}
	
	auth := &model.Auth{UserID:user.ID, Token:tokenString}
	GormDB.Create(&auth)
	DbClose(GormDB)
	
	return auth, nil
}
func (repository sqlrepository) Logout(token string) (*httperors.HttpError) {
	auth := model.Auth{}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return err1
	}
	if GormDB.First(&auth, "token =?", token).RecordNotFound(){
		return httperors.NewNotFoundError("Something went wrong logging out!")
	 }
	
	GormDB.Model(&auth).Where("token =?", token).First(&auth)
	
	GormDB.Delete(auth)
	DbClose(GormDB)
	
	return  nil
}
func (repository sqlrepository) GetOne(id int) (*model.User, *httperors.HttpError) {
	ok := repository.UserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("User with that id does not exists!")
	}
	user := model.User{}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	
	GormDB.Model(&user).Where("id = ?", id).First(&user)
	DbClose(GormDB)
	
	return &user, nil
}

func (repository sqlrepository) GetAll(users []model.User) ([]model.User, *httperors.HttpError) {
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	User := model.User{}
	GormDB.Model(&User).Find(&users)
	
	DbClose(GormDB)
	if len(users) == 0 {
		return nil, httperors.NewNotFoundError("No results found!")
	}
	return users, nil
}

func (repository sqlrepository) Update(id int, user *model.User) (*model.User, *httperors.HttpError) {
	ok := repository.UserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("User with that id does not exists!")
	}
	
	
	hashpassword, err2 := user.HashPassword(user.Password)
	if err2 != nil {
		return nil, err2
	}
	user.Password = hashpassword
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	User := model.User{}
	uuser := model.User{}
	
	GormDB.Model(&User).Where("id = ?", id).First(&uuser)
	if user.FName  == "" {
		user.FName = uuser.FName
	}
	if user.LName  == "" {
		user.LName = uuser.LName
	}
	if user.UName  == "" {
		user.UName = uuser.UName
	}
	if user.Phone  == "" {
		user.Phone = uuser.Phone
	}
	if user.Address  == "" {
		user.Address = uuser.Address
	}
	if user.Picture  == "" {
		user.Picture = uuser.Picture
	}
	if user.Email  == "" {
		user.Email = uuser.Email
	}
	if user.Password  == "" {
		user.Password = uuser.Password
	}
	GormDB.Model(&User).Where("id = ?", id).First(&User).Update(&user)
	
	DbClose(GormDB)

	return user, nil
}
func (repository sqlrepository) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := repository.UserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("User with that id does not exists!")
	}
	user := model.User{}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&user).Where("id = ?", id).First(&user)
	GormDB.Delete(user)
	DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (repository sqlrepository)UserExist(email string) bool {
	user := model.User{}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return false
	}
	if GormDB.First(&user, "email =?", email).RecordNotFound(){
	   return false
	}
	DbClose(GormDB)
	return true
	
}
func (repository sqlrepository)UserExistByid(id int) bool {
	user := model.User{}
	GormDB, err1 := Getconnected()
	if err1 != nil {
		return false
	}
	if GormDB.First(&user, "id =?", id).RecordNotFound(){
	   return false
	}
	DbClose(GormDB)
	return true
	
}
