package lc253meetingrooms

import "testing"

// Assisted by WCA@IBM
// Latest GenAI contribution: ibm/granite-20b-code-instruct-v2
// TestMinMeetingRooms tests the minMeetingRooms function.
func TestMinMeetingRooms(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				intervals: [][]int{{7, 10}, {2, 4}},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				intervals: [][]int{{7, 10}, {2, 4}, {15, 20}},
			},
			want: 1,
		},
		{
			name: "case 4",
			args: args{
				intervals: [][]int{{9, 10}, {4, 17}, {4, 9}},
			},
			want: 2,
		},
		{
			name: "case 5",
			args: args{
				intervals: [][]int{{9, 10}, {4, 17}, {4, 8}, {9, 15}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMeetingRooms(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
