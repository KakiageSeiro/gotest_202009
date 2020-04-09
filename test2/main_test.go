package test2

import "testing"

var (
	u1 = user{"Taro", 1}
	u2 = user{"Hanako", 2}
	u3 = user{"Takashi", 3}
	u4 = user{"Hajime", 4}
	u5 = user{"Jiro", 5}
	u6 = user{"Kazuko", 6}
	u7 = user{"Jenny", 7}
	u8 = user{"Abe", 8}
)

// addUserでユーザ登録
func Test_addUser(t *testing.T) {
	if err := addUser(u1); err != nil {
		t.Error(err)
	}

	if err := addUser(u2); err != nil {
		t.Error(err)
	}

	if err := addUser(u3); err != nil {
		t.Error(err)
	}

	if err := addUser(u1); err != ErrAlreadyExists {
		t.Errorf("%sを期待したが、%s", ErrAlreadyExists, err)
	}

	if 3 != len(users) {
		t.Errorf("usersには3名存在しているはずが、%d名", len(users))
	}
}

// addUserで追加したユーザを今度は、syncAddUsersでまとめて登録
func Test_syncAddUsers(t *testing.T) {
	us := []user{u1, u2, u3, u3}
	if err := syncAddUsers(us); err != ErrAlreadyExists {
		t.Errorf("%sを期待したが、%s", ErrAlreadyExists, err)
	}

	if 3 != len(users) {
		t.Errorf("usersには3名存在しているはずが、%d名", len(users))
	}
}

// addUserで追加したユーザを今度は、asyncAddUsersでまとめて登録
// 内部では非同期処理の実装が求められます。
// https://golang.org/pkg/sync/
func Test_asyncAddUsers(t *testing.T) {
	us := []user{u4, u5, u8, u6, u1, u2, u3, u7, u4}
	if err := asyncAddUsers(us); err != ErrAlreadyExists {
		t.Errorf("%sを期待したが、%s", ErrAlreadyExists, err)
	}

	if 8 != len(users) {
		t.Errorf("usersには8名存在しているはずが、%d名", len(users))
	}
}
