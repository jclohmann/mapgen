package test

import (
	"testing"
)

func TestMapCar(t *testing.T) {
	var cars CarSlice
	cars = append(cars, Car{"BMW", "white"})
	cars = append(cars, Car{"VW", "Black"})
	cars = append(cars, Car{"Seat", "Black"})
	cars = append(cars, Car{"VW", "silver"})
	cars = append(cars, Car{"Ford", "brown"})

	names := cars.Map(func(car Car) interface{} {
		return car.brand
	})

	for index, name := range names {
		car := cars[index]
		if car.brand != name {
			t.Errorf("name should be %v, but was %v", car.brand, name)
		}
	}
}
