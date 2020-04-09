package test1

import "errors"

var Err1 = errors.New("エラーメッセージ")

func returnErr(i int) error {
	if i == 0 {
		return Err1
	}

	return nil
}
