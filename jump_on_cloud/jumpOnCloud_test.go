package jumponcloud

import "testing"

func TestJumpingOnClouds(t *testing.T) {
	cases := []struct {
		name     string
		clouds   []int32
		expected int32
	}{
		{"Example 1", []int32{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, 6},
		{"Example 2", []int32{0, 0, 0, 1, 0, 0}, 3},
		{"Example 3", []int32{0, 0, 1, 0, 0, 1, 0}, 4},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := jumpingOnClouds(tc.clouds)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
