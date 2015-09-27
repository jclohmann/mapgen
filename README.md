MapGen
======
This package generates map-, filter and each-methods for slices.

##Building
```
go get github.com/jclohmann/mapgen
go install github.com/jclohmann/mapgen/...
```

##Usage
Add a comment of the following format to your source-file:
```
package test

// +mapgen type=Car
type Car struct {
	brand string
	color string
}
```
Call the mapgen-executable:
```
mapgen -source car.go
```
Use the generated methods in your code:
```
var cars CarSlice
cars = append(cars, Car{"BMW", "red"})
cars = append(cars, Car{"VW", "green"})

brands := cars.Map(func(car Car) interface{} {
	return car.brand
})
```

Todo
====
* go generate-Integration
