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
