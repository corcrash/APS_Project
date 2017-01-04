package main

import (
	"APS_Project/database"
	"testing"
)

var (
	dummyDrawing  = Drawing{Answer: "Answer!", RoundNumber: 1, Data: "Data!"}
	updateDrawing = Drawing{Data: "New data!"}
)

func TestCreateDrawing(t *testing.T) {
	dummyDrawing.Game, _ = ReadGame(database.GlobalDB, Game{ID: 1})
	err := CreateDrawing(database.GlobalDB, &dummyDrawing)

	if err != nil {
		t.Error("CreateDrawing failed: ", err)
	}
}

func TestReadDrawing(t *testing.T) {
	drawing, err := ReadDrawing(database.GlobalDB, Drawing{ID: dummyDrawing.ID})

	if err != nil {
		t.Error("ReadDrawing failed: ", err)
	}

	if drawing.Data != dummyDrawing.Data {
		t.Error("Drawing read not valid: ", drawing, dummyDrawing)
	}
}

func TestUpdateDrawing(t *testing.T) {
	err := UpdateDrawing(database.GlobalDB, updateDrawing)

	if err != nil {
		t.Error("UpdateDrawing failed: ", err)
	}

	drawing, err := ReadDrawing(database.GlobalDB, Drawing{ID: dummyDrawing.ID})

	if err != nil {
		t.Error("ReadDrawing failed: ", err)
	}

	if drawing.Data != updateDrawing.Data {
		t.Error("UpdateDrawing not valid: ", drawing, updateDrawing)
	}
}

func TestDeleteDrawing(t *testing.T) {
	err := DeleteDrawing(database.GlobalDB, Drawing{ID: dummyDrawing.ID})

	if err != nil {
		t.Error("DeleteDrawing failed: ", err)
	}
}
