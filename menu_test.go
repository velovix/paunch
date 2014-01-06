package paunch

import (
	"testing"
)

func TestAddToMenu(t *testing.T) {

	menu := NewMenu()

	err := menu.AddItem("Play")
	if err != nil {
		t.Errorf("menu.AddItem(\"Play\") returned %s", err)
	}

	err = menu.AddItem("Instructions")
	if err != nil {
		t.Errorf("menu.AddItem(\"Instructions\") returned %s", err)
	}

	err = menu.SelectItem("Play")
	if err != nil {
		t.Errorf("menu.SelectItem(\"Play\") returned %s", err)
	}

	err = menu.SelectItem("Instructions")
	if err != nil {
		t.Errorf("menu.SelectItem(\"Instructions\") returned %s", err)
	}
}

func TestMoveSelector(t *testing.T) {

	menu := NewMenu()

	err := menu.AddItem("Play")
	if err != nil {
		t.Errorf("menu.AddItem(\"Play\") returned %s", err)
	}

	err = menu.AddItem("Instructions")
	if err != nil {
		t.Errorf("menu.AddItem(\"Instructions\") returned %s", err)
	}

	err = menu.SetItemNeighbor("Play", "Instructions", Left)
	if err != nil {
		t.Errorf("menu.SetItemNeighbor(\"Play\", \"Instructions\", Left) returned %s", err)
	}

	err = menu.SelectItem("Play")
	if err != nil {
		t.Errorf("menu.SelectItem(\"Play\") returned %s", err)
	}

	var moved bool
	moved, err = menu.MoveSelector(Left)
	if err != nil {
		t.Errorf("menu.MoveSelector(Left) returned %s", err)
	}
	if !moved {
		t.Errorf("menu.MoveSelector(Left) returned false when true was expected")
	}
	if menu.GetSelection() != "Instructions" {
		t.Errorf("menu.GetSelection returned %s when \"Instructions\" was expected", menu.GetSelection)
	}
}

func TestDeleteItem(t *testing.T) {

	menu := NewMenu()

	err := menu.AddItem("Play")
	if err != nil {
		t.Errorf("menu.AddItem(\"Play\") returned %s", err)
	}

	err = menu.AddItem("Instructions")
	if err != nil {
		t.Errorf("menu.AddItem(\"Instructions\") returned %s", err)
	}

	err = menu.SetItemNeighbor("Play", "Instructions", Left)
	if err != nil {
		t.Errorf("menu.SetItemNeighbor(\"Play\", \"Instructions\", Left) returned %s", err)
	}

	err = menu.SelectItem("Play")
	if err != nil {
		t.Errorf("menu.SelectItem(\"Play\") returned %s", err)
	}

	err = menu.DeleteItem("Instructions")
	if err != nil {
		t.Errorf("menu.DeleteItem(\"Instructions\") returned %s", err)
	}

	err = menu.SelectItem("Instructions")
	if err == nil {
		t.Errorf("menu.SelectItem(\"Instructions\") returned no error when an error was expected")
	}

	var moved bool
	moved, err = menu.MoveSelector(Left)
	if err != nil {
		t.Errorf("menu.MoveSelector(Left) returned %s", err)
	}
	if moved {
		t.Errorf("menu.MoveSelector(Left) returned true when false was expected")
	}
}
