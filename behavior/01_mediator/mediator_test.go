package mediator

import (
	"fmt"
	"testing"
)

func TestMediator(t *testing.T) {
	imp := DepUserMediatorImp{}

	dep1 := Dep{}
	dep1.SetdepID("d1")
	dep1.SetdepName("d1")

	dep2 := Dep{}
	dep2.SetdepID("d2")
	dep2.SetdepName("d2")

	user1 := User{}
	user1.SetuserID("u1")
	user1.SetUserName("u1")

	user2 := User{}
	user2.SetuserID("u2")
	user2.SetUserName("u2")

	imp.depUserCol = append(imp.depUserCol, 
		DepUserMediator{
			depUserID: "1",
			depID:     dep1.depID,
			userID:    user1.userID,
	},
		DepUserMediator{
			depUserID: "2",
			depID:     dep2.depID,
			userID:    user2.userID},
		DepUserMediator{
			depUserID: "3",
			depID:     dep2.depID,
			userID:    user1.userID},
		DepUserMediator{
			depUserID: "4",
			depID:     dep1.depID,
			userID:    user2.userID},
	)

	fmt.Println("d1部门人数")
	imp.showDepUser("d1")
	fmt.Println("u1所在部门")
	imp.showUserDep("u1")

	fmt.Println("删除d1部门")
	imp.deleteOp("d1")
	fmt.Println("删除u1用户")
	imp.deleteUser("u1")

	fmt.Println("d2部门人数")
	imp.showDepUser("d2")

	fmt.Println("u2所在部门")
	imp.showUserDep("u2")
}
