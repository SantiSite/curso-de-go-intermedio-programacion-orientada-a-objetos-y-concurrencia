package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(5, 5)

	if total != 10 {
		t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	}

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total = sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d, expected %d", total, item.n)
		}
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(5, 3)

	if result != 2 {
		t.Errorf("Subtract was incorrect, got %d, expected %d", result, 2)
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		description string
		a           int
		b           int
		n           int
	}{
		{"TestMax 1", 4, 2, 4},
		{"TestMax 2", 3, 2, 3},
		{"TestMax 3", 20, 10, 20},
		{"TestMax 4", 20, 31, 31},
	}

	for _, item := range tables {
		t.Run(item.description, func(t *testing.T) {
			max := getMax(item.a, item.b)
			if max != item.n {
				t.Errorf("getMax was incorrect, got %d, expected %d", max, item.n)
			}
		})
	}
}

func TestFib(t *testing.T) {
	table := []struct {
		a int
		n int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{50, 12_586_269_025},
	}

	for _, item := range table {
		fib := fibonacci(item.a)
		if fib != item.n {
			t.Errorf("fibonacci was incorret, got %d, expected %d", fib, item.n)
		}
	}
}
