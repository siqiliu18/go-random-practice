package spiralmatrix

import (
	"reflect"
	"testing"
)

// TestSpiralOrder tests the spiral order function
func TestSpiralOrder1(t *testing.T) {
	// Test case 1
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	result := spiralOrder(matrix)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected, result)
	}
}

func TestSpiralOrder2(t *testing.T) {
	// Test case 2
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	result := spiralOrder(matrix)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected, result)
	}
}

func TestSpiralOrder3(t *testing.T) {
	// Test case 2
	matrix := [][]int{{1}}
	expected := []int{1}
	result := spiralOrder(matrix)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected, result)
	}
}

func TestSpiralOrder4(t *testing.T) {
	// Test case 2
	matrix := [][]int{{2, 3}}
	expected := []int{2, 3}
	result := spiralOrder(matrix)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected, result)
	}
}
