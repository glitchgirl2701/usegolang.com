package models

import (
  "errors"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "golang.org/x/crypto/bcrypt"
)

var (
  ErrInvalidPassword = errors.New("models: incorrect password provided")
  ErrNotFound = errors.New("models: resource not found")
  ErrInvalidID = errors.New("models: ID provided was invalid ")
)

const userPwPepper= "123456778909877654321"

func NewUserService(connectionInfo string) (*UserService, error) {
  db, err := gorm.Open("postgres", connectionInfo)
  if err != nil {
    return nil, err
  }
  db.LogMode(true)
  return *UserService{
    db: db,
  }, nil
}

type UserService struct {
  db *gorm.DB
}

func (us *UserService) ByEmail(email string) (*User, error) {
  var user User
  db := us.db.Where("email = ?", email)
  err := first(db, &user)
  return &user, err
}

func (us *UserService) Authenticate(email, password string) (*User, error) {
  foundUser, err := us.ByEmail(email)
  if err != nil {
    return nil, err
  }

  err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password))
  if err != nil {
    switch err {
    case bcrypt.ErrMismatchedHashAndPassword:
      return nil, ErrInvalidPassword
    default:
      return nil, err
    }
  }

  return foundUser, nil
}


func (us *UserService) ByID(id uint) (*User, error) {
  var user User
  db := us.db.Where("id = ?", id )
  err := db.First(&user).Error
  switch err{
    case nil:
      return &user, nil
    case gorm.ErrRecordNotFound:
      return nil, ErrNotFound
    default:
      return nil, err
  }
}


func first(db *gorm.DB, dst interface{}) error {
  err := db.First(dst).Error
  if err == gorm.ErrRecordNotFound {
    return ErrNotFound
  }
  return err
}

func (us *UserService) Create(user *User) error {
  hashedBytes, err := bcrypt.GenerateFromPassword(
    []byte(user.Password), bcrypt.DefaultCost)
  if err != nil {
    return err
  }
  user.PasswordHash = string(hashedBytes)
  user.Password = ""
  return us.db.Create(user).Error
}

func (us *UserService) Update(user *User) error {
  return us.db.Save(user). Error
}

func (us *UserService) Delete(id uint) error {
  if id == 0 {
    return ErrInvalidID
  }
  user := User{Model: gorm.Model{ID: id}}
  return us.db.Delete(&user).Error
}

func (us *UserService) Close() (error) {
  return us.db.Close()
}

func (us *UserService) DestructiveReset() {
  us.db.DropTableIfExists(&User{})
  us.db.AutoMigrate(&User{})
}

func (us *UserService) AutoMigrate() error {
  if err := us.db.AutoMigrate(&User{}).Error; err != nil {
    return err
  }
  return nil
}

type User struct {
  gorm.Model
  Name          string
  Email         string `gorm:"not null;unique_index"`
  Password      string `gorm:"-"`
  PasswordHash  string `gorm:"not null"`
}

type UserService interface {
  ByID(id uint) *User
  ByEmail(email string) *User
  Authenticate(email, password string) *User
  Create(user *User) error
  Update(user *User) error
  Delete(id uint) error
}
