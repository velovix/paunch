package paunch

import (
	"fmt"
)

type menuItem struct {
	name        string
	surrounding map[Direction]menuItem
}

// Menu is an object to help with making in-game menus.
type Menu struct {
	currItem string
	items    map[string]menuItem
}

// NewMenu creates a new Menu object.
func NewMenu() Menu {

	var menu Menu
	menu.items = make(map[string]menuItem)

	return menu
}

// AddItem adds a new menu item to the Menu object, or returns an error if the
// item name has already been used.
func (menu *Menu) AddItem(name string) error {

	if _, ok := menu.items[name]; ok {
		return fmt.Errorf("method AddItem: menu item %s already exists", name)
	}

	menu.items[name] = menuItem{name: name, surrounding: make(map[Direction]menuItem)}
	return nil
}

// SetItemNeighbor sets the adjacet item of the first specified item. If either
// items do not exist in the Menu object, the method returns an error.
func (menu *Menu) SetItemNeighbor(name, neighbor string, direction Direction) error {

	if _, ok := menu.items[name]; !ok {
		return fmt.Errorf("method SetItemNeighbor: menu item %s doesn't exist", name)
	}

	if _, ok := menu.items[neighbor]; !ok {
		return fmt.Errorf("method SetItemNeighbor: menu item neighbor %s doesn't exist", neighbor)
	}

	menu.items[name].surrounding[direction] = menu.items[neighbor]
	return nil
}

// DeleteItem removes an item from the Menu object, and removes any instances
// of the specified item being a neighbor to another item. If the currently
// selected item is the item to be deleted, no item will be selected.
func (menu *Menu) DeleteItem(name string) error {

	if _, ok := menu.items[name]; !ok {
		return fmt.Errorf("method DeleteItem: menu item %s doesn't exist", name)
	}

	if menu.currItem == name {
		menu.currItem = ""
	}

	for i := range menu.items {
		if val, ok := menu.items[i].surrounding[Left]; ok && val.name == menu.items[name].name {
			delete(menu.items[i].surrounding, Left)
		}
		if val, ok := menu.items[i].surrounding[Right]; ok && val.name == menu.items[name].name {
			delete(menu.items[i].surrounding, Right)
		}
		if val, ok := menu.items[i].surrounding[Up]; ok && val.name == menu.items[name].name {
			delete(menu.items[i].surrounding, Up)
		}
		if val, ok := menu.items[i].surrounding[Down]; ok && val.name == menu.items[name].name {
			delete(menu.items[i].surrounding, Down)
		}
	}

	delete(menu.items, name)

	return nil
}

// SelectItem sets the specified item to be the currently selected item, or
// returns an error if the item does not exist in the Menu object.
func (menu *Menu) SelectItem(name string) error {

	if _, ok := menu.items[name]; !ok {
		return fmt.Errorf("method SelectItem: menu item %s doesn't exist", name)
	}

	menu.currItem = name
	return nil
}

// Selection returns the currently selected item's name, or a blank string
// if nothing is selected.
func (menu *Menu) Selection() string {

	return menu.currItem
}

// MoveSelector moves the selector in the specified direction. It returns an
// error if no item is currently selected or if the currently selected item no
// longer exists. It returns a boolean value representing whether or not the
// selector was actually moved.
func (menu *Menu) MoveSelector(direction Direction) (bool, error) {

	if menu.currItem == "" {
		return false, fmt.Errorf("method MoveSelector: no current item set")
	}

	if _, ok := menu.items[menu.currItem]; !ok {
		return false, fmt.Errorf("method MoveSelector: current item %s doesn't exist", menu.currItem)
	}

	if _, ok := menu.items[menu.currItem].surrounding[direction]; !ok {
		return false, nil
	}

	menu.currItem = menu.items[menu.currItem].surrounding[direction].name
	return true, nil
}
