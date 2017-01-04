package main

import (
	"APS_Project/database"
	"testing"
)

var (
	dummyTerm  = Term{Term: "Term!", Explanation: "Explanation!"}
	updateTerm = Term{Explanation: "New Explanation!"}
)

func TestCreateTerm(t *testing.T) {
	err := CreateTerm(database.GlobalDB, &dummyTerm)

	if err != nil {
		t.Error("CreateTerm failed: ", err)
	}
}

func TestReadTerm(t *testing.T) {
	term, err := ReadTerm(database.GlobalDB, Term{ID: dummyTerm.ID})

	if err != nil {
		t.Error("ReadTerm failed: ", err)
	}

	if term.Explanation != dummyTerm.Explanation {
		t.Error("Term read not valid: ", term, dummyTerm)
	}
}

func TestUpdateTerm(t *testing.T) {
	err := UpdateTerm(database.GlobalDB, updateTerm)

	if err != nil {
		t.Error("UpdateTerm failed: ", err)
	}

	term, err := ReadTerm(database.GlobalDB, Term{ID: dummyTerm.ID})

	if err != nil {
		t.Error("ReadTerm failed: ", err)
	}

	if term.Explanation != updateTerm.Explanation {
		t.Error("UpdateTerm not valid: ", term, updateTerm)
	}
}

func TestDeleteTerm(t *testing.T) {
	err := DeleteTerm(database.GlobalDB, Term{ID: dummyTerm.ID})

	if err != nil {
		t.Error("DeleteTerm failed: ", err)
	}
}
