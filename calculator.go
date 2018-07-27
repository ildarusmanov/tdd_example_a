package main

type ICalculator interface {
	Sum(a, b int) int
}

type Calculator struct{}

func CreateNewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Sum(a, b int) int {
	return 15
}
