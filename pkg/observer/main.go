package main

import (
	"fmt"
)

type subject interface {
	register(o observer)
	remove(o observer)
	notify()
}

type observer interface {
	update(int, int , int)
}

type CurrentConditionDisplay struct {
	temperature int
	pressure    int
	humidity    int
	weatherData *weatherData
}

func NewCurrentConditionDisplay(w *weatherData) *CurrentConditionDisplay {
	cd := &CurrentConditionDisplay{
		weatherData: w,
	}
	return cd
}

func (c *CurrentConditionDisplay) update(temperature, pressure, humidity int) {
	c.temperature = temperature
	c.pressure = pressure
	c.humidity = humidity
	c.display()
}

func (c *CurrentConditionDisplay) display() {
		fmt.Printf("current condition: temp: %d, pressure: %d, humidity: %d\n",c.temperature,c.pressure,c.humidity)
}

type weatherData struct {
	temperature int
	pressure int
	humidity int
	observers []observer
}

func NewWeatherData() *weatherData{
	return &weatherData{
		observers: make([]observer,0),
	}
}

func (w *weatherData) register(o observer) {
	w.observers = append(w.observers, o)

}

func (w *weatherData) remove(o observer) {
	newList := make([]observer, 0)
	for _, selObserver := range w.observers {
		if selObserver != o {
			newList = append(newList,selObserver)
		}
	}
	w.observers = newList
}

func (w *weatherData) notify() {
	for _, obs := range w.observers {
		obs.update(w.temperature,w.pressure,w.humidity)
	}
}

func (w *weatherData) measurementChanged(temperature, pressure, humidity int) {
	w.temperature = temperature
	w.pressure = pressure
	w.humidity = humidity
	w.notify()
}


func main() {
	wd := NewWeatherData()
	currentCondition := NewCurrentConditionDisplay(wd)

	wd.register(currentCondition)
	wd.measurementChanged(1,2,3)

	wd.remove(currentCondition)
	wd.measurementChanged(3,3,3)

}