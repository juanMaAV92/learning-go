package main

import "testing"

func TestSum(t *testing.T) {

	// total := Sum(5, 5)
	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	// }

	tableTest := []struct {
		description string
		a           int
		b           int
		total       int
	}{
		{
			description: "Sum of 5 + 5",
			a:           5,
			b:           5,
			total:       10,
		},
		{
			description: "Sum of 12 + 26",
			a:           12,
			b:           26,
			total:       38,
		},
	}

	for _, item := range tableTest {
		t.Run(item.description, func(t *testing.T) {
			if totalSum := Sum(item.a, item.b); totalSum != item.total {
				t.Errorf("Sum was incorrect, got: %d, want: %d.", totalSum, item.total)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	tableTest := []struct {
		description string
		a           int
		total       int
	}{
		{
			description: "Fibonacci of 0",
			a:           0,
			total:       0,
		}, {
			description: "Fibonacci of 1",
			a:           1,
			total:       1,
		}, {
			description: "Fibonacci of 10",
			a:           10,
			total:       55,
		}, {
			description: "Fibonacci of 40",
			a:           40,
			total:       102334155,
		},
	}

	for _, item := range tableTest {
		t.Run(item.description, func(t *testing.T) {
			if totalFibonacci := Fibonacci(item.a); totalFibonacci != item.total {
				t.Errorf("Fibonacci was incorrect, got: %d, want: %d.", totalFibonacci, item.total)
			}
		})
	}
}
