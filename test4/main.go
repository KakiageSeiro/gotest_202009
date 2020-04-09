package test4

import "errors"

var (
	ErrFire = errors.New("Error")
)

// deferを使う
func deferTest()(err error){
	//defer fire()
	defer func(){
		err2 := fire()
		if err2 != nil {
			err = ErrFire
		}
	}()
	return errors.New("ダミーエラー")
}

func fire() error {
	return errors.New("エラー")
}
