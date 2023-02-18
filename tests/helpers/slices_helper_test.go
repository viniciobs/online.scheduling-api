package tests

import (
	"testing"

	"github.com/online.scheduling-api/src/helpers"
)

func TestShouldConcatIntegersCorrectly(t *testing.T) {
	// Arrange
	oneToFive := []int{1, 2, 3, 4, 5}
	sixToTen := []int{6, 7, 8, 9, 10}
	elevenToFifteen := []int{11, 12, 13, 14, 15}

	// Act
	result := helpers.Concat([][]int{
		oneToFive,
		sixToTen,
		elevenToFifteen,
	})

	// Assert
	for i := 1; i <= 15; i++ {
		if result[i-1] != i {
			t.Errorf("Expected %d but got %d", i, result[i-1])
		}
	}
}

func TestShouldConcatStructCorrectly(t *testing.T) {
	// Arrange
	type Person struct {
		Name string
		Age  int
	}

	babies := []Person{
		{Name: "Seu Antônio", Age: 1},
		{Name: "Dona Cilia", Age: 3},
	}

	elder := []Person{
		{Name: "Carol", Age: 88},
		{Name: "Beatriz", Age: 70},
	}

	// Act
	result := helpers.Concat([][]Person{
		babies,
		elder,
	})

	// Assert
	if len(result) != 4 {
		t.Error("Expected to have 4 elements in array")
	}
	if babies[0].Name != "Seu Antônio" &&
		babies[0].Age != 1 {
		t.Error("Expected another object")
	}
}
