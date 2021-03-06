package test

/*
	DO NOT EDIT THIS FILE... or your changes will be overwritten.
	This file was generated by mapgen. (https://github.com/jclohmann/mapgen)
*/

type CarSlice []Car



type CarMapToCarFunc func(Car) Car

func (slice CarSlice) MapToCar(fn CarMapToCarFunc) []Car {
	var output []Car
	for _, item := range slice {
		out := fn(item)
		output = append(output, out)
	}
	return output
}

type CarMapToInterfaceFunc func(Car) interface{}

func (slice CarSlice) MapToInterface(fn CarMapToInterfaceFunc) []interface{} {
	var output []interface{}
	for _, item := range slice {
		out := fn(item)
		output = append(output, out)
	}
	return output
}

type CarMapToStringFunc func(Car) string

func (slice CarSlice) MapToString(fn CarMapToStringFunc) []string {
	var output []string
	for _, item := range slice {
		out := fn(item)
		output = append(output, out)
	}
	return output
}


type CarFilterFunc func(Car) bool

func (slice CarSlice) Filter(fn CarFilterFunc) CarSlice {
	var output CarSlice
	for _, item := range slice {
		if fn(item) {
			output = append(output, item)
		}
	}
	return output
}

type CarEachFunc func(Car)

func (slice CarSlice) Each(fn CarEachFunc) {
	for _, item := range slice {
		fn(item)
	}
}
