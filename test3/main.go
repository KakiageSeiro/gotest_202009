package test3

import "errors"

var (
	ErrNotFound = errors.New("ユーザが存在しません")

)

// 認証後取得するユーザ情報
type User struct {
	Name string
}

// Auth 認証処理のインターフェース
type Auth interface {
	Signup(email, password, name string) error
	Signin(email, password string) (User, error)
	Delete(email, password string) error
}

// Authいんたーふぇーすに準拠したImplを作成して書き換えてください。・
func NewMemoryAuth() Auth {
	return nil
}
