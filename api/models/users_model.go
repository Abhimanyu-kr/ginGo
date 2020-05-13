package models

import (
//   "github.com/gin-gonic/gin"
	"log"
	"errors"
	"html"
	// "fmt"
	"strings"
	"time"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"ginGo/api/utils/validator"
  

	// "ginGo/api/auth"
	// "ginGo/api/models"
	// "ginGo/api/responses"
	// "ginGo/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
//   "ginGo/models"
)

type Response struct {
	Message string
	Status string
    Nrows int32
}

type TUsers struct{
	User_id        uint32 `gorm:"primary_key;auto_increment" form:"user_id"`
    Username     string    `gorm:"size:100;not null;unique" form:"username"`
    Email  string    `gorm:"size:255" form:"email"`
	Password string    `gorm:"size:100;not null;" form:"password"`
	Mobile string    `gorm:"size:12;not null;" form:"mobile"`
	Country_code string    `gorm:"size:100;not null;" form:"country_code"`
	Country string    `gorm:"size:100;not null;" form:"country"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" form:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" form:"updated_at"`
}


func (u *TUsers) Prepare() {
	
	u.User_id = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.Country = html.EscapeString(strings.TrimSpace(u.Country))
	u.Country_code = html.EscapeString(strings.TrimSpace(u.Country_code))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *TUsers) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
func (u *TUsers) Validate(action string) error {
	
	switch strings.ToLower(action) {
	case "update":
		if err :=  validator.ValidateString("Username", u.Username, 5, 50, ""); err != nil{
			return err
		}
		if u.Username == "" {
			return errors.New("Required Username")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Mobile == "" {
			return errors.New("Required Mobile")
		}
		if u.Country == "" {
			return errors.New("Required Country")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		
		if (u.Username == "" && u.Email == "" ) {
			return errors.New("Required Username or Email")
		}
		if err :=  validator.ValidateString("Password", u.Password, 5, 100, ``); err != nil{
			return err
		}
		return nil

	default:
		if err :=  validator.ValidateString("Username", u.Username, 5, 100, ``); err != nil{
			return err
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		if err :=  validator.ValidateString("Country", u.Country, 2, 100, ``); err != nil{
			return err
		}
		if err :=  validator.ValidateMobile("Mobile", u.Mobile, 10, 12, `^[0-9]{10}$`); err != nil{
			return err
		}
		if err :=  validator.ValidateString("Country Code", u.Country_code, 1, 5,``); err != nil{
			return err
		}
		if err :=  validator.ValidateString("Password", u.Password, 5, 100, ``); err != nil{
			return err
		}
		return nil
	}
}

func (u *TUsers)CreateNewUser(db *gorm.DB) (*TUsers, error)  {

	var err error

	// To hash the password
	err = u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	// save real hassed password
	var password string = u.Password

	err = db.Debug().Table("tUsers").Create(&u).Error
	if err != nil {
		return &TUsers{}, err
	}
	
	// save real hassed password
	err =  db.Debug().Table("tUsers").Where("user_id = ?", u.User_id).Update("password", password).Error
	if err != nil {
		return &TUsers{}, err
	}
	log.Println(u.User_id)
	
	return u, nil
}

// CreateUser : insert user detail in data base
func CreateUser( username string,
				email string,  
				mobile string,
				password string,
				country string ) (Response , error) {

		
	log.Println("IN create function")
	log.Println(username)

	// var err error

	

	// err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	// if err != nil {
	// 	return "", err
	// }
	// err = models.VerifyPassword(user.Password, password)
	// if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
	// 	return "", err
	// }

	
	
	return Response{
		Message: "YOO YOOO !",
		Status:     "SUCCESS",
		Nrows:    0,
	},nil
}


