package main

import (
	"APS_Project/database"
	"testing"
	"time"
)

var (
	dummyGame  = Game{ChatData: "Some chat data!", DateTime: time.Now()}
	updateGame = Game{ChatData: "Some new chat data!", DateTime: dummyGame.DateTime}
)

func TestCreateGame(t *testing.T) {
	err := CreateGame(database.GlobalDB, &dummyGame)

	if err != nil {
		t.Error("CreateGame failed: ", err)
	}

	updateGame.ID = dummyGame.ID
}

func TestReadGame(t *testing.T) {
	game, err := ReadGame(database.GlobalDB, Game{ID: dummyGame.ID})

	if err != nil {
		t.Error("ReadGame failed: ", err)
	}

	if game.ChatData != dummyGame.ChatData {
		t.Error("Game read not valid: ", game, Game{ID: dummyGame.ID})
	}
}

func TestUpdateGame(t *testing.T) {
	err := UpdateGame(database.GlobalDB, updateGame)

	if err != nil {
		t.Error("UpdateGame failed: ", err)
	}

	game, err := ReadGame(database.GlobalDB, Game{ID: dummyGame.ID})

	if err != nil {
		t.Error("ReadGame failed: ", err)
	}

	if game.ChatData != updateGame.ChatData {
		t.Error("UpdateGame not valid: ", game, updateGame)
	}
}

func TestDeleteGame(t *testing.T) {
	err := DeleteGame(database.GlobalDB, Game{ID: dummyGame.ID})

	if err != nil {
		t.Error("DeleteGame failed: ", err)
	}
}
