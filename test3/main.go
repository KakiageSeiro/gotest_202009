package test3

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("ユーザが存在しません")
	userDetails = []userDetail{}
)

// 認証後取得するユーザ情報
type User struct {
	Name string
}

type userDetail struct {
	user User
	email string
	password string
}

// Auth 認証処理のインターフェース
type Auth interface {
	Signup(email, password, name string) error
	Signin(email, password string) (User, error)
	Delete(email, password string) error
}

// Authいんたーふぇーすに準拠したImplを作成して書き換えてください。・
func NewMemoryAuth() Auth {
	return AuthImpl{}
}

type AuthImpl struct{}

func (AuthImpl) Signup(email, password, name string) error {
	user := User{name}
	detail := userDetail{
		user:     user,
		email:    email,
		password: password,
	}
	userDetails = append(userDetails, detail)
	return nil
}

func (AuthImpl) Signin(email, password string) (User, error) {
	for i := range userDetails {
		if userDetails[i].email == email && userDetails[i].password == password{
			return userDetails[i].user, nil
		}
	}
	return User{}, ErrNotFound
}

func (AuthImpl) Delete(email, password string) error {
	for i := range userDetails {
		if userDetails[i].email == email && userDetails[i].password == password {
			fmt.Println(userDetails)
			fmt.Println(i)
			newUserDetails := unset(userDetails, i)
			fmt.Println(newUserDetails)
			userDetails = newUserDetails

			return nil
		}
	}
	return ErrNotFound
}

func unset(uds []userDetail, i int) []userDetail {
	if i >= len(uds) {
		return uds
	}

	newSlise := []userDetail{}
	for i2 := range uds {
		if(i != i2){
			newSlise = append(newSlise, uds[i2])
		}
	}

	return newSlise
}
