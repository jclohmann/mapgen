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

func TestFilterCar(t *testing.T) {
	var cars CarSlice
	cars = append(cars, Car{"BMW", "white"})
	cars = append(cars, Car{"VW", "Black"})
	cars = append(cars, Car{"Seat", "Black"})
	cars = append(cars, Car{"VW", "silver"})
	cars = append(cars, Car{"Ford", "brown"})

	vws := cars.Filter(func(car Car) bool {
		return car.brand == "VW"
	})

	if len(vws) != 2 {
		t.Errorf("the length of the slice should be %v, but was %v", 2, len(vws))
	}

	for _, car := range vws {
		if car.brand != "VW" {
			t.Errorf("%v should be in the slice", car)
		}
	}
}
