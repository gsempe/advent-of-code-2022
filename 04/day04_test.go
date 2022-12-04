package day04

import (
	"testing"
)

func TestPair_areOverlapping(t *testing.T) {
	type fields struct {
		FirstRange  Range
		SecondRange Range
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "not overlap",
			fields: fields{
				FirstRange:  Range{Min: 2, Max: 4},
				SecondRange: Range{Min: 6, Max: 8},
			},
			want: false,
		},
		{
			name: "not overlap",
			fields: fields{
				FirstRange:  Range{Min: 2, Max: 3},
				SecondRange: Range{Min: 4, Max: 5},
			},
			want: false,
		},
		{
			name: "overlap",
			fields: fields{
				FirstRange:  Range{Min: 2, Max: 7},
				SecondRange: Range{Min: 7, Max: 9},
			},
			want: true,
		},
		{
			name: "overlap",
			fields: fields{
				FirstRange:  Range{Min: 6, Max: 6},
				SecondRange: Range{Min: 4, Max: 6},
			},
			want: true,
		},
		{
			name: "included",
			fields: fields{
				FirstRange:  Range{Min: 2, Max: 8},
				SecondRange: Range{Min: 3, Max: 7},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pair{
				FirstRange:  tt.fields.FirstRange,
				SecondRange: tt.fields.SecondRange,
			}
			if got := p.areOverlapping(); got != tt.want {
				t.Errorf("areOverlapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
