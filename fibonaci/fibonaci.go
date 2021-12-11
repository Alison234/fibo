package fibonaci

import "errors"

type Fibonaci struct {
	index int
	value int
}

func MakeFibonaci(startIndex int, endIndex int) ([]Fibonaci, error) {
	if startIndex < 0 || endIndex < 0 {
		return nil, errors.New("index must be above zero")
	}

	if startIndex > endIndex {
		return nil, errors.New("end index must be more start")
	}
	result := make([]Fibonaci, 0)
	result = append(result, Fibonaci{index: 0, value: 0})
	x := 0
	y := 1
	value := 1

	for i := 0; i < endIndex; i++ {
		result = append(result, Fibonaci{index: i + 1, value: value})
		value = x + y
		x = y
		y = value
	}

	return result[startIndex : endIndex+1], nil
}
