package api

import (
	helper "project/GoTrainingAssignment/UserProfileManagement/api/helper"

	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func GetAllUser() error {
	res := []User{}
	if err := helper.Collection().Find(nil).All(&res); err != nil {
		return err
	}

	userList = res
	return nil
}

func SaveUserDb(user User) error {

	//uuidVar, _ := uuid.NewV4()
	//user.Id = uuidVar.String()
	return helper.Collection().Insert(user)
}

func DeleteUserDb(id string) error {
	fmt.Println(id)
	r:= helper.Collection().Remove(bson.M{"id": id})
	fmt.Println(r)
	return r
}
