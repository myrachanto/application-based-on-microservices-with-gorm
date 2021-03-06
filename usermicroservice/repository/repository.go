package repository
import (
	"log"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/myrachanto/allmicro/gormmicro/usermicroservice/httperors"
	"github.com/myrachanto/allmicro/gormmicro/usermicroservice/model"
)

type Redirectrepository interface{
	Create(user *model.User) (*model.User, *httperors.HttpError)
	Login(auser *model.LoginUser) (*model.Auth, *httperors.HttpError)
	Logout(token string) (*httperors.HttpError)
	GetOne(id int) (*model.User, *httperors.HttpError)
	GetAll(customers []model.User) ([]model.User, *httperors.HttpError)
	Update(id int, user *model.User) (*model.User, *httperors.HttpError)
	Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError)
}

/////////////////////////////////////////////////////////////////////////////////////
////////////////figure how to switch repositories automatically//////////////////////////////////
func ChooseRepo() (repository Redirectrepository) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	switch os.Getenv("DbType") {
	case "mysql":
		_, err1 := NewGormRepository()
		if err1 != nil {
			log.Fatal(err1)
		}
		repository = Sqlrepository
		// model.CheckMongo(gorm)
		return repository
	// case "mongo":
	// 	_, err1 := NewMongoRepository()
	// 	if err1 != nil {
	// 		log.Fatal(err1)
	// 	}
	// 	repository = Mongorepository
		// model.CheckMongo(mongo)
		// return repository
	// case "postgress":
	// 	repository, err1 := NewMongoRepository()
	// 	if err1 != nil {
	// 		log.Fatal(err1)
	// 	}
	// 	return repository
	// case "redis":
	// 	repository, err1 := NewMongoRepository()
	// 	if err1 != nil {
	// 		log.Fatal(err1)
	// 	}
	// 	return repository
		// 	return Repo
		// case "redis":
		// 	Repo := r.RedisRepository
		// 	return Repo
		// default:
		// 	repository = MockRepository
		// 	return repository
	}
	return
	
}
func Enkey()string{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key := os.Getenv("EncryptionKey")
	return key
}
func NewGormRepository()(Redirectrepository, error){
	dbURI := "root@/micro?charset=utf8&parseTime=True&loc=Local"
	GormDB, err1 := gorm.Open("mysql", dbURI)
	if err1 != nil {
		return nil, err1
	}
	GormDB.AutoMigrate(&model.User{})
	GormDB.AutoMigrate(&model.Auth{})
	return Sqlrepository, nil
}
// func NewMongoRepository()(Redirectrepository, error){
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	//Mongo := os.Getenv("MongoDb")
// 	host := os.Getenv("Mongohost")

// 	_, err = mgo.Dial(host)
// 	if err != nil{
// 		return nil, err
// 	}
// 	return Mongorepository, nil
// }
// func NewGormRepository()(Redirectrepository, error){
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	dbDriver := os.Getenv("DbType")
// 	dbUser	:= os.Getenv("DbUsername")
// 	dbPass := os.Getenv("DbPassword")
// 	dbName := os.Getenv("DbName")
// 	repo, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return repo, nil
// }