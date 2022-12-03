package day03

import (
	"testing"
)

func Test_rucksackToCompartment(t *testing.T) {
	type args struct {
		rucksack string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name: "Simple test",
			args: args{
				rucksack: "abcdef",
			},
			want:  "abc",
			want1: "def",
		},
		{
			name: "Odd",
			args: args{
				rucksack: "abcdefg",
			},
			want:  "abc",
			want1: "defg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := rucksackToCompartment(tt.args.rucksack)
			if got != tt.want {
				t.Errorf("rucksackToCompartment() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("rucksackToCompartment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_itemTypeToPriority(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test b",
			args: args{
				i: "b",
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "Test z",
			args: args{
				i: "z",
			},
			want:    26,
			wantErr: false,
		},
		{
			name: "Test A",
			args: args{
				i: "A",
			},
			want:    27,
			wantErr: false,
		},
		{
			name: "Test Z",
			args: args{
				i: "A",
			},
			want:    52,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := itemTypeToPriority(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("itemTypeToPriority() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("itemTypeToPriority() got = %v, want %v", got, tt.want)
			}
		})
	}
}
