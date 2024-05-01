package nychaos

import "testing"

func TestMinimumBribes1(t *testing.T) {
	q := []int32{2, 1, 5, 3, 4}
	expected := int32(3)
	result := minimumBribes(q)
	if result != expected {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected, result)
	}
}

func TestMinimumBribes2(t *testing.T) {
	q := []int32{5, 1, 2, 3, 7, 8, 6, 4}
	expected := int32(-1)
	result := minimumBribes(q)
	if result != expected {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected, result)
	}
}

func TestMinimumBribes3(t *testing.T) {
	q := []int32{1, 2, 5, 3, 7, 8, 6, 4}
	expected := int32(7)
	result := minimumBribes(q)
	if result != expected {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected, result)
	}
}

func TestMinimumBribes4(t *testing.T) {
	q := []int32{2, 5, 1, 3, 4}
	expected := int32(-1)
	result := minimumBribes(q)
	if result != expected {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected, result)
	}
}

func TestMinimumBribes5(t *testing.T) {
	q := []int32{1, 2, 5, 3, 4, 7, 8, 6}
	expected := int32(4)
	result := minimumBribes(q)
	if result != expected {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected, result)
	}
}
