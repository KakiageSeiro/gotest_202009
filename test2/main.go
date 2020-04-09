package test2

import (
	"errors"
	"sync"
	"time"
)

var users []user
var ErrAlreadyExists = errors.New("Already exists")

type user struct {
	Name string
	Age  int
}

type addUserCommand struct {
	user user
	wg *sync.WaitGroup
	ch chan error
}

func facadeAddUser(command addUserCommand) {
	if command.wg != nil {
		defer command.wg.Done()
	}

	command.ch <- addUser(command.user)
}

func addUser(u user) error {
	// timeは擬似的に登録に時間がかかる処理を演じているだけなので、変更不可
	time.Sleep(time.Duration(u.Age) * time.Second)

	//同じユーザーががすでにいる場合はエラー
	for i := range users {
		if users[i].Name == u.Name && users[i].Age == u.Age {
			return ErrAlreadyExists
		}
	}

	//ユーザーを追加
	users = append(users, u)
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

	wg := new(sync.WaitGroup)
	wg.Add(len(us))

	//各ユーザーを登録する
	ch := make(chan error, len(us))
	for _, u := range us {
		// go funcの記述は残すこと、非同期であれば実装部分を変更しても可能
		//go addUser(u)

		command := addUserCommand{u, wg, ch}
		go facadeAddUser(command)
	}

	wg.Wait()

	for err := range ch {
		if err != nil {
			return err
		}
	}

	return nil
}
