package fibonacci

import "errors"

type Fibonacci struct {
	Index int `json:"index"`
	Value int `json:"value"`
}

type FibonacciProvider struct {
	FibonacciSequence []Fibonacci
}

func NewFibonacciProvider() *FibonacciProvider {
	return &FibonacciProvider{FibonacciSequence: []Fibonacci{}}
}

func (f *FibonacciProvider) Calculate(startIndex int, endIndex int) error {
	if startIndex < 0 || endIndex < 0 {
		return errors.New("index must be above zero")
	}

	if startIndex > endIndex {
		return errors.New("end index must be more start")
	}

	f.FibonacciSequence = makeFibonacci(startIndex, endIndex)
	return nil
}

func makeFibonacci(startIndex int, endIndex int) []Fibonacci {
	result := make([]Fibonacci, 0)
	result = append(result, Fibonacci{Index: 0, Value: 0})
	x := 0
	y := 1
	value := 1

	for i := 0; i < endIndex; i++ {
		result = append(result, Fibonacci{Index: i + 1, Value: value})
		value = x + y
		x = y
		y = value
	}

	return result[startIndex : endIndex+1]
}
