package test

type CarSlice []Car
type CarMapFunc func(Car) interface{}

func (slice CarSlice) Map(fn CarMapFunc) []interface{} {
	var output []interface{}
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
