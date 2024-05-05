package minswap

import "testing"

// Assisted by WCA@IBM
// Latest GenAI contribution: ibm/granite-20b-code-instruct-v2
// TestMinimumSwaps tests the minimumSwaps function.
func TestMinimumSwaps(t *testing.T) {
	tests := []struct {
		name string
		arr  []int32
		want int32
	}{
		// {
		// 	name: "Test case 1",
		// 	arr:  []int32{4, 3, 1, 2},
		// 	want: 3,
		// },
		{
			name: "Test case 2",
			arr:  []int32{2, 3, 4, 1, 5},
			want: 3,
		},
		// {
		// 	name: "Test case 3",
		// 	arr:  []int32{1, 3, 5, 2, 4, 6, 7},
		// 	want: 3,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSwaps(tt.arr); got != tt.want {
				t.Errorf("minimumSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
