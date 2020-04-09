package test1

import "errors"

var Err1 = errors.New("エラーメッセージ")

func returnErr(i int) error {
	_ = i
	return nil
}
