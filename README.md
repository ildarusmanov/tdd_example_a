# TDD with Golang

* [https://github.com/golang/dep](https://github.com/golang/dep)
* [http://onsi.github.io/ginkgo/](http://onsi.github.io/ginkgo/)
* [https://onsi.github.io/gomega/](https://onsi.github.io/gomega/)

## Setup

```
go get -u github.com/golang/dep/cmd/dep
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega/...

cd [project_path]

dep init
ginkgo bootstrap
```

## Generate Test
```
ginkgo generate book
```

## Write down the test into generated file
```
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
```

## Add code
```
package main

type Calculator struct{}

func CreateNewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Sum(a, b int) int {
	return a + b
}
```