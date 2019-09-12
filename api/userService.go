package api

import (
	"fmt"
)

func NewUser() User {
	return User{}
}

func GetUserList() []User {
	if userList == nil || len(userList) == 0 {
		userList = []User{}
		fmt.Println("helloinside")
		return userList
	}
	fmt.Println("hello")
	return userList
}
func AddUserService(user User) {

	SaveUserDb(user)
	GetAllUser()
}

func DeleteService(id string){
	fmt.Println("From SERVICE",id)
	DeleteUserDb(id)
	GetAllUser()
}