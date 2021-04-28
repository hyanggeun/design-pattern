package main

import "fmt"

type CharacterIf interface {
	setWeapon(behavior WeaponBehavior)
	setName(name string)
	fight()
}

type Character struct {
	w    WeaponBehavior
	Name string
}

func (c *Character) setWeapon(behavior WeaponBehavior) {
	c.w = behavior
}

func (c *Character) setName(name string) {
	c.Name = name
}

func (c *Character) fight() {
	fmt.Printf("%s: fight using %s", c.Name, c.w.useWeapon())
}

type King struct {
	Character
}

type Queen struct {
	Character
}
