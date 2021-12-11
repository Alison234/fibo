package fibonaci

import "errors"

type Fibonaci struct {
	Index int `json:"index"`
	Value int `json:"value"`
}

type FibonaciProvider struct {
	FibonaciSequence []Fibonaci
}

func (f *FibonaciProvider) NewFibonaciProvider() *FibonaciProvider {
	return &FibonaciProvider{FibonaciSequence: []Fibonaci{}}
}

func (f *FibonaciProvider) Calculate(startIndex int, endIndex int) error {
	if startIndex < 0 || endIndex < 0 {
		return errors.New("index must be above zero")
	}

	if startIndex > endIndex {
		return errors.New("end index must be more start")
	}

	f.FibonaciSequence = makeFibonaci(startIndex, endIndex)
	return nil
}

func makeFibonaci(startIndex int, endIndex int) []Fibonaci {
	result := make([]Fibonaci, 0)
	result = append(result, Fibonaci{Index: 0, Value: 0})
	x := 0
	y := 1
	value := 1

	for i := 0; i < endIndex; i++ {
		result = append(result, Fibonaci{Index: i + 1, Value: value})
		value = x + y
		x = y
		y = value
	}

	return result[startIndex : endIndex+1]
}
