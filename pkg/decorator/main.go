package main

import (
	"fmt"
)

type BeverageIf interface {
	getCost() int
	setCost(int)
	getDescription() string
	setDescription(string)
	withDeco(decoratorIf DecoratorIf) DecoratorIf
}

type Americano struct {
	cost int
	description string
}

func NewAmericano() *Americano{
	return &Americano{
		cost: 1000,
		description: "아메리카노",
	}
}

func (a *Americano) getCost() int {
	return a.cost
}
func (a *Americano) setCost(cost int)  {
	a.cost = cost
}

func (a *Americano) getDescription() string {
	return a.description
}

func (a *Americano) setDescription(desc string)  {
	a.description = desc
}

func (a *Americano) withDeco(deco DecoratorIf) DecoratorIf {
	deco.setCost(deco.getCost()+a.getCost())
	deco.setDescription(fmt.Sprintf("%s with %s",a.getDescription(), deco.getDescription()))
	return deco
}

type DecoratorIf interface {
	BeverageIf
	build() BeverageIf
}

type Milk struct {
	a BeverageIf
	cost int
	description string
}

func NewMilk() *Milk {
	return &Milk{
		description: "우유",
		cost: 500,
	}
}

func (m *Milk) getCost() int {
	return m.cost
}

func (m *Milk) setCost(i int) {
	m.cost = i
}

func (m *Milk) getDescription() string {
	return m.description
}

func (m *Milk) setDescription(s string) {
	m.description = s
}

func (m *Milk) withDeco(deco DecoratorIf) DecoratorIf {
	deco.setCost(deco.getCost()+m.getCost())
	deco.setDescription(fmt.Sprintf("%s with %s",m.getDescription(), deco.getDescription()))
	return deco
}

func (m *Milk) build() BeverageIf{
	return m
}

type Whip struct {
	a BeverageIf
	cost int
	description string
}

func NewWhip() *Whip{
	return &Whip{
		cost: 200,
		description: "휘핑",
	}
}

func (w *Whip) getCost() int {
	return w.cost
}

func (w *Whip) setCost(i int) {
	w.cost = i
}

func (w *Whip) getDescription() string {
	return w.description
}

func (w *Whip) setDescription(s string) {
	w.description = s
}

func (w *Whip) withDeco(deco DecoratorIf) DecoratorIf {
	deco.setCost(deco.getCost()+w.getCost())
	deco.setDescription(fmt.Sprintf("%s with %s",w.getDescription(), deco.getDescription()))
	return deco
}

func (w *Whip) build() BeverageIf {
	return w
}

func main(){
	americanoWithMilk := NewAmericano().withDeco(NewMilk()).withDeco(NewWhip()).build()
	fmt.Println(americanoWithMilk.getDescription())
	fmt.Println(americanoWithMilk.getCost())
}
