package db

import (
	"fmt"
	"os"
	"sync"
	"time"

	guuid "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB //database
var once sync.Once

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

// Postgre cocnnection start
func InitPostgre() {
	username := getEnvVariable("LOCALHOST_USERNAME_DB")
	password := getEnvVariable("LOCALHOST_PASSWORD_DB")
	dbName := getEnvVariable("LOCALHOST_DB_NAME")
	dbHost := getEnvVariable("LOCALHOST_HOST_DB")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{}, &Company{}, &Division{}, &JobPosition{}, &AppToken{}, &ApiLog{}, &ApiType{}, &Activity{}, &Workflow{}, &ActivityLink{}, &ActivityLog{}, &AuditTrails{})

	company := Company{Name: "AIA", Email: "aia@aia.com", CreatedAt: time.Now()}
	db.FirstOrCreate(&company)
	apptoken := AppToken{Company: company}
	db.FirstOrCreate(&apptoken)
	fmt.Println("Successfully connected!")
}

//returns a handle to the DB object
func GetDb() *gorm.DB {
	// Initiate value if there is no instance
	once.Do(func() {
		InitPostgre()
	})
	return db
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", guuid.New().String())
	scope.SetColumn("IsActive", true)
	return nil
}
func (company *Company) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", guuid.New().String())
	scope.SetColumn("IsActive", true)
	return nil
}
func (apptoken *AppToken) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Token", guuid.New().String())
	scope.SetColumn("IsActive", true)
	return nil
}
func (workflow *Workflow) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", guuid.New().String())
	scope.SetColumn("IsActive", true)
	return nil
}
func (useractivity *UserActivity) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", guuid.New().String())
	scope.SetColumn("IsActive", true)
	return nil
}
