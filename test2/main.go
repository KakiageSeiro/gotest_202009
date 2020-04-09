package test2

import (
	"errors"
	"time"
)

var users []user
var ErrAlreadyExists = errors.New("Already exists")

type user struct {
	Name string
	Age int
}

func addUser(u user) error {
	// timeは擬似的に登録に時間がかかる処理を演じているだけなので、変更不可
	time.Sleep(time.Duration(u.Age) * time.Second)
	return nil
}

func syncAddUsers(us []user) error {
	for _, u := range us {
		if err := addUser(u); err != nil {
			return err
		}
	}
	return nil
}

// asyncAddUsersユーザの追加処理が全県終わってから処理を終了させたい。直列だと遅いので、並列化する目的
func asyncAddUsers(us []user) error {
	for _, u := range us {
		// go funcの記述は残すこと、非同期であれば実装部分を変更しても可能
		go addUser(u)
	}
	return nil
}

