package test4

import "errors"

var (
	ErrFire = errors.New("Error")
)

// deferを使う
func deferTest() error {
	defer fire()
	return errors.New("ダミーエラー")
}

func fire() error {
	return errors.New("エラー")
}
