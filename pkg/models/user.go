package models

import (
	"net/http"
	"time"

	"github.com/adeindra6/intikom_test/pkg/config"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string `gorm:"type:text"json:"name"`
	Email    string `gorm:"type:varchar(255)"json:"email"`
	Password string `gorm:"type:text"json:"password,omitempty"`
}

type LoginRes struct {
	Token  string    `gorm:""json:"token"`
	Status string    `json:"status"`
	Code   int64     `json:"code"`
	Exp    time.Time `json:"exp"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	password := []byte(u.Password)
	hashPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashPassword)

	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("id = ?", id).Find(&getUser)
	return &getUser, db
}

func DeleteUserById(id int64) User {
	var user User
	db.Where("id = ?", id).Delete(&user)
	return user
}

func (u *User) Login() *LoginRes {
	var LoginUser User
	var PasswordCorrect bool
	var Token LoginRes
	db.Where("email = ?", u.Email).Find(&LoginUser)

	password := []byte(u.Password)
	storedPassword := []byte(LoginUser.Password)
	PasswordCorrect = false

	err := bcrypt.CompareHashAndPassword(storedPassword, password)
	if err == nil {
		PasswordCorrect = true
	}

	if PasswordCorrect {
		expTime := time.Now().Add(time.Hour * 24)
		jwtToken, err := CreateJWTToken(LoginUser.ID, LoginUser.Email, expTime)
		if err != nil {
			panic(err)
		}

		Token.Token = jwtToken
		Token.Status = "SUCCESS"
		Token.Code = http.StatusOK
		Token.Exp = expTime
	}

	return &Token
}

func CreateJWTToken(id uint, email string, expTime time.Time) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["customer_id"] = id
	atClaims["username"] = email
	atClaims["exp"] = expTime.Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)

	token, err := at.SignedString([]byte("Intikom_Test"))
	if err != nil {
		return "", err
	}

	return token, nil
}
