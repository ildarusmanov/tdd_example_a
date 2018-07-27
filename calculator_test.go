package main_test

import (
	. "github.com/ildarusmanov/tdd_example_a"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
	var (
		a          int
		b          int
		calculator *Calculator
	)

	BeforeEach(func() {
		calculator = CreateNewCalculator()
		a = 5
		b = 10
	})

	Describe("Sum()", func() {
		Context("With a and b", func() {
			It("should be equal to expected sum", func() {
				Expect(calculator.Sum(a, b)).To(Equal(a + b))
				Expect(calculator.Sum(a, 1)).To(Equal(a + 1))
				Expect(calculator.Sum(b, 1)).To(Equal(b + 1))
			})
		})
	})
})
