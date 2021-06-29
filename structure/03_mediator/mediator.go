package mediator

import "fmt"

type Dep struct {
	depID   string
	depName string
}

func (d *Dep) GetdepID() string {
	return d.depID
}

func (d *Dep) SetdepID(depID string) {
	d.depID = depID
}

func (d *Dep) GetdepName() string {
	return d.depName
}

func (d *Dep) SetdepName(depName string) {
	d.depName = depName
}

func (d *Dep) deleteOp() bool {
	depUserMediator := DepUserMediatorImp{}
	depUserMediator.deleteOp(d.depID)
	return true
}

type User struct {
	userID   string
	userName string
}

func (u *User) GetuserID() string {
	return u.userID
}

func (u *User) SetuserID(userID string) {
	u.userID = userID
}

func (u *User) GetUserName() string {
	return u.userName
}

func (d *User) SetUserName(userName string) {
	d.userName = userName
}

func (d *User) dimisson() bool {
	depUserMediator := DepUserMediatorImp{}
	depUserMediator.deleteUser(d.userID)
	return true
}

type DepUserMediator struct {
	depUserID string
	userID    string
	depID     string
}

type DepUserMediatorImp struct {
	depUserCol []DepUserMediator
}

func (d *DepUserMediatorImp) deleteOp(depID string) bool {
	newDep := make([]DepUserMediator, len(d.depUserCol))
	for _, dep := range d.depUserCol {
        if dep.depID != depID {
			newDep = append(newDep, dep)
		}
	}
	d.depUserCol = newDep
	return true
}

func (d *DepUserMediatorImp) deleteUser(userID string) bool {
	newUser := make([]DepUserMediator, len(d.depUserCol))
	for _, dep := range d.depUserCol {
        if dep.userID != userID {
			newUser = append(newUser, dep)
		}
	}
	d.depUserCol = newUser
	return true
}

func (d *DepUserMediatorImp) showDepUser(depID string) {
	for _, dep := range d.depUserCol {
        if dep.depID == depID {
			fmt.Println("部门编号：", dep.depID, "人员编号：", dep.userID)
		}
	}
}

func (d *DepUserMediatorImp) showUserDep(userID string) {
	for _, dep := range d.depUserCol {
        if dep.userID == userID {
			fmt.Println("人员编号：", dep.userID, "部门编号：", dep.depID)
		}
	}
}
