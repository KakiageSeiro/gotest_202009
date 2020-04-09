package test3

import (
	"reflect"
	"testing"
)

var (
	email    = "test@example.com"
	password = "test1234TEST"
	name     = "test.taro"
)

func TestNewMemoryAuth(t *testing.T) {
	a := NewMemoryAuth()
	_, ok := interface{}(a).(Auth)
	if !ok {
		t.Fatal("インターフェースに準拠していません")
	}

	if err := a.Signup(email, password, name); err != nil {
		t.Error(err)
	}

	u, err := a.Signin(email, password)
	if err != nil {
		t.Error(err)
	}

	u2 := User{name}
	if !reflect.DeepEqual(u, u2) {
		t.Errorf("%vと%vは同じであるべき", u, u2)
	}

	if err := a.Delete(email, password); err != nil {
		t.Error(err)
	}

	if _, err := a.Signin(email, password); err != ErrNotFound {
		t.Errorf("エラーメッセージが、%sであるべきが%s", ErrNotFound, err)
	}
}
