package main

import (
	"APS_Project/database"
	"testing"
)

var (
	dummyUser  = User{Name: "Stefan", Surname: "Mitic", Email: "corcrash@gmail.com"}
	updateUser = User{Name: "Milena", Surname: "Petrovic", Email: "menci@gmail.com", Google: true}
)

func TestCreateUser(t *testing.T) {
	err := CreateUser(database.GlobalDB, &dummyUser)

	if err != nil {
		t.Error("CreateUser failed: ", err)
	}

	updateUser.ID = dummyUser.ID
}

func TestReadUser(t *testing.T) {
	user, err := ReadUser(database.GlobalDB, User{ID: dummyUser.ID})

	if err != nil {
		t.Error("ReadUser failed: ", err)
	}

	if user.Email != dummyUser.Email {
		t.Error("User read not valid: ", user, dummyUser)
	}
}

func TestUpdateUser(t *testing.T) {
	err := UpdateUser(database.GlobalDB, updateUser)

	if err != nil {
		t.Error("UpdateUser failed: ", err)
	}

	user, err := ReadUser(database.GlobalDB, User{ID: dummyUser.ID})

	if err != nil {
		t.Error("ReadUser failed: ", err)
	}

	if user.Email != updateUser.Email {
		t.Error("UpdateUser not valid: ", user, updateUser)
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser(database.GlobalDB, User{ID: dummyUser.ID})

	if err != nil {
		t.Error("DeleteUser failed: ", err)
	}
}
